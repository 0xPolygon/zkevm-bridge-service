package server

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"

	"github.com/0xPolygonHermez/zkevm-bridge-service/bridgectrl"
	"github.com/0xPolygonHermez/zkevm-bridge-service/bridgectrl/pb"
	"github.com/0xPolygonHermez/zkevm-bridge-service/log"
	"github.com/0xPolygonHermez/zkevm-bridge-service/utils"
	"github.com/0xPolygonHermez/zkevm-bridge-service/utils/gerror"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/estimatetime"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/localcache"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/messagepush"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/pushtask"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/redisstorage"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/tokenlogoinfo"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/utils/messagebridge"

	ctmtypes "github.com/0xPolygonHermez/zkevm-bridge-service/claimtxman/types"
	xlUtils "github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/utils"
)

const (
	mtHeight           = 32 // For sending mtProof to bridge contract, it requires constant-sized array...
	defaultMinDuration = 1
)

const (
	minReadyTimeLimitForWaitClaimSeconds int64 = 24 * 60 * 1000
)

// Put in global variables (separate from struct fields)
// NOTE: Beware of concurrency issues as gRPC handlers are run inside go routines.
var (
	rsstore             redisstorage.RedisStorage
	mainCoinsCache      localcache.MainCoinsCache
	messagePushProducer messagepush.KafkaProducer

	// Read-only variables. No need thread-safe.
	// Initialized once on creation.
	nodeClientsMap map[uint]*utils.Client
	authMap        map[uint]*bind.TransactOpts
)

func (s *bridgeService) WithRedisStorage(storage redisstorage.RedisStorage) *bridgeService {
	rsstore = storage
	return s
}

func (s *bridgeService) WithMainCoinsCache(cache localcache.MainCoinsCache) *bridgeService {
	mainCoinsCache = cache
	return s
}

func (s *bridgeService) WithMessagePushProducer(producer messagepush.KafkaProducer) *bridgeService {
	messagePushProducer = producer
	return s
}

func (s *bridgeService) LogConfig() {
	// Currently logs apollo loaded configs as a sanity check
	log.Info("BridgeServer.DefaultPageLimit = ", s.defaultPageLimit)
	log.Info("BridgeServer.MaxPageLimit = ", s.maxPageLimit)
}

func (s *bridgeService) WithL2Clients(
	l2NodeClients []*utils.Client,
	l2Auths []*bind.TransactOpts,
	networks []uint32,
) *bridgeService {
	nodeClientsMap = make(map[uint]*utils.Client, len(networks))
	authMap = make(map[uint]*bind.TransactOpts, len(networks))
	for i, network := range networks {
		if i > 0 { // Why skip first index?
			id := uint(network) // nolint:gosec
			nodeClientsMap[id] = l2NodeClients[i-1]
			authMap[id] = l2Auths[i-1]
		}
	}
	return s
}

func (s *bridgeService) GetSmtProof(ctx context.Context, req *pb.GetSmtProofRequest) (*pb.CommonProofResponse, error) {
	globalExitRoot, merkleProof, rollupMerkleProof, err := s.GetClaimProof(uint32(req.Index), req.FromChain, nil) // nolint:gosec
	if err != nil || len(merkleProof) != len(rollupMerkleProof) {
		log.Errorf("GetSmtProof err[%v] merkleProofLen[%v] rollupMerkleProofLen[%v]", err, len(merkleProof), len(rollupMerkleProof))
		return &pb.CommonProofResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
			Msg:  err.Error(),
		}, nil
	}
	var (
		proof       []string
		rollupProof []string
	)
	for i := 0; i < len(merkleProof); i++ {
		proof = append(proof, "0x"+hex.EncodeToString(merkleProof[i][:]))
		rollupProof = append(rollupProof, "0x"+hex.EncodeToString(rollupMerkleProof[i][:]))
	}

	return &pb.CommonProofResponse{
		Code: uint32(pb.ErrorCode_ERROR_OK),
		Data: &pb.ProofDetail{
			SmtProof:        proof,
			RollupSmtProof:  rollupProof,
			MainnetExitRoot: globalExitRoot.ExitRoots[0].Hex(),
			RollupExitRoot:  globalExitRoot.ExitRoots[1].Hex(),
		},
	}, nil
}

