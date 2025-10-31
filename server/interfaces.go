package server

import (
	"context"

	"github.com/0xPolygon/zkevm-bridge-service/etherman"
	"github.com/ethereum/go-ethereum/common"
)

type bridgeServiceStorage interface {
	Get(ctx context.Context, key []byte, dbTx interface{}) ([][]byte, error)
	GetRoot(ctx context.Context, depositCnt, network uint32, dbTx interface{}) ([]byte, error)
	GetDepositCountByRoot(ctx context.Context, root []byte, network uint32, dbTx interface{}) (uint32, error)
	GetLatestExitRoot(ctx context.Context, networkID, destNetwork uint32, dbTx interface{}) (*etherman.GlobalExitRoot, error)
	GetL1ExitRootByGER(ctx context.Context, ger common.Hash, dbTx interface{}) (*etherman.GlobalExitRoot, error)
	GetLatestTrustedExitRoot(ctx context.Context, networkID uint32, dbTx interface{}) (*etherman.GlobalExitRoot, error)
	GetClaim(ctx context.Context, index, originNetworkID, networkID uint32, dbTx interface{}) (*etherman.Claim, error)
	GetClaims(ctx context.Context, destAddr string, limit, offset uint32, dbTx interface{}) ([]*etherman.Claim, error)
	GetClaimCount(ctx context.Context, destAddr string, dbTx interface{}) (uint64, error)
	GetDeposit(ctx context.Context, depositCnt, networkID uint32, dbTx interface{}) (*etherman.Deposit, error)
	GetDeposits(ctx context.Context, destAddr string, limit, offset uint32, dbTx interface{}) ([]*etherman.Deposit, error)
	GetDepositCount(ctx context.Context, destAddr string, dbTx interface{}) (uint64, error)
	GetTokenWrapped(ctx context.Context, originalNetwork uint32, originalTokenAddress common.Address, dbTx interface{}) (*etherman.TokenWrapped, error)
	GetRollupExitLeavesByRoot(ctx context.Context, root common.Hash, dbTx interface{}) ([]etherman.RollupExitLeaf, error)
	GetPendingDepositsToClaim(ctx context.Context, destAddress common.Address, destNetwork, leafType, limit, offset uint32, fromNetwork int8, dbTx interface{}) ([]*etherman.Deposit, uint64, error)
	GetSyncStatus(ctx context.Context, dbTx interface{}) ([]*etherman.SyncStatus, error)
}
