package operations

import (
	"context"

	"github.com/0xPolygon/zkevm-bridge-service/bridgectrl/pb"
	"github.com/0xPolygon/zkevm-bridge-service/etherman"
	"github.com/ethereum/go-ethereum/common"
)

// StorageInterface is a storage interface.
type StorageInterface interface {
	GetLastBlock(ctx context.Context, networkID uint32, dbTx interface{}) (*etherman.Block, error)
	GetLatestExitRoot(ctx context.Context, networkID, destNetwork uint32, dbTx interface{}) (*etherman.GlobalExitRoot, error)
	GetLatestL1SyncedExitRoot(ctx context.Context, dbTx interface{}) (*etherman.GlobalExitRoot, error)
	GetLatestTrustedExitRoot(ctx context.Context, networkID uint32, dbTx interface{}) (*etherman.GlobalExitRoot, error)
	GetTokenWrapped(ctx context.Context, originalNetwork uint32, originalTokenAddress common.Address, dbTx interface{}) (*etherman.TokenWrapped, error)
	GetDepositCountByRoot(ctx context.Context, root []byte, network uint32, dbTx interface{}) (uint32, error)
	UpdateBlocksForTesting(ctx context.Context, networkID uint32, blockNum uint64, dbTx interface{}) error
	GetClaim(ctx context.Context, depositCount, origNetworkID, networkID uint32, dbTx interface{}) (*etherman.Claim, error)
	GetClaims(ctx context.Context, destAddr string, limit uint32, offset uint32, dbTx interface{}) ([]*etherman.Claim, error)
	UpdateDepositsStatusForTesting(ctx context.Context, dbTx interface{}) error
	GetLatestMonitoredTxGroupID(ctx context.Context, dbTx interface{}) (uint64, error)
	// synchronizer
	AddBlock(ctx context.Context, block *etherman.Block, dbTx interface{}) (uint64, error)
	AddGlobalExitRoot(ctx context.Context, exitRoot *etherman.GlobalExitRoot, dbTx interface{}) error
	AddTrustedGlobalExitRoot(ctx context.Context, trustedExitRoot *etherman.GlobalExitRoot, dbTx interface{}) (bool, error)
	AddDeposit(ctx context.Context, deposit *etherman.Deposit, dbTx interface{}) (uint64, error)
	AddClaim(ctx context.Context, claim *etherman.Claim, dbTx interface{}) error
	AddTokenWrapped(ctx context.Context, tokenWrapped *etherman.TokenWrapped, dbTx interface{}) error
	// atomic
	Rollback(ctx context.Context, dbTx interface{}) error
	BeginDBTransaction(ctx context.Context) (interface{}, error)
	Commit(ctx context.Context, dbTx interface{}) error
}

// BridgeServiceInterface is an interface for the bridge service.
type BridgeServiceInterface interface {
	GetBridges(ctx context.Context, req *pb.GetBridgesRequest) (*pb.GetBridgesResponse, error)
	GetClaims(ctx context.Context, req *pb.GetClaimsRequest) (*pb.GetClaimsResponse, error)
	GetProof(ctx context.Context, req *pb.GetProofRequest) (*pb.GetProofResponse, error)
	GetProofByGER(ctx context.Context, req *pb.GetProofByGERRequest) (*pb.GetProofResponse, error)
}
