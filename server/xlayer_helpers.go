package server

import (
	"context"
	"math/big"
	"time"

	"github.com/0xPolygonHermez/zkevm-bridge-service/bridgectrl/pb"
	"github.com/0xPolygonHermez/zkevm-bridge-service/etherman"
	"github.com/0xPolygonHermez/zkevm-bridge-service/log"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/pushtask"
)

func (s *bridgeService) getGlobalIndex(deposit *etherman.Deposit) *big.Int {
	mainnetFlag := deposit.NetworkID == 0
	rollupIndex := deposit.NetworkID - 1
	localExitRootIndex := deposit.DepositCount
	return etherman.GenerateGlobalIndex(mainnetFlag, rollupIndex, localExitRootIndex)
}

func (s *bridgeService) setDurationForL2Deposit(ctx context.Context, l2AvgCommitDuration uint64, l2AvgVerifyDuration uint64, currTime time.Time,
	tx *pb.Transaction, depositCreateTime time.Time) {
	var duration int
	if tx.Status == uint32(pb.TransactionStatus_TX_CREATED) {
		duration = pushtask.GetLeftCommitTime(depositCreateTime, l2AvgCommitDuration, currTime)
	} else {
		duration = pushtask.GetLeftVerifyTime(ctx, rsstore, tx.BlockNumber, depositCreateTime, l2AvgCommitDuration, l2AvgVerifyDuration, currTime)
	}
	if duration <= 0 {
		log.Debugf("count EstimateTime for L2 -> L1 over range, so use min default duration: %v", defaultMinDuration)
		tx.EstimateTime = uint32(defaultMinDuration)
		return
	}
	tx.EstimateTime = uint32(duration)
}

// For testing purposes
func (s *bridgeService) GetFakePushMessages(ctx context.Context, req *pb.GetFakePushMessagesRequest) (*pb.GetFakePushMessagesResponse, error) {
	if messagePushProducer == nil {
		return &pb.GetFakePushMessagesResponse{
			Code: uint32(pb.ErrorCode_ERROR_DEFAULT),
			Msg:  "producer is nil",
		}, nil
	}

	return &pb.GetFakePushMessagesResponse{
		Code: uint32(pb.ErrorCode_ERROR_OK),
		Data: messagePushProducer.GetFakeMessages(req.Topic),
	}, nil
}