// GetCoinPrice returns the price for each coin symbol in the request
// Bridge rest API endpoint
func (s *bridgeService) GetCoinPrice(ctx context.Context, req *pb.GetCoinPriceRequest) (*pb.CommonCoinPricesResponse, error) {
	// convert inner chainId to standard chain id
	for _, symbol := range req.SymbolInfos {
		symbol.ChainId = xlUtils.GetStandardChainIdByInnerId(symbol.ChainId)
	}
	priceList, err := rsstore.GetCoinPrice(ctx, req.SymbolInfos)
	if err != nil {
		log.Errorf("get coin price from redis failed for symbol: %v, error: %v", req.SymbolInfos, err)
		return &pb.CommonCoinPricesResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
			Msg:  gerror.ErrInternalErrorForRpcCall.Error(),
		}, nil
	}
	// convert standard chainId to ok inner chainId
	for _, priceInfo := range priceList {
		priceInfo.ChainId = xlUtils.GetInnerChainIdByStandardId(priceInfo.ChainId)
	}
	return &pb.CommonCoinPricesResponse{
		Code: uint32(pb.ErrorCode_ERROR_OK),
		Data: priceList,
	}, nil
}

// GetMainCoins returns the info of the main coins in a network
// Bridge rest API endpoint
func (s *bridgeService) GetMainCoins(ctx context.Context, req *pb.GetMainCoinsRequest) (*pb.CommonCoinsResponse, error) {
	coins, err := mainCoinsCache.GetMainCoinsByNetwork(ctx, req.NetworkId)
	if err != nil {
		log.Errorf("get main coins from cache failed for net: %v, error: %v", req.NetworkId, err)
		return &pb.CommonCoinsResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
			Msg:  gerror.ErrInternalErrorForRpcCall.Error(),
		}, nil
	}
	// use ok inner chain id
	for _, coinInfo := range coins {
		coinInfo.ChainId = xlUtils.GetInnerChainIdByStandardId(coinInfo.ChainId)
	}
	return &pb.CommonCoinsResponse{
		Code: uint32(pb.ErrorCode_ERROR_OK),
		Data: coins,
	}, nil
}

