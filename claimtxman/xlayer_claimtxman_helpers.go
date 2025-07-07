package claimtxman

import (
	"context"
	"fmt"
	"math/big"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"

	"github.com/0xPolygonHermez/zkevm-bridge-service/bridgectrl/pb"
	"github.com/0xPolygonHermez/zkevm-bridge-service/etherman"
	"github.com/0xPolygonHermez/zkevm-bridge-service/log"
	"github.com/0xPolygonHermez/zkevm-bridge-service/utils"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/metrics"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/pushtask"

	ctmtypes "github.com/0xPolygonHermez/zkevm-bridge-service/claimtxman/types"
)

func (tm *ClaimTxManager) updateDepositsStatusXLayer(ger *etherman.GlobalExitRoot) error {
	if optClaim {
		if ger.BlockID != 0 {
			updateDepositsL2Mutex.Lock()
			defer updateDepositsL2Mutex.Unlock()
			return tm.processDepositStatusL2(ger)
		} else {
			updateDepositsL1Mutex.Lock()
			defer updateDepositsL1Mutex.Unlock()
			return tm.processDepositStatusL1(ger)
		}
	}

	dbTx, err := tm.storage.BeginDBTransaction(tm.ctx)
	if err != nil {
		return err
	}
	if err = tm.processDepositStatusXLayer(ger, dbTx); err != nil {
		log.Errorf("error processing ger. Error: %v", err)
		rollbackErr := tm.storage.Rollback(tm.ctx, dbTx)
		if rollbackErr != nil {
			log.Errorf("claimtxman error rolling back state. RollbackErr: %v, err: %s", rollbackErr, err.Error())
			return rollbackErr
		}
		return err
	}
	if err = tm.storage.Commit(tm.ctx, dbTx); err != nil {
		log.Errorf("AddClaimTx committing dbTx. Err: %v", err)
		rollbackErr := tm.storage.Rollback(tm.ctx, dbTx)
		if rollbackErr != nil {
			log.Fatalf("claimtxman error rolling back state. RollbackErr: %s, err: %s", rollbackErr.Error(), err.Error())
		}
		log.Fatalf("AddClaimTx committing dbTx, err: %s", err.Error())
	}
	log.Debugf("updateDepositsStatus done")

	return nil
}

func (tm *ClaimTxManager) processDepositStatusL2(ger *etherman.GlobalExitRoot) error {
	dbTx, err := tm.storage.BeginDBTransaction(tm.ctx)
	if err != nil {
		return err
	}
	log.Infof("Rollup exitroot %v is updated", ger.ExitRoots[1])
	deposits, err := tm.storage.UpdateL2DepositsStatusXLayer(tm.ctx, ger.ExitRoots[1][:], uint(tm.rollupID), uint(tm.l2NetworkID), dbTx)
	if err != nil {
		log.Errorf("error getting and updating L2DepositsStatus. Error: %v", err)
		rollbackErr := tm.storage.Rollback(tm.ctx, dbTx)
		if rollbackErr != nil {
			log.Errorf("claimtxman error rolling back state. RollbackErr: %v, err: %s", rollbackErr, err.Error())
			return rollbackErr
		}
		return err
	}
	err = tm.storage.Commit(tm.ctx, dbTx)
	if err != nil {
		log.Errorf("AddClaimTx committing dbTx. Err: %v", err)
		rollbackErr := tm.storage.Rollback(tm.ctx, dbTx)
		if rollbackErr != nil {
			log.Fatalf("claimtxman error rolling back state. RollbackErr: %s, err: %s", rollbackErr.Error(), err.Error())
		}
		log.Fatalf("AddClaimTx committing dbTx, err: %s", err.Error())
	}
	log.Debugf("begin send deposits for l1 ready_claim, blockId: %v, blockNumber: %v, deposit size: %v", ger.BlockID, ger.BlockNumber,
		len(deposits))
	for _, deposit := range deposits {
		// Notify FE that tx is pending auto claim
		go tm.pushTransactionUpdate(deposit, uint32(pb.TransactionStatus_TX_PENDING_USER_CLAIM))
		// Record order waiting time metric
		metrics.RecordOrderWaitTime(uint32(deposit.NetworkID), uint32(deposit.DestinationNetwork), time.Since(deposit.Time))
	}
	return nil
}

