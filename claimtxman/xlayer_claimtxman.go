package claimtxman

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/0xPolygonHermez/zkevm-bridge-service/bridgectrl/pb"
	"github.com/0xPolygonHermez/zkevm-bridge-service/etherman"
	"github.com/0xPolygonHermez/zkevm-bridge-service/log"
	"github.com/0xPolygonHermez/zkevm-bridge-service/utils"
	"github.com/0xPolygonHermez/zkevm-bridge-service/utils/gerror"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/messagepush"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/metrics"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/redisstorage"

	ctmtypes "github.com/0xPolygonHermez/zkevm-bridge-service/claimtxman/types"
)

var (
	// Flags/configs
	freeGas  bool
	optClaim bool

	updateDepositsL1Mutex sync.Mutex
	updateDepositsL2Mutex sync.Mutex
	isDone                bool

	// Producer to push the transaction status change to front end
	messagePushProducer messagepush.KafkaProducer
	redisStorage        redisstorage.RedisStorage
	monitorTxsLimit     uint = 128

	// NOTE: originated from zkevm-node package (which is now removed)
	// ErrNonceTooLow is returned if the nonce of a transaction is lower than the
	// one present in the local chain.
	ErrNonceTooLow = errors.New("nonce too low")
	// ErrNonceTooHigh is returned if the nonce of a transaction is higher than the
	// current + the configured AccountQueue.
	ErrNonceTooHigh = errors.New("nonce too high")
)

func (tm *ClaimTxManager) SetProducer(producer messagepush.KafkaProducer) *ClaimTxManager {
	messagePushProducer = producer
	return tm
}

func (tm *ClaimTxManager) SetRedis(rs redisstorage.RedisStorage) *ClaimTxManager {
	redisStorage = rs
	return tm
}

func (tm *ClaimTxManager) SetFreegas(isFree bool) *ClaimTxManager {
	freeGas = isFree
	return tm
}

func (tm *ClaimTxManager) SetOptClaim(flag bool) *ClaimTxManager {
	optClaim = flag
	return tm
}

func (tm *ClaimTxManager) SetMonitorTxsLimit(limit uint) *ClaimTxManager {
	monitorTxsLimit = limit
	return tm
}

// StartXLayer will start the tx management, reading txs from storage,
// send then to the blockchain and keep monitoring them until they
// get mined
func (tm *ClaimTxManager) StartXLayer() {
	go tm.startMonitorTxs()

	compressorTicker := time.NewTicker(tm.cfg.GroupingClaims.FrequencyToProcessCompressedClaims.Duration)
	var ger = &etherman.GlobalExitRoot{}
	var latestProcessedGer common.Hash
	for {
		select {
		case <-tm.ctx.Done():
			isDone = true
			return
		case netID := <-tm.chSynced:
			if netID == tm.l2NetworkID && !tm.l2Synced {
				log.Info("NetworkID synced: ", netID)
				tm.l2Synced = true
			}
			if netID == 0 && !tm.l1Synced {
				log.Infof("RollupID: %d, L1 network synced. GERs ready for ClaimTxManager", tm.rollupID)
				tm.l1Synced = true
			}
		case ger = <-tm.chExitRootEvent:
			if tm.l2Synced && tm.l1Synced {
				log.Debugf("RollupID: %d UpdateDepositsStatus for ger: %s", tm.rollupID, ger.GlobalExitRoot.String())
				if tm.cfg.GroupingClaims.Enabled {
					log.Debugf("rollupID: %d, Ger value updated and ready to be processed...", tm.rollupID)
					continue
				}
				go func() {
					err := tm.updateDepositsStatusXLayer(ger)
					if err != nil {
						log.Errorf("rollupID: %d, failed to update deposits status: %v", tm.rollupID, err)
					}
				}()
			} else {
				log.Infof("Waiting for networkID %d to be synced before processing deposits", tm.l2NetworkID)
			}
		case <-compressorTicker.C:
			if tm.l2Synced && tm.l1Synced && tm.cfg.GroupingClaims.Enabled && ger.GlobalExitRoot != latestProcessedGer {
				log.Infof("RollupID: %d,Processing deposits for ger: %s", tm.rollupID, ger.GlobalExitRoot.String())
				go func() {
					err := tm.updateDepositsStatusXLayer(ger)
					if err != nil {
						log.Errorf("rollupID: %d, failed to update deposits status: %v", tm.rollupID, err)
					}
				}()
				latestProcessedGer = ger.GlobalExitRoot
			}
		}
	}
}

