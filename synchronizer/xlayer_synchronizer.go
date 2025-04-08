package synchronizer

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/0xPolygonHermez/zkevm-bridge-service/bridgectrl/pb"
	"github.com/0xPolygonHermez/zkevm-bridge-service/etherman"
	"github.com/0xPolygonHermez/zkevm-bridge-service/log"
	"github.com/0xPolygonHermez/zkevm-bridge-service/utils"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/estimatetime"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/messagepush"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/metrics"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/pushtask"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/redisstorage"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/tokenlogoinfo"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/utils/messagebridge"

	xlUtils "github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/utils"
)

var (
	// Related to filtering large transactions
	largeTxUsdLimit uint64 = 100000
)

func (s *ClientSynchronizer) SetLargeTxUsdLimit(limit uint64) *ClientSynchronizer {
	largeTxUsdLimit = limit
	log.Info("Synchronizer.LargeTxUsdLimit = ", largeTxUsdLimit)
	return s
}

func (s *ClientSynchronizer) beforeProcessDeposit(deposit *etherman.Deposit) {
	messagebridge.ReplaceDepositDestAddresses(deposit)
}

func (s *ClientSynchronizer) afterProcessDeposit(deposit *etherman.Deposit, depositID uint64, dbTx pgx.Tx) error {
	// Add the deposit to Redis for L1
	if deposit.NetworkID == 0 {
		err := s.redisStorage.AddBlockDeposit(s.ctx, deposit)
		if err != nil {
			log.Errorf("networkID: %d, failed to add block deposit to Redis, BlockNumber: %d, Deposit: %+v, err: %s", s.networkID, deposit.BlockNumber, deposit, err)
			rollbackErr := s.storage.Rollback(s.ctx, dbTx)
			if rollbackErr != nil {
				log.Errorf("networkID: %d, error rolling back state to store block. BlockNumber: %v, rollbackErr: %v, err: %s",
					s.networkID, deposit.BlockNumber, rollbackErr, err.Error())
				return rollbackErr
			}
			return err
		}
	}

	err := s.processWstETHDeposit(deposit, dbTx)
	if err != nil {
		log.Errorf("networkID: %d, failed to process wstETH deposit, BlockNumber: %d, Deposit: %+v, err: %s", s.networkID, deposit.BlockNumber, deposit, err)
		rollbackErr := s.storage.Rollback(s.ctx, dbTx)
		if rollbackErr != nil {
			log.Errorf("networkID: %d, error rolling back state to store block. BlockNumber: %v, rollbackErr: %v, err: %s",
				s.networkID, deposit.BlockNumber, rollbackErr, err.Error())
			return rollbackErr
		}
		return err
	}

	// Original address is needed for message allow list check, but it may be changed when we replace USDC info
	origAddress := deposit.OriginalAddress
	// Replace the USDC info here so that the metrics can report the correct token info
	messagebridge.ReplaceDepositInfo(deposit, true)

	// Notify FE about a new deposit
	go func() {
		if s.messagePushProducer == nil {
			log.Errorf("kafka push producer is nil, so can't push tx status change msg!")
			return
		}
		if deposit.LeafType != uint8(utils.LeafTypeAsset) {
			if !messagebridge.IsAllowedContractAddress(origAddress) {
				log.Infof("transaction is not asset, so skip push update change, hash: %v", deposit.TxHash)
				return
			}
		}
		rollupWorkId := xlUtils.GetRollupNetworkId()
		if uint(deposit.NetworkID) != rollupWorkId && uint(deposit.DestinationNetwork) != rollupWorkId {
			log.Infof("transaction is not x layer, so skip push msg and filter large tx, hash: %v", deposit.TxHash)
			return
		}

		transaction := &pb.Transaction{
			FromChain:       uint32(deposit.NetworkID),
			ToChain:         uint32(deposit.DestinationNetwork),
			BridgeToken:     deposit.OriginalAddress.Hex(),
			TokenAmount:     deposit.Amount.String(),
			EstimateTime:    s.getEstimateTimeForDepositCreated(uint(deposit.NetworkID)),
			Time:            uint64(deposit.Time.UnixMilli()),
			TxHash:          deposit.TxHash.String(),
			Id:              depositID,
			Index:           uint64(deposit.DepositCount),
			Status:          uint32(pb.TransactionStatus_TX_CREATED),
			BlockNumber:     deposit.BlockNumber,
			DestAddr:        deposit.DestinationAddress.Hex(),
			FromChainId:     xlUtils.GetChainIdByNetworkId(uint(deposit.NetworkID)),
			ToChainId:       xlUtils.GetChainIdByNetworkId(uint(deposit.DestinationNetwork)),
			GlobalIndex:     s.getGlobalIndex(deposit).String(),
			LeafType:        uint32(deposit.LeafType),
			OriginalNetwork: uint32(deposit.OriginalNetwork),
		}
		transactionMap := make(map[string][]*pb.Transaction)
		chainId := xlUtils.GetChainIdByNetworkId(uint(deposit.OriginalNetwork))
		logoCacheKey := tokenlogoinfo.GetTokenLogoMapKey(transaction.GetBridgeToken(), chainId)
		transactionMap[logoCacheKey] = append(transactionMap[logoCacheKey], transaction)
		tokenlogoinfo.FillLogoInfos(s.ctx, s.redisStorage, transactionMap)
		err := s.messagePushProducer.PushTransactionUpdate(transaction)
		if err != nil {
			log.Errorf("PushTransactionUpdate error: %v", err)
		}
		// filter and cache large transactions
		s.filterLargeTransaction(s.ctx, transaction, uint(chainId))
	}()

	metrics.RecordOrder(uint32(deposit.NetworkID), uint32(deposit.DestinationNetwork), uint32(deposit.LeafType), uint32(deposit.OriginalNetwork), deposit.OriginalAddress, deposit.Amount)
	return nil
}