// GetPendingTransactions returns the pending transactions of an account
// Bridge rest API endpoint
func (s *bridgeService) GetPendingTransactions(ctx context.Context, req *pb.GetPendingTransactionsRequest) (*pb.CommonTransactionsResponse, error) {
	limit := req.Limit
	if limit == 0 {
		limit = s.defaultPageLimit
	}
	if limit > s.maxPageLimit {
		limit = s.maxPageLimit
	}

	deposits, err := s.storage.GetPendingTransactions(ctx, req.DestAddr, uint(limit+1), uint(req.Offset), messagebridge.GetContractAddressList(), nil)
	if err != nil {
		log.Errorf("get pending tx failed for address: %v, limit: %v, offset: %v, error: %v", req.DestAddr, limit, req.Offset, err)
		return &pb.CommonTransactionsResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
			Msg:  gerror.ErrInternalErrorForRpcCall.Error(),
		}, nil
	}

	hasNext := len(deposits) > int(limit)
	if hasNext {
		deposits = deposits[:limit]
	}

	l1BlockNum, _ := rsstore.GetL1BlockNum(ctx)
	l2CommitBlockNum, _ := rsstore.GetCommitMaxBlockNum(ctx)
	l2AvgCommitDuration := pushtask.GetAvgCommitDuration(ctx, rsstore)
	l2AvgVerifyDuration := pushtask.GetAvgVerifyDuration(ctx, rsstore)
	currTime := time.Now()

	var pbTransactions []*pb.Transaction
	transactionMap := make(map[string][]*pb.Transaction)
	for _, deposit := range deposits {
		// replace contract address to real token address
		messagebridge.ReplaceDepositInfo(deposit, false)
		transaction := xlUtils.EthermanDepositToPbTransaction(deposit)
		transaction.EstimateTime = estimatetime.GetDefaultCalculator().Get(uint(deposit.NetworkID)) // nolint:gosec
		transaction.Status = uint32(pb.TransactionStatus_TX_CREATED)
		transaction.GlobalIndex = s.getGlobalIndex(deposit).String()
		if deposit.ReadyForClaim {
			transaction.Status = uint32(pb.TransactionStatus_TX_PENDING_USER_CLAIM)
			// For L1->L2, if backend is trying to auto-claim, set the status to 0 to block the user from manual-claim
			// When the auto-claim failed, set status to 1 to let the user claim manually through front-end
			if deposit.NetworkID == 0 {
				mTx, err := s.storage.GetClaimTxById(ctx, uint(deposit.DepositCount), nil)
				if err == nil && mTx.Status != ctmtypes.MonitoredTxStatusFailed {
					transaction.Status = uint32(pb.TransactionStatus_TX_PENDING_AUTO_CLAIM)
				}
			}
		} else {
			// For L1->L2, when ready_for_claim is false, but there have been more than 64 block confirmations,
			// should also display the status as "L2 executing" (pending auto claim)
			if deposit.NetworkID == 0 {
				if l1BlockNum-deposit.BlockNumber >= xlUtils.L1TargetBlockConfirmations {
					transaction.Status = uint32(pb.TransactionStatus_TX_PENDING_AUTO_CLAIM)
				}
			} else {
				if l2CommitBlockNum >= deposit.BlockNumber {
					transaction.Status = uint32(pb.TransactionStatus_TX_PENDING_VERIFICATION)
				}
				s.setDurationForL2Deposit(ctx, l2AvgCommitDuration, l2AvgVerifyDuration, currTime, transaction, deposit.Time)
			}
		}
		// chain id convert to ok inner chain id
		if transaction.FromChainId != 0 {
			transaction.FromChainId = uint32(xlUtils.GetInnerChainIdByStandardId(uint64(transaction.FromChainId)))
		}
		if transaction.ToChainId != 0 {
			transaction.ToChainId = uint32(xlUtils.GetInnerChainIdByStandardId(uint64(transaction.ToChainId)))
		}
		pbTransactions = append(pbTransactions, transaction)
		chainId := xlUtils.GetChainIdByNetworkId(uint(deposit.OriginalNetwork)) // nolint:gosec
		logoCacheKey := tokenlogoinfo.GetTokenLogoMapKey(transaction.GetBridgeToken(), chainId)
		transactionMap[logoCacheKey] = append(transactionMap[logoCacheKey], transaction)
	}
	tokenlogoinfo.FillLogoInfos(ctx, rsstore, transactionMap)
	return &pb.CommonTransactionsResponse{
		Code: uint32(pb.ErrorCode_ERROR_OK),
		Data: &pb.TransactionDetail{HasNext: hasNext, Transactions: pbTransactions},
	}, nil
}

