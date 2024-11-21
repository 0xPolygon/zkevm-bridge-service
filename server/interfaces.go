package server

import (
	"context"

	"github.com/0xPolygonHermez/zkevm-bridge-service/etherman"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
)

type bridgeServiceStorage interface {
	Get(ctx context.Context, key []byte, dbTx pgx.Tx) ([][]byte, error)
	GetRoot(ctx context.Context, depositCnt, network uint32, dbTx pgx.Tx) ([]byte, error)
	GetDepositCountByRoot(ctx context.Context, root []byte, network uint32, dbTx pgx.Tx) (uint32, error)
	GetLatestExitRoot(ctx context.Context, networkID, destNetwork uint32, dbTx pgx.Tx) (*etherman.GlobalExitRoot, error)
	GetExitRootByGER(ctx context.Context, ger common.Hash, dbTx pgx.Tx) (*etherman.GlobalExitRoot, error)
	GetClaim(ctx context.Context, index, originNetworkID, networkID uint32, dbTx pgx.Tx) (*etherman.Claim, error)
	GetClaims(ctx context.Context, destAddr string, limit, offset uint32, dbTx pgx.Tx) ([]*etherman.Claim, error)
	GetClaimCount(ctx context.Context, destAddr string, dbTx pgx.Tx) (uint64, error)
	GetDeposit(ctx context.Context, depositCnt, networkID uint32, dbTx pgx.Tx) (*etherman.Deposit, error)
	GetDeposits(ctx context.Context, destAddr string, limit, offset uint32, dbTx pgx.Tx) ([]*etherman.Deposit, error)
	GetDepositCount(ctx context.Context, destAddr string, dbTx pgx.Tx) (uint64, error)
	GetTokenWrapped(ctx context.Context, originalNetwork uint32, originalTokenAddress common.Address, dbTx pgx.Tx) (*etherman.TokenWrapped, error)
	GetRollupExitLeavesByRoot(ctx context.Context, root common.Hash, dbTx pgx.Tx) ([]etherman.RollupExitLeaf, error)
	GetPendingDepositsToClaim(ctx context.Context, destAddress common.Address, destNetwork, leafType, limit, offset uint32, dbTx pgx.Tx) ([]*etherman.Deposit, uint64, error)
}