func (tm *ClaimTxManager) startMonitorTxs() {
	ticker := time.NewTicker(tm.cfg.FrequencyToMonitorTxs.Duration)
	for range ticker.C {
		if isDone {
			return
		}
		traceID := utils.GenerateTraceID()
		ctx := context.WithValue(tm.ctx, utils.CtxTraceID, traceID)
		logger := log.WithFields(utils.TraceID, traceID)
		logger.Infof("MonitorTxs begin %d", tm.l2NetworkID)
		err := tm.monitorTxsXLayer(ctx)
		if err != nil {
			logger.Errorf("failed to monitor txs: %v", err)
		}
		logger.Infof("MonitorTxs end %d", tm.l2NetworkID)
	}
}

func (tm *ClaimTxManager) monitorTxsXLayer(ctx context.Context) error {
	mLog := log.WithFields(utils.TraceID, ctx.Value(utils.CtxTraceID))

	dbTx, err := tm.storage.BeginDBTransaction(ctx)
	if err != nil {
		return err
	}
	mLog.Infof("monitorTxs begin")

	statusesFilter := []ctmtypes.MonitoredTxStatus{ctmtypes.MonitoredTxStatusCreated}
	mTxs, err := tm.storage.GetClaimTxsByStatusWithLimit(ctx, statusesFilter, monitorTxsLimit, 0, dbTx)
	if err != nil {
		mLog.Errorf("failed to get created monitored txs: %v", err)
		rollbackErr := tm.storage.Rollback(tm.ctx, dbTx)
		if rollbackErr != nil {
			mLog.Errorf("claimtxman error rolling back state. RollbackErr: %s, err: %v", rollbackErr.Error(), err)
			return rollbackErr
		}
		return fmt.Errorf("failed to get created monitored txs: %v", err)
	}
	mLog.Infof("found %v monitored tx to process", len(mTxs))
	metrics.RecordPendingMonitoredTxsCount(len(mTxs))

	isResetNonce := false // it will reset the nonce in one cycle
	for _, mTx := range mTxs {
		mTx := mTx // force variable shadowing to avoid pointer conflicts
		mTxLog := mLog.WithFields("monitoredTx", mTx.DepositID)
		mTxLog.Infof("processing tx with nonce %d", mTx.Nonce)
		// Check the claim table to see whether the transaction has already been claimed by some other methods
		_, err = tm.storage.GetClaim(ctx, uint32(mTx.DepositID), 0, tm.l2NetworkID, dbTx) // TODO: original network ID?
		if err != nil && err != gerror.ErrStorageNotFound {
			mTxLog.Errorf("failed to get claim tx: %v", err)
			return err
		}
		if err == nil {
			mTxLog.Infof("Tx has already been claimed")
			mTx.Status = ctmtypes.MonitoredTxStatusConfirmed
			// Update monitored txs status to confirmed
			err = tm.storage.UpdateClaimTx(ctx, mTx, dbTx)
			if err != nil {
				mTxLog.Errorf("failed to update tx status to confirmed: %v", err)
			}
			metrics.RecordMonitoredTxsResult(string(mTx.Status))
			metrics.RecordAutoClaimDuration(time.Since(mTx.CreatedAt))
			continue
		}

		// check if any of the txs in the history was mined
		mined := false
		var receipt *types.Receipt
		hasFailedReceipts := false
		allHistoryTxMined := true
		receiptSuccessful := false

		for txHash := range mTx.History {
			mTxLog.Infof("Checking if tx %s is mined", txHash.String())
			mined, receipt, err = tm.l2Node.CheckTxWasMined(ctx, txHash)
			if err != nil {
				mTxLog.Errorf("failed to check if tx %s was mined: %v", txHash.String(), err)
				continue
			}

			// if the tx is not mined yet, check that not all the tx were mined and go to the next
			if !mined {
				// check if the tx is in the pending pool
				_, _, err = tm.l2Node.TransactionByHash(ctx, txHash)
				if err != nil {
					mTxLog.Errorf("error getting txByHash %s. Error: %v", txHash.String(), err)
					// Retry if the tx has not appeared in the pool yet.
					for i := 0; i < tm.cfg.RetryNumber && err != nil; i++ {
						mTxLog.Warn("waiting and retrying to find the tx in the pool. TxHash: %s. Error: %v", txHash.String(), err)
						time.Sleep(tm.cfg.RetryInterval.Duration)
						_, _, err = tm.l2Node.TransactionByHash(ctx, txHash)
					}
					if errors.Is(err, ethereum.NotFound) {
						mTxLog.Error("maximum retries and the tx is still missing in the pool. TxHash: ", txHash.String())
						hasFailedReceipts = true
						continue
					} else if err != nil {
						mTxLog.Errorf("failed to retry to get tx %s: %v", txHash.String(), err)
						continue
					}
				}
				mTxLog.Infof("tx: %s not mined yet", txHash.String())

				allHistoryTxMined = false
				continue
			}

			// if the tx was mined successfully we can break the loop and proceed
			if receipt.Status == types.ReceiptStatusSuccessful {
				mTxLog.Infof("tx %s was mined successfully", txHash.String())
				receiptSuccessful = true
				mTx.Status = ctmtypes.MonitoredTxStatusConfirmed
				// update monitored tx changes into storage
				err = tm.storage.UpdateClaimTx(ctx, mTx, dbTx)
				if err != nil {
					mTxLog.Errorf("failed to update monitored tx when confirmed: %v", err)
				}
				metrics.RecordMonitoredTxsResult(string(mTx.Status))
				metrics.RecordAutoClaimDuration(time.Since(mTx.CreatedAt))
				break
			}

			// if the tx was mined but failed, we continue to consider it was not mined
			// and store the failed receipt to be used to check if nonce needs to be reviewed
			mined = false
			hasFailedReceipts = true
		}

		if receiptSuccessful {
			continue
		}

		// if the history size reaches the max history size, this means something is really wrong with
		// this Tx and we are not able to identify automatically, so we mark this as failed to let the
		// caller know something is not right and needs to be review and to avoid to monitor this
		// tx infinitely
		if allHistoryTxMined && len(mTx.History) >= maxHistorySize {
			mTx.Status = ctmtypes.MonitoredTxStatusFailed
			mTxLog.Infof("marked as failed because reached the history size limit (%d)", maxHistorySize)
			// update monitored tx changes into storage
			err = tm.storage.UpdateClaimTx(ctx, mTx, dbTx)
			if err != nil {
				mTxLog.Errorf("failed to update monitored tx when max history size limit reached: %v", err)
			}
			metrics.RecordMonitoredTxsResult(string(mTx.Status))

			// Notify FE that tx is pending user claim
			go func() {
				// Retrieve L1 transaction info
				deposit, err := tm.storage.GetDeposit(ctx, uint32(mTx.DepositID), 0, nil)
				if err != nil {
					mTxLog.Errorf("push message: GetDeposit error: %v", err)
					return
				}
				tm.pushTransactionUpdate(deposit, uint32(pb.TransactionStatus_TX_PENDING_USER_CLAIM))
			}()
			continue
		}

		// if we have failed receipts, this means at least one of the generated txs was mined
		// so maybe the current nonce was already consumed, then we need to check if there are
		// tx that were not mined yet, if so, we just need to wait, because maybe one of them
		// will get mined successfully
		if allHistoryTxMined {
			// in case of all tx were mined and none of them were mined successfully, we need to
			// review the tx information
			if hasFailedReceipts {
				mTxLog.Infof("monitored tx needs to be updated")
				err := tm.ReviewMonitoredTxXLayer(ctx, &mTx)
				if err != nil {
					mTxLog.Errorf("failed to review monitored tx: %v", err)
					continue
				}
			}

			// GasPrice is set here to use always the proper and most accurate value right before sending it to L2
			gasPrice := big.NewInt(0)
			if !freeGas {
				gasPrice, err = tm.l2Node.SuggestGasPrice(ctx)
				if err != nil {
					mTxLog.Errorf("failed to get suggested gasPrice. Error: %v", err)
					continue
				}
			}

			//Multiply gasPrice by 10 to increase the efficiency of the tx in the sequence
			mTx.GasPrice = big.NewInt(0).Mul(gasPrice, big.NewInt(10)) //nolint:gomnd
			mTxLog.Infof("Using gasPrice: %s. The gasPrice suggested by the network is %s", mTx.GasPrice.String(), gasPrice.String())

			// Calculate nonce before signing
			if err = tm.setTxNonce(&mTx); err != nil {
				mTxLog.Errorf("failed to set tx nonce: %v", err)
				continue
			}

			// rebuild transaction
			tx := mTx.Tx()
			mTxLog.Debugf("unsigned tx created for monitored tx")

			var signedTx *types.Transaction
			// sign tx
			signedTx, err = tm.auth.Signer(mTx.From, tx)
			if err != nil {
				mTxLog.Errorf("failed to sign tx %v created from monitored tx: %v", tx.Hash().String(), err)
				continue
			}
			mTxLog.Debugf("signed tx %v created using gasPrice: %s, nonce: %v", signedTx.Hash().String(), signedTx.GasPrice().String(), signedTx.Nonce())

			// add tx to monitored tx history
			err = mTx.AddHistory(signedTx)
			if errors.Is(err, ctmtypes.ErrAlreadyExists) {
				mTxLog.Infof("signed tx already existed in the history")
			} else if err != nil {
				mTxLog.Errorf("failed to add signed tx to monitored tx history: %v", err)
				continue
			}

			// check if the tx is already in the network, if not, send it
			_, _, err = tm.l2Node.TransactionByHash(ctx, signedTx.Hash())
			if errors.Is(err, ethereum.NotFound) {
				err := tm.l2Node.SendTransaction(ctx, signedTx)
				if err != nil {
					mTxLog.Errorf("failed to send tx %s to network: %v", signedTx.Hash().String(), err)
					if err.Error() == ErrNonceTooLow.Error() {
						mTxLog.Infof("nonce error detected, Nonce used: %d", signedTx.Nonce())
						if !isResetNonce {
							isResetNonce = true
							tm.nonceCache.Remove(mTx.From.Hex())
							mTxLog.Infof("nonce cache cleared for address %v", mTx.From.Hex())
						}
					}
					if err.Error() == ErrNonceTooHigh.Error() {
						mTxLog.Infof("nonce error detected, Nonce used: %d", signedTx.Nonce())
						if !isResetNonce {
							isResetNonce = true
							tm.nonceCache.Remove(mTx.From.Hex())
							mTxLog.Infof("nonce cache cleared for address %v", mTx.From.Hex())
						}
					}
					tm.nonceCache.decr(mTx.From)
					mTx.RemoveHistory(signedTx)
					// we should rebuild the monitored tx to fix the nonce
					err := tm.ReviewMonitoredTxXLayer(ctx, &mTx)
					if err != nil {
						mTxLog.Errorf("failed to review monitored tx: %v", err)
					}
				}
			} else if err != nil && !errors.Is(err, ethereum.NotFound) {
				mTxLog.Error("unexpected error getting TransactionByHash. Error: ", err)
			} else {
				mTxLog.Infof("signed tx %v already found in the network for the monitored tx.", signedTx.Hash().String())
			}

			// update monitored tx changes into storage
			err = tm.storage.UpdateClaimTx(ctx, mTx, dbTx)
			if err != nil {
				mTxLog.Errorf("failed to update monitored tx: %v", err)
				continue
			}
			mTxLog.Infof("signed tx %s added to the monitored tx history", signedTx.Hash().String())
		}
	}
	mLog.Infof("monitorTxs end")

	err = tm.storage.Commit(tm.ctx, dbTx)
	if err != nil {
		mLog.Errorf("UpdateClaimTx committing dbTx, err: %v", err)
		rollbackErr := tm.storage.Rollback(tm.ctx, dbTx)
		if rollbackErr != nil {
			mLog.Errorf("claimtxman error rolling back state. RollbackErr: %s, err: %v", rollbackErr.Error(), err)
			return rollbackErr
		}
		return err
	}

	mLog.Infof("monitorTxs committed")

	return nil
}