// GetAllTransactions returns all the transactions of an account, similar to GetBridges
// Bridge rest API endpoint
func (s *bridgeService) GetAllTransactions(ctx context.Context, req *pb.GetAllTransactionsRequest) (*pb.CommonTransactionsResponse, error) {
	limit := req.Limit
	if limit == 0 {
		limit = s.defaultPageLimit
	}
	if limit > s.maxPageLimit {
		limit = s.maxPageLimit
	}

	deposits, err := s.storage.GetDepositsXLayer(ctx, req.DestAddr, uint(limit+1), uint(req.Offset), messagebridge.GetContractAddressList(), nil)
	if err != nil {
		log.Errorf("get deposits from db failed for address: %v, limit: %v, offset: %v, error: %v", req.DestAddr, limit, req.Offset, err)
		return &pb.CommonTransactionsResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
			Msg:  gerror.ErrInternalErrorForRpcCall.Error(),
		}, nil
	}

	hasNext := len(deposits) > int(limit)
	if hasNext {
		deposits = deposits[0:limit]
	}

	l1BlockNum, _ := rsstore.GetL1BlockNum(ctx)
	l2CommitBlockNum, _ := rsstore.GetCommitMaxBlockNum(ctx)
	l2AvgCommitDuration := pushtask.GetAvgCommitDuration(ctx, rsstore)
	l2AvgVerifyDuration := pushtask.GetAvgVerifyDuration(ctx, rsstore)
	currTime := time.Now()

	var pbTransactions []*pb.Transaction
	transactionMap := make(map[string][]*pb.Transaction)
	for _, deposit := range deposits {
		// replace contract address to real token address
		messagebridge.ReplaceDepositInfo(deposit, false)
		transaction := xlUtils.EthermanDepositToPbTransaction(deposit)
		transaction.EstimateTime = estimatetime.GetDefaultCalculator().Get(uint(deposit.NetworkID)) // nolint:gosec
		transaction.Status = uint32(pb.TransactionStatus_TX_CREATED)                                // Not ready for claim
		transaction.GlobalIndex = s.getGlobalIndex(deposit).String()
		if deposit.ReadyForClaim {
			// Check whether it has been claimed or not
			claim, err := s.storage.GetClaim(ctx, deposit.DepositCount, deposit.NetworkID, deposit.DestinationNetwork, nil)
			transaction.Status = uint32(pb.TransactionStatus_TX_PENDING_USER_CLAIM) // Ready but not claimed
			if err != nil {
				if !errors.Is(err, gerror.ErrStorageNotFound) {
					return &pb.CommonTransactionsResponse{
						Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
						Msg:  errors.Wrap(err, "load claim error").Error(),
					}, nil
				}
				// For L1->L2, if backend is trying to auto-claim, set the status to 0 to block the user from manual-claim
				// When the auto-claim failed, set status to 1 to let the user claim manually through front-end
				if deposit.NetworkID == 0 {
					mTx, err := s.storage.GetClaimTxById(ctx, uint(deposit.DepositCount), nil) // nolint:gosec
					if err == nil && mTx.Status != ctmtypes.MonitoredTxStatusFailed {
						transaction.Status = uint32(pb.TransactionStatus_TX_PENDING_AUTO_CLAIM)
					}
				}
			} else {
				transaction.Status = uint32(pb.TransactionStatus_TX_CLAIMED) // Claimed
				transaction.ClaimTxHash = claim.TxHash.String()
				transaction.ClaimTime = uint64(claim.Time.UnixMilli())
			}
		} else {
			// For L1->L2, when ready_for_claim is false, but there have been more than 64 block confirmations,
			// should also display the status as "L2 executing" (pending auto claim)
			if deposit.NetworkID == 0 {
				if l1BlockNum-deposit.BlockNumber >= xlUtils.L1TargetBlockConfirmations {
					transaction.Status = uint32(pb.TransactionStatus_TX_PENDING_AUTO_CLAIM)
				}
			} else {
				if l2CommitBlockNum >= deposit.BlockNumber {
					transaction.Status = uint32(pb.TransactionStatus_TX_PENDING_VERIFICATION)
				}
				s.setDurationForL2Deposit(ctx, l2AvgCommitDuration, l2AvgVerifyDuration, currTime, transaction, deposit.Time)
			}
		}
		// chain id convert to ok inner chain id
		if transaction.FromChainId != 0 {
			transaction.FromChainId = uint32(xlUtils.GetInnerChainIdByStandardId(uint64(transaction.FromChainId)))
		}
		if transaction.ToChainId != 0 {
			transaction.ToChainId = uint32(xlUtils.GetInnerChainIdByStandardId(uint64(transaction.ToChainId)))
		}
		pbTransactions = append(pbTransactions, transaction)
		chainId := xlUtils.GetChainIdByNetworkId(uint(deposit.OriginalNetwork)) // nolint:gosec
		logoCacheKey := tokenlogoinfo.GetTokenLogoMapKey(transaction.GetBridgeToken(), chainId)
		transactionMap[logoCacheKey] = append(transactionMap[logoCacheKey], transaction)
	}
	tokenlogoinfo.FillLogoInfos(ctx, rsstore, transactionMap)

	return &pb.CommonTransactionsResponse{
		Code: uint32(pb.ErrorCode_ERROR_OK),
		Data: &pb.TransactionDetail{HasNext: hasNext, Transactions: pbTransactions},
	}, nil
}

