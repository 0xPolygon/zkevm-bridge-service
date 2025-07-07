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
	UpdateL1DepositsStatus(ctx context.Context, exitRoot []byte, destinationNetwork uint32, dbTx interface{}) ([]*etherman.Deposit, error)
	UpdateL2DepositsStatus(ctx context.Context, exitRoot []byte, rollupID, networkID uint32, dbTx interface{}) error
	GetDepositsFromOtherL2ToClaim(ctx context.Context, destinationNetwork uint32, dbTx interface{}) ([]*etherman.Deposit, error)
	GetLatestTrustedGERByDeposit(ctx context.Context, depositCnt, networkID, destinationNetwork uint32, dbTx interface{}) (common.Hash, error)
	AddClaimTx(ctx context.Context, mTx types.MonitoredTx, dbTx interface{}) error
	UpdateClaimTx(ctx context.Context, mTx types.MonitoredTx, dbTx interface{}) error
	GetClaimTxsByStatus(ctx context.Context, statuses []types.MonitoredTxStatus, rollupID uint32, dbTx interface{}) ([]types.MonitoredTx, error)
	// atomic
	Rollback(ctx context.Context, dbTx interface{}) error
	BeginDBTransaction(ctx context.Context) (interface{}, error)
	Commit(ctx context.Context, dbTx interface{}) error

	// XLayer
	UpdateL1DepositsStatusXLayer(ctx context.Context, exitRoot []byte, dbTx interface{}) ([]*etherman.Deposit, error)
	UpdateL2DepositsStatusXLayer(ctx context.Context, exitRoot []byte, rollupID, networkID uint, dbTx interface{}) ([]*etherman.Deposit, error)
	GetL1Deposits(ctx context.Context, exitRoot []byte, dbTx interface{}) ([]*etherman.Deposit, error)
	UpdateL1DepositStatus(ctx context.Context, depositCount uint, dbTx interface{}) error
	GetDeposit(ctx context.Context, depositCnt, networkID uint32, dbTx interface{}) (*etherman.Deposit, error)
	GetClaim(ctx context.Context, index, depositCount, networkID uint32, dbTx interface{}) (*etherman.Claim, error)
	GetClaimTxsByStatusWithLimit(ctx context.Context, statuses []types.MonitoredTxStatus, limit, offset uint, dbTx interface{}) ([]types.MonitoredTx, error)
}

type bridgeServiceInterface interface {
	GetDepositStatus(ctx context.Context, depositCount, networkID, destNetworkID uint32) (string, error)
	GetClaimProofForCompressed(ger common.Hash, depositCnt, networkID uint32, dbTx interface{}) (*etherman.GlobalExitRoot, [][bridgectrl.KeyLen]byte, [][bridgectrl.KeyLen]byte, error)

	// From XLayer branch
	GetClaimProof(depositCnt, networkID uint32, dbTx interface{}) (*etherman.GlobalExitRoot, [][bridgectrl.KeyLen]byte, [][bridgectrl.KeyLen]byte, error)
}