func (tm *ClaimTxManager) processDepositStatusL1(newGer *etherman.GlobalExitRoot) error {
	deposits, err := tm.getDeposits(newGer)
	if err != nil {
		return err
	}
	log.Debug("createClaimTx deposits-num:", len(deposits))
	for _, deposit := range deposits {
		dbTx, err := tm.storage.BeginDBTransaction(tm.ctx)
		if err != nil {
			return err
		}
		if err = tm.storage.UpdateL1DepositStatus(tm.ctx, uint(deposit.DepositCount), dbTx); err != nil {
			log.Errorf("error update deposit %d status. Error: %v", deposit.DepositCount, err)
			tm.rollbackStore(dbTx)
			return err
		}
		ignore := false
		if tm.l2NetworkID != deposit.DestinationNetwork {
			log.Infof("Ignoring deposit: %d: dest_net: %d, we are:%d", deposit.DepositCount, deposit.DestinationNetwork, tm.l2NetworkID)
			ignore = true
		} else {
			claimHash, err := tm.bridgeService.GetDepositStatus(tm.ctx, deposit.DepositCount, deposit.NetworkID, deposit.DestinationNetwork)
			if err != nil {
				log.Errorf("error getting deposit status for deposit %d. Error: %v", deposit.DepositCount, err)
				tm.rollbackStore(dbTx)
				return err
			}
			if len(claimHash) > 0 || (deposit.LeafType == LeafTypeMessage && !tm.isDepositMessageAllowed(deposit)) {
				log.Infof("Ignoring deposit: %d, leafType: %d, claimHash: %s", deposit.DepositCount, deposit.LeafType, claimHash)
				ignore = true
			}
		}

		if ignore {
			// todo: optimize it
			err = tm.storage.Commit(tm.ctx, dbTx)
			if err != nil {
				log.Errorf("AddClaimTx committing dbTx. Err: %v", err)
				rollbackErr := tm.storage.Rollback(tm.ctx, dbTx)
				if rollbackErr != nil {
					log.Fatalf("claimtxman error rolling back state. RollbackErr: %s, err: %s", rollbackErr.Error(), err.Error())
					return rollbackErr
				}
				log.Fatalf("AddClaimTx committing dbTx, err: %s", err.Error())
				return err
			}
			continue
		}
		log.Infof("create the claim tx for the deposit %d", deposit.DepositCount)
		ger, proof, rollupProof, err := tm.bridgeService.GetClaimProof(deposit.DepositCount, deposit.NetworkID, dbTx)
		if err != nil {
			log.Errorf("error getting Claim Proof for deposit %d. Error: %v", deposit.DepositCount, err)
			tm.rollbackStore(dbTx)
			return err
		}
		log.Infof("get the claim proof for the deposit %d successfully", deposit.DepositCount)
		var (
			mtProof       [mtHeight][keyLen]byte
			mtRollupProof [mtHeight][keyLen]byte
		)
		for i := 0; i < mtHeight; i++ {
			mtProof[i] = proof[i]
			mtRollupProof[i] = rollupProof[i]
		}
		tx, err := tm.l2Node.BuildSendClaimXLayer(tm.ctx, deposit, mtProof, mtRollupProof,
			&etherman.GlobalExitRoot{
				ExitRoots: []common.Hash{
					ger.ExitRoots[0],
					ger.ExitRoots[1],
				}}, 1, 1, 1, uint(tm.rollupID),
			tm.auth)
		if err != nil {
			log.Errorf("error BuildSendClaimXLayer tx for deposit %d. Error: %v", deposit.DepositCount, err)
			tm.rollbackStore(dbTx)
			return err
		}
		if err = tm.addClaimTxXLayer(uint(deposit.DepositCount), tm.auth.From, tx.To(), nil, tx.Data(), dbTx); err != nil {
			log.Errorf("error adding claim tx for deposit %d. Error: %v", deposit.DepositCount, err)
			tm.rollbackStore(dbTx)
			return err
		}

		// There can be cases that the deposit can be ready for claim (and even claimed) before it reached 64 block confirmations
		// (for example, in devnet where the block confirmations required is lower)
		// To prevent duplicated push in such cases (which can cause unexpected behavior), we need to remove the tx from
		// the block->tx mapping cache
		err = redisStorage.DeleteBlockDeposit(tm.ctx, deposit)
		if err != nil {
			log.Errorf("failed to delete deposit %d from block num cache, error: %v", deposit.DepositCount, err)
			tm.rollbackStore(dbTx)
			return err
		}

		err = tm.storage.Commit(tm.ctx, dbTx)
		if err != nil {
			log.Errorf("AddClaimTx committing dbTx. Err: %v", err)
			rollbackErr := tm.storage.Rollback(tm.ctx, dbTx)
			if rollbackErr != nil {
				log.Fatalf("claimtxman error rolling back state. RollbackErr: %s, err: %s", rollbackErr.Error(), err.Error())
				return rollbackErr
			}
			log.Fatalf("AddClaimTx committing dbTx, err: %s", err.Error())
			return err
		}
		log.Infof("add claim tx for the deposit %d blockID %d successfully", deposit.DepositCount, deposit.BlockID)

		// Notify FE that tx is pending auto claim
		go tm.pushTransactionUpdate(deposit, uint32(pb.TransactionStatus_TX_PENDING_AUTO_CLAIM))
		// Record order waiting time metric
		metrics.RecordOrderWaitTime(uint32(deposit.NetworkID), uint32(deposit.DestinationNetwork), time.Since(deposit.Time))
	}
	return nil
}