func (s *ClientSynchronizer) filterLargeTransaction(ctx context.Context, transaction *pb.Transaction, chainId uint) {
	if transaction.LogoInfo == nil {
		log.Infof("failed to get logo info, so skip filter large transaction, tx: %v", transaction.GetTxHash())
		return
	}
	symbolInfo := &pb.SymbolInfo{
		ChainId: uint64(chainId),
		Address: transaction.BridgeToken,
	}
	priceInfos, err := s.redisStorage.GetCoinPrice(ctx, []*pb.SymbolInfo{symbolInfo})
	if err != nil || len(priceInfos) == 0 {
		log.Errorf("not find coin price for coin: %v, chain: %v, so skip monitor large tx: %v", symbolInfo.Address, symbolInfo.ChainId, transaction.GetTxHash())
		return
	}
	originNum, success := new(big.Float).SetPrec(uint(transaction.GetLogoInfo().Decimal)).SetString(transaction.GetTokenAmount())
	if !success {
		log.Errorf("failed to convert token num to big.Float, so skip monitor large tx, tx: %v, token num: %v", transaction.GetTxHash(), transaction.GetTokenAmount())
		return
	}
	tokenDecimal := new(big.Float).SetPrec(uint(transaction.GetLogoInfo().Decimal)).SetFloat64(math.Pow10(int(transaction.GetLogoInfo().Decimal)))
	tokenAmount, _ := new(big.Float).Quo(originNum, tokenDecimal).Float64()
	usdAmount := priceInfos[0].Price * tokenAmount
	if usdAmount < float64(largeTxUsdLimit) {
		log.Infof("tx usd amount less than limit, so skip, tx usd amount: %v, tx: %v", usdAmount, transaction.GetTxHash())
		return
	}
	s.freshLargeTxCache(ctx, transaction, chainId, tokenAmount, usdAmount)
}

func (s *ClientSynchronizer) freshLargeTxCache(ctx context.Context, transaction *pb.Transaction, chainId uint, tokenAmount float64, usdAmount float64) {
	largeTxInfo := &pb.LargeTxInfo{
		ChainId:   uint64(chainId),
		Symbol:    transaction.LogoInfo.Symbol,
		Amount:    tokenAmount,
		UsdAmount: usdAmount,
		Hash:      transaction.TxHash,
		Address:   transaction.DestAddr,
	}
	key := utils.GetLargeTxRedisKeySuffix(uint(transaction.ToChain), utils.OpWrite)
	size, err := s.redisStorage.AddLargeTransaction(ctx, key, largeTxInfo)
	if err != nil {
		log.Errorf("failed set large tx cache for tx: %v, err: %v", transaction.GetTxHash(), err)
	}
	log.Debugf("success push tx for key: %v, size: %v", key, size)
	if size == 1 {
		log.Infof("success init new cache list for large transaction, key: %v", key)
		ret, err := s.redisStorage.ExpireLargeTransactions(ctx, key, utils.GetLargeTxCacheExpireDuration())
		if err != nil || !ret {
			log.Errorf("failed to expire large tx key: %v, err: %v", key, err)
			return
		}
		log.Infof("success expire large tx key: %v", key)
	}
}

func (s *ClientSynchronizer) getEstimateTimeForDepositCreated(networkId uint) uint32 {
	if networkId == 0 {
		return estimatetime.GetDefaultCalculator().Get(networkId)
	}
	return uint32(pushtask.GetAvgCommitDuration(s.ctx, s.redisStorage))
}

