package claimtxman

import (
	"context"

	"github.com/0xPolygonHermez/zkevm-bridge-service/bridgectrl"
	"github.com/0xPolygonHermez/zkevm-bridge-service/claimtxman/types"
	"github.com/0xPolygonHermez/zkevm-bridge-service/etherman"
	"github.com/ethereum/go-ethereum/common"
)

type StorageInterface interface {
	AddBlock(ctx context.Context, block *etherman.Block, dbTx interface{}) (uint64, error)
	UpdateL1DepositsStatus(ctx context.Context, exitRoot []byte, destinationNetwork uint32, dbTx interface{}) error
	UpdateL2DepositsStatus(ctx context.Context, exitRoot []byte, rollupID, networkID uint32, dbTx interface{}) error
	GetDepositsFromOtherL2ToClaim(ctx context.Context, destinationNetwork uint32, dbTx interface{}) ([]*etherman.Deposit, error)
	GetPendingDepositsToClaim(ctx context.Context, destAddress common.Address, destNetwork, leafType, limit, offset uint32, fromNetwork int8, dbTx interface{}) ([]*etherman.Deposit, uint64, error)
	GetLatestTrustedGERByDeposit(ctx context.Context, depositCnt, networkID, destinationNetwork uint32, dbTx interface{}) (common.Hash, error)
	AddClaimTx(ctx context.Context, mTx types.MonitoredTx, dbTx interface{}) error
	UpdateClaimTx(ctx context.Context, mTx types.MonitoredTx, dbTx interface{}) error
	GetClaimTxsByStatus(ctx context.Context, statuses []types.MonitoredTxStatus, rollupID uint32, dbTx interface{}) ([]types.MonitoredTx, error)
	GetDepositByDepositID(ctx context.Context, depositID uint64, dbTx interface{}) (*etherman.Deposit, error)
	// atomic
	Rollback(ctx context.Context, dbTx interface{}) error
	BeginDBTransaction(ctx context.Context) (interface{}, error)
	Commit(ctx context.Context, dbTx interface{}) error
}

type bridgeServiceInterface interface {
	GetClaimProofForCompressed(ctx context.Context, ger common.Hash, depositCnt, networkID uint32, dbTx interface{}) (*etherman.GlobalExitRoot, [][bridgectrl.KeyLen]byte, [][bridgectrl.KeyLen]byte, error)
	GetDepositStatus(ctx context.Context, depositCount, networkID, destNetworkID uint32) (string, error)
}
