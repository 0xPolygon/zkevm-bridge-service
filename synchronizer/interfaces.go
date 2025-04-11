package synchronizer

import (
	"context"
	"math/big"

	"github.com/0xPolygonHermez/zkevm-bridge-service/etherman"
	rpcTypes "github.com/0xPolygonHermez/zkevm-bridge-service/jsonrpcclient/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ethermanInterface contains the methods required to interact with ethereum.
type ethermanInterface interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	GetRollupInfoByBlockRange(ctx context.Context, fromBlock uint64, toBlock *uint64) ([]etherman.Block, map[common.Hash][]etherman.Order, error)
	GetNetworkID() uint32
}

type storageInterface interface {
	GetLastBlock(ctx context.Context, networkID uint32, dbTx interface{}) (*etherman.Block, error)
	Rollback(ctx context.Context, dbTx interface{}) error
	BeginDBTransaction(ctx context.Context) (interface{}, error)
	Commit(ctx context.Context, dbTx interface{}) error
	AddBlock(ctx context.Context, block *etherman.Block, dbTx interface{}) (uint64, error)
	AddGlobalExitRoot(ctx context.Context, exitRoot *etherman.GlobalExitRoot, dbTx interface{}) error
	AddDeposit(ctx context.Context, deposit *etherman.Deposit, dbTx interface{}) (uint64, error)
	AddClaim(ctx context.Context, claim *etherman.Claim, dbTx interface{}) error
	AddTokenWrapped(ctx context.Context, tokenWrapped *etherman.TokenWrapped, dbTx interface{}) error
	Reset(ctx context.Context, blockNumber uint64, networkID uint32, dbTx interface{}) error
	GetPreviousBlock(ctx context.Context, networkID uint32, offset uint64, dbTx interface{}) (etherman.Block, error)
	GetNumberDeposits(ctx context.Context, origNetworkID uint32, blockNumber uint64, dbTx interface{}) (uint32, error)
	AddTrustedGlobalExitRoot(ctx context.Context, trustedExitRoot *etherman.GlobalExitRoot, dbTx interface{}) (bool, error)
	GetLatestL1SyncedExitRoot(ctx context.Context, dbTx interface{}) (*etherman.GlobalExitRoot, error)
	GetLatestTrustedExitRoot(ctx context.Context, networkID uint32, dbTx interface{}) (*etherman.GlobalExitRoot, error)
	CheckIfRootExists(ctx context.Context, root []byte, network uint32, dbTx interface{}) (bool, error)
	GetL1ExitRootByGER(ctx context.Context, ger common.Hash, dbTx interface{}) (*etherman.GlobalExitRoot, error)
	GetL2ExitRootsByGER(ctx context.Context, ger common.Hash, dbTx interface{}) ([]etherman.GlobalExitRoot, error)
	UpdateL2GER(ctx context.Context, ger etherman.GlobalExitRoot, dbTx interface{}) error
	AddRemoveL2GER(ctx context.Context, globalExitRoot etherman.GlobalExitRoot, dbTx interface{}) error
}

type bridgectrlInterface interface {
	AddDeposit(ctx context.Context, deposit *etherman.Deposit, dbTx interface{}) error
	ReorgMT(ctx context.Context, depositCount, networkID uint32, dbTx interface{}) error
	RollbackMT(ctx context.Context, networkID uint32, dbTx interface{}) error
	AddRollupExitLeaf(ctx context.Context, rollupLeaf etherman.RollupExitLeaf, dbTx interface{}) error
}

type zkEVMClientInterface interface {
	GetLatestGlobalExitRoot(ctx context.Context) (common.Hash, error)
	ExitRootsByGER(ctx context.Context, globalExitRoot common.Hash) (*rpcTypes.ExitRoots, error)
}