// GetNotReadyTransactions returns all deposit transactions with ready_for_claim = false
func (s *bridgeService) GetNotReadyTransactions(ctx context.Context, req *pb.GetNotReadyTransactionsRequest) (*pb.CommonTransactionsResponse, error) {
	limit := req.Limit
	if limit == 0 {
		limit = s.defaultPageLimit
	}
	if limit > s.maxPageLimit {
		limit = s.maxPageLimit
	}

	deposits, err := s.storage.GetNotReadyTransactions(ctx, uint(limit+1), uint(req.Offset), nil)
	if err != nil {
		return &pb.CommonTransactionsResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
			Msg:  err.Error(),
		}, nil
	}

	hasNext := len(deposits) > int(limit)
	if hasNext {
		deposits = deposits[0:limit]
	}

	var pbTransactions []*pb.Transaction
	for _, deposit := range deposits {
		transaction := xlUtils.EthermanDepositToPbTransaction(deposit)
		transaction.EstimateTime = estimatetime.GetDefaultCalculator().Get(uint(deposit.NetworkID)) // nolint:gosec
		transaction.Status = uint32(pb.TransactionStatus_TX_CREATED)
		transaction.GlobalIndex = s.getGlobalIndex(deposit).String()
		pbTransactions = append(pbTransactions, transaction)
	}

	return &pb.CommonTransactionsResponse{
		Code: uint32(pb.ErrorCode_ERROR_OK),
		Data: &pb.TransactionDetail{HasNext: hasNext, Transactions: pbTransactions},
	}, nil
}

// GetMonitoredTxsByStatus returns list of monitored transactions, filtered by status
func (s *bridgeService) GetMonitoredTxsByStatus(ctx context.Context, req *pb.GetMonitoredTxsByStatusRequest) (*pb.CommonMonitoredTxsResponse, error) {
	limit := req.Limit
	if limit == 0 {
		limit = s.defaultPageLimit
	}
	if limit > s.maxPageLimit {
		limit = s.maxPageLimit
	}

	mTxs, err := s.storage.GetClaimTxsByStatusWithLimit(ctx, []ctmtypes.MonitoredTxStatus{ctmtypes.MonitoredTxStatus(req.Status)}, uint(limit+1), uint(req.Offset), nil)
	if err != nil {
		return &pb.CommonMonitoredTxsResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
			Msg:  err.Error(),
		}, nil
	}

	hasNext := len(mTxs) > int(limit)
	if hasNext {
		mTxs = mTxs[0:limit]
	}

	var pbTransactions []*pb.MonitoredTx
	for _, mTx := range mTxs {
		transaction := &pb.MonitoredTx{
			Id:        uint64(mTx.DepositID),
			From:      mTx.From.String(),
			To:        mTx.To.String(),
			Nonce:     mTx.Nonce,
			Value:     mTx.Value.String(),
			Data:      "0x" + hex.EncodeToString(mTx.Data),
			Gas:       mTx.Gas,
			GasPrice:  mTx.GasPrice.String(),
			Status:    string(mTx.Status),
			CreatedAt: uint64(mTx.CreatedAt.UnixMilli()),
			UpdatedAt: uint64(mTx.UpdatedAt.UnixMilli()),
		}
		for h := range mTx.History {
			transaction.History = append(transaction.History, h.String())
		}
		pbTransactions = append(pbTransactions, transaction)
	}

	return &pb.CommonMonitoredTxsResponse{
		Code: uint32(pb.ErrorCode_ERROR_OK),
		Data: &pb.MonitoredTxsDetail{HasNext: hasNext, Transactions: pbTransactions},
	}, nil
}

