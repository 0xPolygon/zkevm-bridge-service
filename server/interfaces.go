package server

import (
	"context"
	"math/big"
	"time"

	"github.com/0xPolygonHermez/zkevm-bridge-service/etherman"
	"github.com/ethereum/go-ethereum/common"
	pgx "github.com/jackc/pgx/v4"

	ctmtypes "github.com/0xPolygonHermez/zkevm-bridge-service/claimtxman/types"
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
	GetPendingDepositsToClaim(ctx context.Context, destAddress common.Address, destNetwork, leafType, limit, offset uint32, dbTx interface{}) ([]*etherman.Deposit, uint64, error)

	// XLayer
	GetDepositByHash(ctx context.Context, destAddr string, networkID uint, txHash string, dbTx pgx.Tx) (*etherman.Deposit, error)
	GetDepositsXLayer(ctx context.Context, destAddr string, limit uint, offset uint, messageAllowlist []common.Address, dbTx pgx.Tx) ([]*etherman.Deposit, error)
	GetPendingTransactions(ctx context.Context, destAddr string, limit uint, offset uint, messageAllowlist []common.Address, dbTx pgx.Tx) ([]*etherman.Deposit, error)
	GetNotReadyTransactions(ctx context.Context, limit uint, offset uint, dbTx pgx.Tx) ([]*etherman.Deposit, error)
	GetReadyPendingTransactions(ctx context.Context, networkID uint, limit uint, offset uint, minReadyTime time.Time, dbTx pgx.Tx) ([]*etherman.Deposit, error)
	GetClaimTxById(ctx context.Context, id uint, dbTx pgx.Tx) (*ctmtypes.MonitoredTx, error)
	GetClaimTxsByStatusWithLimit(ctx context.Context, statuses []ctmtypes.MonitoredTxStatus, limit uint, offset uint, dbTx pgx.Tx) ([]ctmtypes.MonitoredTx, error)
	GetDepositsForUnitTest(ctx context.Context, destAddr string, limit uint, offset uint, dbTx pgx.Tx) ([]*etherman.Deposit, error)
	GetBridgeBalance(ctx context.Context, originalTokenAddr common.Address, networkID uint, forUpdate bool, dbTx pgx.Tx) (*big.Int, error)
	SetBridgeBalance(ctx context.Context, originalTokenAddr common.Address, networkID uint, balance *big.Int, dbTx pgx.Tx) error
}