func (s *ClientSynchronizer) afterProcessClaim(claim *etherman.Claim, dbTx pgx.Tx) {
	// Try to retrieve deposit transaction info
	deposit, err := s.storage.GetDeposit(s.ctx, claim.Index, claim.OriginalNetwork, nil)
	if err != nil || deposit == nil {
		log.Warnf("failed to get deposit for claim, claim: %+v, err: %v", claim, err)
		return
	}

	err = s.processWstETHClaim(deposit, dbTx)
	if err != nil {
		log.Warnf("failed to process wstETH claim, claim: %+v, err: %v", claim, err)
		return
	}

	// Notify FE that the tx has been claimed
	go func() {
		if s.messagePushProducer == nil {
			log.Errorf("kafka push producer is nil, so can't push tx status change msg!")
			return
		}

		if deposit.LeafType != uint8(utils.LeafTypeAsset) {
			if !messagebridge.IsAllowedContractAddress(deposit.OriginalAddress) {
				log.Infof("transaction is not asset, so skip push update change, hash: %v", deposit.TxHash)
				return
			}
		}
		err = s.messagePushProducer.PushTransactionUpdate(&pb.Transaction{
			FromChain:   uint32(deposit.NetworkID),
			ToChain:     uint32(deposit.DestinationNetwork),
			TxHash:      deposit.TxHash.String(),
			Index:       uint64(deposit.DepositCount),
			Status:      uint32(pb.TransactionStatus_TX_CLAIMED),
			ClaimTxHash: claim.TxHash.Hex(),
			ClaimTime:   uint64(claim.Time.UnixMilli()),
			DestAddr:    deposit.DestinationAddress.Hex(),
			GlobalIndex: s.getGlobalIndex(deposit).String(),
		})
		if err != nil {
			log.Errorf("PushTransactionUpdate error: %v", err)
		}
	}()
}

func (s *ClientSynchronizer) getGlobalIndex(deposit *etherman.Deposit) *big.Int {
	isMainnet := deposit.NetworkID == 0
	rollupIndex := s.rollupID - 1
	return etherman.GenerateGlobalIndex(isMainnet, uint32(rollupIndex), deposit.DepositCount)
}

// recordLatestBlockNum continuously records the latest block number to prometheus metrics
func (s *ClientSynchronizer) RecordLatestBlockNum() {
	log.Debugf("Start recordLatestBlockNum")
	ticker := time.NewTicker(2 * time.Second) //nolint:gomnd

	for range ticker.C {
		// Get the latest block header
		header, err := s.etherMan.HeaderByNumber(s.ctx, nil)
		if err != nil {
			log.Errorf("HeaderByNumber err: %v", err)
			continue
		}
		metrics.RecordLatestBlockNum(uint32(s.networkID), header.Number.Uint64())
	}
}

// processWstETHDeposit increases the l2TokenNotWithdrawn value for wstETH bridge, used for reconciliation purpose
func (s *ClientSynchronizer) processWstETHDeposit(deposit *etherman.Deposit, dbTx pgx.Tx) error {
	amount := deposit.Amount
	processor := messagebridge.GetProcessorByType(messagebridge.WstETH)
	if processor == nil || !processor.CheckContractAddress(deposit.OriginalAddress) {
		return nil
	}
	_, amount = processor.DecodeMetadataFn(deposit.Metadata)
	return s.processWstETHCommon(deposit, func(value *big.Int) {
		value.Add(value, amount)
	}, dbTx)
}

// processWstETHClaim decrease the l2TokenNotWithdrawn value for wstETH bridge, used for reconciliation purpose
func (s *ClientSynchronizer) processWstETHClaim(deposit *etherman.Deposit, dbTx pgx.Tx) error {
	amount := deposit.Amount
	processor := messagebridge.GetProcessorByType(messagebridge.WstETH)
	if processor == nil || !processor.CheckContractAddress(deposit.OriginalAddress) {
		return nil
	}
	_, amount = processor.DecodeMetadataFn(deposit.Metadata)
	return s.processWstETHCommon(deposit, func(value *big.Int) {
		value.Sub(value, amount)
	}, dbTx)
}

func (s *ClientSynchronizer) processWstETHCommon(deposit *etherman.Deposit, valueUpdateFn func(*big.Int), dbTx pgx.Tx) error {
	processor := messagebridge.GetProcessorByType(messagebridge.WstETH)
	if processor == nil {
		return nil
	}
	// Check if this deposit is for wstETH
	if !processor.CheckContractAddress(deposit.OriginalAddress) {
		return nil
	}

	// Update DB using the token original address
	tokenAddr := processor.GetTokenAddressList()[0]
	value, err := s.storage.GetBridgeBalance(s.ctx, tokenAddr, uint(deposit.NetworkID), true, dbTx)
	if err != nil {
		return errors.Wrap(err, "GetBridgeBalance from DB err")
	}
	// Update the value
	valueUpdateFn(value)
	log.Debugf("setting wstETH bridgeBalance to %v, network_id = %v", value.String(), deposit.NetworkID)
	err = s.storage.SetBridgeBalance(s.ctx, tokenAddr, uint(deposit.NetworkID), value, dbTx)
	return errors.Wrap(err, "SetBridgeBalance to DB err")
}

func (s *ClientSynchronizer) SetProducer(messagePushProducer messagepush.KafkaProducer) *ClientSynchronizer {
	s.messagePushProducer = messagePushProducer
	return s
}

func (s *ClientSynchronizer) SetRedis(rs redisstorage.RedisStorage) *ClientSynchronizer {
	s.redisStorage = rs
	return s
}

func (s *ClientSynchronizer) SetRollupID(rollupID uint) *ClientSynchronizer {
	s.rollupID = rollupID
	return s
}