func (tm *ClaimTxManager) processDepositStatusXLayer(ger *etherman.GlobalExitRoot, dbTx interface{}) error {
	if ger.BlockID != 0 { // L2 exit root is updated
		log.Infof("Rollup exitroot %v is updated", ger.ExitRoots[1])
		deposits, err := tm.storage.UpdateL2DepositsStatusXLayer(tm.ctx, ger.ExitRoots[1][:], uint(tm.rollupID), uint(tm.l2NetworkID), dbTx)
		if err != nil {
			log.Errorf("error getting and updating L2DepositsStatus. Error: %v", err)
			return err
		}
		log.Debugf("begin send deposits for l1 ready_claim, blockId: %v, blockNumber: %v, deposit size: %v", ger.BlockID, ger.BlockNumber,
			len(deposits))
		for _, deposit := range deposits {
			// Notify FE that tx is pending auto claim
			go tm.pushTransactionUpdate(deposit, uint32(pb.TransactionStatus_TX_PENDING_USER_CLAIM))
			// Record order waiting time metric
			metrics.RecordOrderWaitTime(uint32(deposit.NetworkID), uint32(deposit.DestinationNetwork), time.Since(deposit.Time))
		}
	} else { // L1 exit root is updated in the trusted state
		log.Infof("Mainnet exitroot %v is updated", ger.ExitRoots[0])
		deposits, err := tm.storage.UpdateL1DepositsStatusXLayer(tm.ctx, ger.ExitRoots[0][:], dbTx)
		if err != nil {
			log.Errorf("error getting and updating L1DepositsStatus. Error: %v", err)
			return err
		}
		log.Debugf("Mainnet deposits count %d", len(deposits))
		for _, deposit := range deposits {
			if tm.l2NetworkID != deposit.DestinationNetwork {
				log.Infof("Ignoring deposit: %d: dest_net: %d, we are:%d", deposit.DepositCount, deposit.DestinationNetwork, tm.l2NetworkID)
				continue
			}

			claimHash, err := tm.bridgeService.GetDepositStatus(tm.ctx, deposit.DepositCount, deposit.NetworkID, deposit.DestinationNetwork)
			if err != nil {
				log.Errorf("error getting deposit status for deposit %d. Error: %v", deposit.DepositCount, err)
				return err
			}
			if len(claimHash) > 0 || (deposit.LeafType == LeafTypeMessage && !tm.isDepositMessageAllowed(deposit)) {
				log.Infof("Ignoring deposit: %d, leafType: %d, claimHash: %s, deposit.OriginalAddress: %s", deposit.DepositCount, deposit.LeafType, claimHash, deposit.OriginalAddress.String())
				continue
			}
			log.Infof("create the claim tx for the deposit %d", deposit.DepositCount)
			ger, proof, rollupProof, err := tm.bridgeService.GetClaimProof(deposit.DepositCount, deposit.NetworkID, dbTx)
			if err != nil {
				log.Errorf("error getting Claim Proof for deposit %d. Error: %v", deposit.DepositCount, err)
				return err
			}
			log.Debugf("get claim proof done for the deposit %d", deposit.DepositCount)
			var (
				mtProof       [mtHeight][keyLen]byte
				mtRollupProof [mtHeight][keyLen]byte
			)
			for i := 0; i < mtHeight; i++ {
				mtProof[i] = proof[i]
				mtRollupProof[i] = rollupProof[i]
			}
			tx, err := tm.l2Node.BuildSendClaimXLayer(tm.ctx, deposit, mtProof, mtRollupProof,
				&etherman.GlobalExitRoot{
					ExitRoots: []common.Hash{
						ger.ExitRoots[0],
						ger.ExitRoots[1],
					}}, 1, 1, 1, uint(tm.rollupID),
				tm.auth)
			if err != nil {
				log.Errorf("error BuildSendClaimXLayer tx for deposit %d. Error: %v", deposit.DepositCount, err)
				return err
			}
			log.Debugf("claimTx for deposit %d build successfully %d", deposit.DepositCount)
			if err = tm.addClaimTxXLayer(uint(deposit.DepositCount), tm.auth.From, tx.To(), nil, tx.Data(), dbTx); err != nil {
				log.Errorf("error adding claim tx for deposit %d. Error: %v", deposit.DepositCount, err)
				return err
			}
			log.Debugf("claimTx for deposit %d save successfully %d", deposit.DepositCount)

			// There can be cases that the deposit can be ready for claim (and even claimed) before it reached 64 block confirmations
			// (for example, in devnet where the block confirmations required is lower)
			// To prevent duplicated push in such cases (which can cause unexpected behavior), we need to remove the tx from
			// the block->tx mapping cache
			err = redisStorage.DeleteBlockDeposit(tm.ctx, deposit)
			if err != nil {
				log.Errorf("failed to delete deposit %d from block num cache, error: %v", deposit.DepositCount, err)
				return err
			}

			// Notify FE that tx is pending auto claim
			go tm.pushTransactionUpdate(deposit, uint32(pb.TransactionStatus_TX_PENDING_AUTO_CLAIM))
			// Record order waiting time metric
			metrics.RecordOrderWaitTime(uint32(deposit.NetworkID), uint32(deposit.DestinationNetwork), time.Since(deposit.Time))
		}
	}
	return nil
}