// GetEstimateTime returns the estimated deposit waiting time for L1 and L2
func (s *bridgeService) GetEstimateTime(ctx context.Context, req *pb.GetEstimateTimeRequest) (*pb.CommonEstimateTimeResponse, error) {
	return &pb.CommonEstimateTimeResponse{
		Code: uint32(pb.ErrorCode_ERROR_OK),
		Data: []uint32{estimatetime.GetDefaultCalculator().Get(0), estimatetime.GetDefaultCalculator().Get(1)},
	}, nil
}

// ManualClaim manually sends a claim transaction for a specific deposit
func (s *bridgeService) ManualClaim(ctx context.Context, req *pb.ManualClaimRequest) (*pb.CommonManualClaimResponse, error) {
	// Only allow L1->L2
	if req.FromChain != 0 {
		return &pb.CommonManualClaimResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
			Msg:  "only allow L1->L2 claim",
		}, nil
	}

	// Query the deposit info from storage
	deposit, err := s.storage.GetDepositByHash(ctx, req.DestAddr, uint(req.FromChain), req.DepositTxHash, nil)
	if err != nil {
		log.Errorf("Failed to get deposit: %v", err)
		return &pb.CommonManualClaimResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
			Msg:  "failed to get deposit info",
		}, nil
	}

	// Only allow to claim ready transactions
	if !deposit.ReadyForClaim {
		return &pb.CommonManualClaimResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
			Msg:  "transaction is not ready for claim",
		}, nil
	}

	// Check whether the deposit has already been claimed
	_, err = s.storage.GetClaim(ctx, deposit.DepositCount, deposit.NetworkID, deposit.DestinationNetwork, nil)
	if err == nil {
		return &pb.CommonManualClaimResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
			Msg:  "transaction has already been claimed",
		}, nil
	}
	if !errors.Is(err, gerror.ErrStorageNotFound) {
		return &pb.CommonManualClaimResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
		}, nil
	}

	destNet := deposit.DestinationNetwork
	client, ok := nodeClientsMap[uint(destNet)] // nolint:gosec
	if !ok || client == nil {
		log.Errorf("node client for networkID %v not found", destNet)
		return &pb.CommonManualClaimResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
		}, nil
	}
	// Get the claim proof
	ger, proves, rollupProves, err := s.GetClaimProof(deposit.DepositCount, deposit.NetworkID, nil)
	if err != nil {
		log.Errorf("failed to get claim proof for deposit %v networkID %v: %v", deposit.DepositCount, deposit.NetworkID, err)
	}
	var (
		mtProves       [mtHeight][bridgectrl.KeyLen]byte
		mtRollupProves [mtHeight][bridgectrl.KeyLen]byte
	)
	for i := 0; i < mtHeight; i++ {
		mtProves[i] = proves[i]
		mtRollupProves[i] = rollupProves[i]
	}
	// Send claim transaction to the node
	tx, err := client.SendClaimXLayer(ctx, deposit, mtProves, mtRollupProves, ger, uint(deposit.NetworkID), authMap[uint(destNet)]) // nolint:gosec
	if err != nil {
		log.Errorf("failed to send claim transaction: %v", err)
		return &pb.CommonManualClaimResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
			Msg:  "failed to send claim transaction",
		}, nil
	}

	return &pb.CommonManualClaimResponse{
		Code: uint32(pb.ErrorCode_ERROR_OK),
		Data: &pb.ManualClaimResponse{
			ClaimTxHash: tx.Hash().String(),
		},
	}, nil
}

// GetReadyPendingTransactions returns all transactions from a network which are ready_for_claim but not claimed
func (s *bridgeService) GetReadyPendingTransactions(ctx context.Context, req *pb.GetReadyPendingTransactionsRequest) (*pb.CommonTransactionsResponse, error) {
	limit := req.Limit
	if limit == 0 {
		limit = s.defaultPageLimit
	}
	if limit > s.maxPageLimit {
		limit = s.maxPageLimit
	}

	minReadyTime := time.Now().Add(time.Duration(-minReadyTimeLimitForWaitClaimSeconds) * time.Second)

	deposits, err := s.storage.GetReadyPendingTransactions(ctx, uint(req.NetworkId), uint(limit+1), uint(req.Offset), minReadyTime, nil)
	if err != nil {
		return &pb.CommonTransactionsResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
		}, nil
	}

	hasNext := len(deposits) > int(limit)
	if hasNext {
		deposits = deposits[:limit]
	}

	var pbTransactions []*pb.Transaction
	for _, deposit := range deposits {
		transaction := xlUtils.EthermanDepositToPbTransaction(deposit)
		transaction.EstimateTime = estimatetime.GetDefaultCalculator().Get(uint(deposit.NetworkID))
		transaction.Status = uint32(pb.TransactionStatus_TX_PENDING_USER_CLAIM)
		transaction.GlobalIndex = s.getGlobalIndex(deposit).String()
		pbTransactions = append(pbTransactions, transaction)
	}

	return &pb.CommonTransactionsResponse{
		Code: uint32(pb.ErrorCode_ERROR_OK),
		Data: &pb.TransactionDetail{HasNext: hasNext, Transactions: pbTransactions},
	}, nil
}

func (s *bridgeService) GetLargeTransactionInfos(ctx context.Context, req *pb.LargeTxsRequest) (*pb.LargeTxsResponse, error) {
	key := utils.GetLargeTxRedisKeySuffix(uint(req.NetworkId), utils.OpRead)
	txInfos, err := rsstore.GetLargeTransactions(ctx, key)
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Errorf("failed to get large tx cache for key: %v", key)
		return &pb.LargeTxsResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
			Msg:  "failed to get cache",
		}, nil
	}
	return &pb.LargeTxsResponse{Code: uint32(pb.ErrorCode_ERROR_OK), Data: txInfos}, nil
}

func (s *bridgeService) GetWstEthTokenNotWithdrawn(ctx context.Context, req *pb.GetWstEthTokenNotWithdrawnRequest) (*pb.GetWstEthTokenNotWithdrawnResponse, error) {
	processor := messagebridge.GetProcessorByType(messagebridge.WstETH)
	if processor == nil {
		return &pb.GetWstEthTokenNotWithdrawnResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
			Msg:  "internal: wstETH processor is not inited",
		}, nil
	}
	tokenAddr := processor.GetTokenAddressList()[0]
	valueL1, errL1 := s.storage.GetBridgeBalance(ctx, tokenAddr, xlUtils.GetMainNetworkId(), false, nil)
	valueL2, errL2 := s.storage.GetBridgeBalance(ctx, tokenAddr, xlUtils.GetRollupNetworkId(), false, nil)
	if errL1 != nil || errL2 != nil {
		log.Errorf("failed to get wstETH TokenNotWithdrawn, errL1: %v, errL2: %v", errL1, errL2)
		return &pb.GetWstEthTokenNotWithdrawnResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
			Msg:  "failed to get from DB",
		}, nil
	}
	return &pb.GetWstEthTokenNotWithdrawnResponse{
		Code: uint32(pb.ErrorCode_ERROR_OK),
		Data: []string{valueL1.String(), valueL2.String()},
	}, nil
}