// ReviewMonitoredTxXLayer checks if tx needs to be updated
// accordingly to the current information stored and the current
// state of the blockchain
func (tm *ClaimTxManager) ReviewMonitoredTxXLayer(ctx context.Context, mTx *ctmtypes.MonitoredTx) error {
	mTxLog := log.WithFields("monitoredTx", mTx.DepositID)
	mTxLog.Debug("reviewing")
	// get gas
	tx := ethereum.CallMsg{
		From:  mTx.From,
		To:    mTx.To,
		Value: mTx.Value,
		Data:  mTx.Data,
	}
	gas, err := tm.l2Node.EstimateGas(ctx, tx)
	for i := 1; err != nil && err.Error() != ErrExecutionReverted.Error() && i < tm.cfg.RetryNumber; i++ {
		mTxLog.Warnf("error during gas estimation. Retrying... Error: %v, Data: %s", err, common.Bytes2Hex(tx.Data))
		time.Sleep(tm.cfg.RetryInterval.Duration)
		gas, err = tm.l2Node.EstimateGas(tm.ctx, tx)
	}
	if err != nil {
		err := fmt.Errorf("failed to estimate gas. Error: %v, Data: %s", err, common.Bytes2Hex(tx.Data))
		mTxLog.Errorf("error: %s", err.Error())
		return err
	}

	// check gas
	if gas > mTx.Gas {
		mTxLog.Infof("monitored tx gas updated from %v to %v", mTx.Gas, gas)
		mTx.Gas = gas
	}

	return nil
}

func (tm *ClaimTxManager) addClaimTxXLayer(depositCount uint, from common.Address, to *common.Address, value *big.Int, data []byte, dbTx interface{}) error {
	// get gas
	tx := ethereum.CallMsg{
		From:  from,
		To:    to,
		Value: value,
		Data:  data,
	}
	log.Debugf("addClaimTx deposit: %d", depositCount)
	gas, err := tm.l2Node.EstimateGas(tm.ctx, tx)
	for i := 1; err != nil && err.Error() != ErrExecutionReverted.Error() && i < tm.cfg.RetryNumber; i++ {
		log.Warnf("error while doing gas estimation. Retrying... Error: %v, Data: %s", err, common.Bytes2Hex(data))
		time.Sleep(tm.cfg.RetryInterval.Duration)
		gas, err = tm.l2Node.EstimateGas(tm.ctx, tx)
	}
	if err != nil {
		log.Errorf("failed to estimate gas. Ignoring tx... Error: %v, data: %s", err, common.Bytes2Hex(data))
		return nil
	}

	// create monitored tx
	mTx := ctmtypes.MonitoredTx{
		DepositID: uint64(depositCount), From: from, To: to,
		Value: value, Data: data,
		Gas: gas, Status: ctmtypes.MonitoredTxStatusCreated,
	}

	// add to storage
	err = tm.storage.AddClaimTx(tm.ctx, mTx, dbTx)
	if err != nil {
		err := fmt.Errorf("failed to add tx to get monitored: %v", err)
		log.Errorf("error adding claim tx to db. Error: %s", err.Error())
		return err
	}
	log.Debugf("addClaimTx successfully depositCount: %d", depositCount)

	return nil
}

// Push message to FE to notify about tx status change
func (tm *ClaimTxManager) pushTransactionUpdate(deposit *etherman.Deposit, status uint32) {
	if messagePushProducer == nil {
		log.Errorf("kafka push producer is nil, so can't push tx status change msg!")
		return
	}
	if deposit.LeafType != uint8(utils.LeafTypeAsset) && !tm.isDepositMessageAllowed(deposit) {
		log.Infof("transaction is not asset, so skip push update change, hash: %v", deposit.TxHash)
		return
	}
	estimateTime := uint64(0)
	if deposit.NetworkID != 0 {
		estimateTime = pushtask.GetAvgVerifyDuration(tm.ctx, redisStorage)
	}
	err := messagePushProducer.PushTransactionUpdate(&pb.Transaction{
		FromChain:    uint32(deposit.NetworkID),
		ToChain:      uint32(deposit.DestinationNetwork),
		TxHash:       deposit.TxHash.String(),
		Index:        uint64(deposit.DepositCount),
		Status:       status,
		DestAddr:     deposit.DestinationAddress.Hex(),
		EstimateTime: uint32(estimateTime),
		GlobalIndex:  etherman.GenerateGlobalIndex(false, tm.rollupID-1, deposit.DepositCount).String(),
	})
	if err != nil {
		log.Errorf("PushTransactionUpdate error: %v", err)
	}
}

// setTxNonce get the next nonce from the nonce cache and set it to the tx
func (tm *ClaimTxManager) setTxNonce(mTx *ctmtypes.MonitoredTx) error {
	var nonce uint64
	var err error
	if nonce, err = tm.nonceCache.GetNextNonce(mTx.From); err != nil {
		return fmt.Errorf("getNextNonce failed = %v", err)
	}
	mTx.Nonce = nonce
	return nil
}

func (tm *ClaimTxManager) getDeposits(ger *etherman.GlobalExitRoot) ([]*etherman.Deposit, error) {
	log.Infof("Mainnet exitroot %v is updated", ger.ExitRoots[0])
	deposits, err := tm.storage.GetL1Deposits(tm.ctx, ger.ExitRoots[0][:], nil)
	if err != nil {
		log.Errorf("error processing ger. Error: %v", err)
		return nil, err
	}
	return deposits, nil
}

func (tm *ClaimTxManager) rollbackStore(dbTx interface{}) {
	if rollbackErr := tm.storage.Rollback(tm.ctx, dbTx); rollbackErr != nil {
		log.Errorf("claimtxman error rolling back state. RollbackErr: %v", rollbackErr)
	}
}
