package db

import (
	"context"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"os"
	"testing"

	"github.com/0xPolygonHermez/zkevm-bridge-service/db/pgstorage"
	"github.com/0xPolygonHermez/zkevm-bridge-service/etherman"
	"github.com/ethereum/go-ethereum/common"
	pgx "github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	err := os.Setenv("ZKEVM_BRIDGE_SYNCDB_DATABASE", "postgres")
	if err != nil {
		panic(err)
	}
	_, exists := os.LookupEnv("ZKEVM_BRIDGE_SYNCDB_DATABASE")
	if !exists {
		panic("ZKEVM_BRIDGE_SYNCDB_DATABASE env var not set")
	}
}

func TestInsertDeposit(t *testing.T) {
	ctx := context.Background()
	testStore, err := newStorageSettings(os.Getenv("ZKEVM_BRIDGE_SYNCDB_DATABASE"))
	require.NoError(t, err)
	tx, err := testStore.BeginDBTransaction(ctx)
	require.NoError(t, err)

	blockID, err := testStore.AddBlock(ctx, &etherman.Block{
		BlockNumber: 1,
		BlockHash:   common.HexToHash("0x29e885adaf8e4b51e4d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9d1"),
	}, nil)
	require.NoError(t, err)
	deposit := &etherman.Deposit{
		NetworkID:          1,
		OriginalNetwork:    4294967295,
		OriginalAddress:    common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
		Amount:             big.NewInt(1000000),
		DestinationNetwork: 4294967295,
		DestinationAddress: common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		BlockNumber:        1,
		BlockID:            blockID,
		DepositCount:       1,
		Metadata:           common.FromHex("0x00"),
	}
	_, err = testStore.AddDeposit(ctx, deposit, tx)
	require.NoError(t, err)
	require.NoError(t, testStore.Rollback(ctx, tx))
}

func TestL1GlobalExitRoot(t *testing.T) {
	ctx := context.Background()
	testStore, err := newStorageSettings(os.Getenv("ZKEVM_BRIDGE_SYNCDB_DATABASE"))
	require.NoError(t, err)
	tx, err := testStore.BeginDBTransaction(ctx)
	require.NoError(t, err)

	block := &etherman.Block{
		BlockNumber: 1,
		BlockHash:   common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f1"),
		NetworkID:   0,
	}

	blockID, err := testStore.AddBlock(ctx, block, tx)
	require.NoError(t, err)
	require.Equal(t, blockID, uint64(1))

	l1GER := &etherman.GlobalExitRoot{
		BlockID:        1,
		ExitRoots:      []common.Hash{common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f1"), common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f1")},
		GlobalExitRoot: common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f1"),
	}

	err = testStore.AddGlobalExitRoot(ctx, l1GER, tx)
	require.NoError(t, err)

	ger, err := testStore.GetLatestL1SyncedExitRoot(ctx, tx)
	require.NoError(t, err)
	require.Equal(t, ger.BlockID, l1GER.BlockID)
	require.Equal(t, ger.GlobalExitRoot, l1GER.GlobalExitRoot)

	latestGER, err := testStore.GetLatestExitRoot(ctx, 1, 0, tx)
	require.NoError(t, err)
	require.Equal(t, latestGER.GlobalExitRoot, l1GER.GlobalExitRoot)
	require.Equal(t, latestGER.BlockNumber, l1GER.BlockNumber)
	require.Equal(t, latestGER.ExitRoots[0], l1GER.ExitRoots[0])
	require.Equal(t, latestGER.ExitRoots[1], l1GER.ExitRoots[1])

	require.NoError(t, testStore.Commit(ctx, tx))
}

func TestAddTrustedGERDuplicated(t *testing.T) {
	ctx := context.Background()
	storageType := os.Getenv("ZKEVM_BRIDGE_SYNCDB_DATABASE")
	testStore, err := newStorageSettings(storageType)
	require.NoError(t, err)
	tx, err := testStore.BeginDBTransaction(ctx)
	require.NoError(t, err)

	ger := &etherman.GlobalExitRoot{
		NetworkID:      1,
		ExitRoots:      []common.Hash{common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f1"), common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f1")},
		GlobalExitRoot: common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f1"),
	}
	isInserted, err := testStore.AddTrustedGlobalExitRoot(ctx, ger, tx)
	require.True(t, isInserted)
	require.NoError(t, err)
	getCount := "select count(*) from %sexit_root where block_id = 0 AND global_exit_root = "
	var result int
	if storageType == "postgres" {
		query := fmt.Sprintf(getCount+"'\\x%s'", "sync.", hex.EncodeToString(ger.GlobalExitRoot.Bytes()))
		err = testStore.QueryRowTesting(ctx, query, tx).(pgx.Row).Scan(&result)
	} else {
		require.NoError(t, fmt.Errorf("database type not supported"))
	}
	require.NoError(t, err)
	assert.Equal(t, 1, result)
	isInserted, err = testStore.AddTrustedGlobalExitRoot(ctx, ger, tx)
	require.False(t, isInserted)
	require.NoError(t, err)
	if storageType == "postgres" {
		err = testStore.QueryRowTesting(ctx, fmt.Sprintf(getCount+"'\\x%s'", "sync.", hex.EncodeToString(ger.GlobalExitRoot.Bytes())), tx).(pgx.Row).Scan(&result)
	} else {
		require.NoError(t, fmt.Errorf("database type not supported"))
	}
	require.NoError(t, err)
	assert.Equal(t, 1, result)
	require.NoError(t, testStore.Commit(ctx, tx))

	tx, err = testStore.BeginDBTransaction(ctx)
	require.NoError(t, err)

	ger1 := &etherman.GlobalExitRoot{
		NetworkID:      1,
		ExitRoots:      []common.Hash{common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f2"), common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f2")},
		GlobalExitRoot: common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f2"),
	}
	isInserted, err = testStore.AddTrustedGlobalExitRoot(ctx, ger, tx)
	require.False(t, isInserted)
	require.NoError(t, err)
	if storageType == "postgres" {
		err = testStore.QueryRowTesting(ctx, fmt.Sprintf(getCount+"'\\x%s'", "sync.", hex.EncodeToString(ger.GlobalExitRoot.Bytes())), tx).(pgx.Row).Scan(&result)
	} else {
		require.NoError(t, fmt.Errorf("database type not supported"))
	}
	require.NoError(t, err)
	assert.Equal(t, 1, result)
	isInserted, err = testStore.AddTrustedGlobalExitRoot(ctx, ger1, tx)
	require.True(t, isInserted)
	require.NoError(t, err)
	getCount2 := "select count(*) from %sexit_root"
	if storageType == "postgres" {
		err = testStore.QueryRowTesting(ctx, fmt.Sprintf(getCount2, "sync."), tx).(pgx.Row).Scan(&result)
	} else {
		require.NoError(t, fmt.Errorf("database type not supported"))
	}
	require.NoError(t, err)
	assert.Equal(t, 2, result)

	blockID, err := testStore.AddBlock(ctx, &etherman.Block{
		BlockNumber: 1,
		BlockHash:   common.HexToHash("0x29e995edaf8e4b51e1d2e05f9da28161d2fb4efb1d53827d9b80a23cf2d7a3f2"),
	}, tx)
	require.NoError(t, err)
	ger2 := ger1
	ger2.BlockID = blockID
	err = testStore.AddGlobalExitRoot(ctx, ger2, tx)
	require.NoError(t, err)

	tGER, err := testStore.GetLatestTrustedExitRoot(ctx, 1, tx)
	require.NoError(t, err)
	require.Equal(t, tGER.GlobalExitRoot, ger1.GlobalExitRoot)

	latestGER, err := testStore.GetLatestExitRoot(ctx, 0, 1, tx)
	require.NoError(t, err)
	require.Equal(t, latestGER.GlobalExitRoot, ger1.GlobalExitRoot)
	require.Equal(t, latestGER.BlockNumber, ger1.BlockNumber)
	require.Equal(t, latestGER.ExitRoots[0], ger1.ExitRoots[0])
	require.Equal(t, latestGER.ExitRoots[1], ger1.ExitRoots[1])

	require.NoError(t, testStore.Commit(ctx, tx))
}

func TestGetLastBlock(t *testing.T) {
	ctx := context.Background()
	testStore, err := newStorageSettings(os.Getenv("ZKEVM_BRIDGE_SYNCDB_DATABASE"))
	require.NoError(t, err)
	tx, err := testStore.BeginDBTransaction(ctx)
	require.NoError(t, err)
	block1 := etherman.Block{
		BlockNumber: 1,
		BlockHash:   common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f1"),
		NetworkID:   0,
	}
	block2 := etherman.Block{
		BlockNumber: 2,
		BlockHash:   common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f3"),
		NetworkID:   0,
	}
	block3 := etherman.Block{
		BlockNumber: 100,
		BlockHash:   common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f5"),
		NetworkID:   1,
	}
	block4 := etherman.Block{
		BlockNumber: 101,
		BlockHash:   common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f7"),
		NetworkID:   1,
	}
	_, err = testStore.AddBlock(ctx, &block1, tx)
	require.NoError(t, err)
	b, err := testStore.GetLastBlock(ctx, 0, tx)
	require.NoError(t, err)
	assert.Equal(t, block1.BlockNumber, b.BlockNumber)
	assert.Equal(t, block1.BlockHash, b.BlockHash)
	assert.Equal(t, block1.NetworkID, b.NetworkID)

	_, err = testStore.AddBlock(ctx, &block2, tx)
	require.NoError(t, err)
	b, err = testStore.GetLastBlock(ctx, 0, tx)
	require.NoError(t, err)
	assert.Equal(t, block2.BlockNumber, b.BlockNumber)
	assert.Equal(t, block2.BlockHash, b.BlockHash)
	assert.Equal(t, block2.NetworkID, b.NetworkID)

	_, err = testStore.AddBlock(ctx, &block3, tx)
	require.NoError(t, err)
	b, err = testStore.GetLastBlock(ctx, 1, tx)
	require.NoError(t, err)
	assert.Equal(t, block3.BlockNumber, b.BlockNumber)
	assert.Equal(t, block3.BlockHash, b.BlockHash)
	assert.Equal(t, block3.NetworkID, b.NetworkID)

	_, err = testStore.AddBlock(ctx, &block4, tx)
	require.NoError(t, err)
	b, err = testStore.GetLastBlock(ctx, 1, tx)
	require.NoError(t, err)
	assert.Equal(t, block4.BlockNumber, b.BlockNumber)
	assert.Equal(t, block4.BlockHash, b.BlockHash)
	assert.Equal(t, block4.NetworkID, b.NetworkID)

	prevBlock, err := testStore.GetPreviousBlock(ctx, 1, 1, tx)
	require.NoError(t, err)
	require.Equal(t, prevBlock.BlockNumber, block3.BlockNumber)
	require.Equal(t, prevBlock.BlockHash, block3.BlockHash)

	require.NoError(t, testStore.Commit(ctx, tx))
}

// Test MerkleTree storage
func TestMTStorage(t *testing.T) {
	ctx := context.Background()
	testStore, err := newStorageSettings(os.Getenv("ZKEVM_BRIDGE_SYNCDB_DATABASE"))
	require.NoError(t, err)
	tx, err := testStore.BeginDBTransaction(ctx)
	require.NoError(t, err)

	leaf1 := common.FromHex("0xa4bfa0908dc7b06d98da4309f859023d6947561bc19bc00d77f763dea1a0b9f5")
	leaf2 := common.FromHex("0x315fee1aa202bf4a6bd0fde560c89be90b6e6e2aaf92dc5e8d118209abc3410f")
	root := common.FromHex("0x88e652896cb1de5962a0173a222059f51e6b943a2ba6dfc9acbff051ceb1abb5")
	deposit := &etherman.Deposit{
		Metadata: common.Hex2Bytes("0x0"),
	}
	depositID, err := testStore.AddDeposit(ctx, deposit, tx)
	require.NoError(t, err)
	err = testStore.SetRoot(ctx, root, depositID, 0, tx)
	require.NoError(t, err)

	err = testStore.Set(ctx, root, [][]byte{leaf1, leaf2}, depositID, tx)
	require.NoError(t, err)

	vals, err := testStore.Get(ctx, root, tx)
	require.NoError(t, err)
	require.Equal(t, leaf1, vals[0])
	require.Equal(t, leaf2, vals[1])

	rRoot, err := testStore.GetRoot(ctx, 0, 0, tx)
	require.NoError(t, err)
	require.Equal(t, root, rRoot)

	count, err := testStore.GetLastDepositCount(ctx, 0, tx)
	require.NoError(t, err)
	require.Equal(t, uint32(0), count)

	dCount, err := testStore.GetDepositCountByRoot(ctx, root, 0, tx)
	require.NoError(t, err)
	require.Equal(t, uint32(0), dCount)

	require.NoError(t, testStore.Commit(ctx, tx))
}

// Test BridgeService storage
func TestBSStorage(t *testing.T) {
	ctx := context.Background()
	testStore, err := newStorageSettings(os.Getenv("ZKEVM_BRIDGE_SYNCDB_DATABASE"))
	require.NoError(t, err)
	tx, err := testStore.BeginDBTransaction(ctx)
	require.NoError(t, err)

	block := &etherman.Block{
		BlockNumber: 1,
		BlockHash:   common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f1"),
		NetworkID:   0,
	}
	_, err = testStore.AddBlock(ctx, block, tx)
	require.NoError(t, err)

	deposit := &etherman.Deposit{
		NetworkID:          0,
		OriginalNetwork:    0,
		OriginalAddress:    common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
		Amount:             big.NewInt(1000000),
		DestinationNetwork: 1,
		DestinationAddress: common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		BlockNumber:        1,
		BlockID:            1,
		DepositCount:       1,
		Metadata:           common.FromHex("0x000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000005436f696e410000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003434f410000000000000000000000000000000000000000000000000000000000"),
	}
	_, err = testStore.AddDeposit(ctx, deposit, tx)
	require.NoError(t, err)

	claim := &etherman.Claim{
		Index:              1,
		OriginalNetwork:    0,
		OriginalAddress:    common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
		Amount:             big.NewInt(1000000),
		DestinationAddress: common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		BlockID:            1,
		BlockNumber:        2,
		NetworkID:          1,
		TxHash:             common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f2"),
		RollupIndex:        1,
		MainnetFlag:        true,
	}
	err = testStore.AddClaim(ctx, claim, tx)
	require.NoError(t, err)

	count, err := testStore.GetDepositCount(ctx, deposit.DestinationAddress.String(), tx)
	require.NoError(t, err)
	require.Equal(t, count, uint64(1))

	rDeposit, err := testStore.GetDeposit(ctx, 1, 0, tx)
	require.NoError(t, err)
	require.Equal(t, rDeposit.DestinationAddress, deposit.DestinationAddress)
	require.Equal(t, rDeposit.DepositCount, deposit.DepositCount)

	rDeposits, err := testStore.GetDeposits(ctx, deposit.DestinationAddress.String(), 10, 0, tx)
	require.NoError(t, err)
	require.Equal(t, len(rDeposits), 1)

	countND, err := testStore.GetNumberDeposits(ctx, 0, 0, tx)
	require.NoError(t, err)
	require.Equal(t, countND, uint32(0))
	countND, err = testStore.GetNumberDeposits(ctx, 0, 1, tx)
	require.NoError(t, err)
	require.Equal(t, countND, uint32(2))

	count, err = testStore.GetClaimCount(ctx, claim.DestinationAddress.String(), tx)
	require.NoError(t, err)
	require.Equal(t, count, uint64(1))

	rClaim, err := testStore.GetClaim(ctx, deposit.DepositCount, deposit.NetworkID, claim.NetworkID, tx)
	require.NoError(t, err)
	require.Equal(t, rClaim.DestinationAddress, claim.DestinationAddress)
	require.Equal(t, rClaim.NetworkID, claim.NetworkID)
	require.Equal(t, rClaim.Index, claim.Index)
	require.Equal(t, rClaim.RollupIndex, claim.RollupIndex)
	require.Equal(t, rClaim.MainnetFlag, claim.MainnetFlag)

	rClaims, err := testStore.GetClaims(ctx, claim.DestinationAddress.String(), 10, 0, tx)
	require.NoError(t, err)
	require.Equal(t, len(rClaims), 1)

	wrappedToken := &etherman.TokenWrapped{
		OriginalNetwork:      0,
		OriginalTokenAddress: deposit.OriginalAddress,
		WrappedTokenAddress:  common.HexToAddress("0x187Bd40226A7073b49163b1f6c2b73d8F2aa8478"),
		BlockID:              1,
		BlockNumber:          1,
		NetworkID:            1,
	}
	metadata, err := testStore.GetTokenMetadata(ctx, wrappedToken.OriginalNetwork, wrappedToken.NetworkID, wrappedToken.OriginalTokenAddress, tx)
	require.NoError(t, err)
	require.Equal(t, metadata, deposit.Metadata)

	err = testStore.AddTokenWrapped(ctx, wrappedToken, tx)
	require.NoError(t, err)

	wt, err := testStore.GetTokenWrapped(ctx, wrappedToken.OriginalNetwork, wrappedToken.OriginalTokenAddress, tx)
	require.NoError(t, err)
	require.Equal(t, wt.WrappedTokenAddress, wrappedToken.WrappedTokenAddress)
	require.Equal(t, wt.Name, "CoinA")
	require.Equal(t, wt.Symbol, "COA")
	require.Equal(t, wt.Decimals, uint8(12))

	require.NoError(t, testStore.Commit(ctx, tx))
}

// Test Set Max uint as networkID into setRoot storage
func TestSetMaxUintNetworkID(t *testing.T) {
	ctx := context.Background()
	testStore, err := newStorageSettings(os.Getenv("ZKEVM_BRIDGE_SYNCDB_DATABASE"))
	require.NoError(t, err)
	tx, err := testStore.BeginDBTransaction(ctx)
	require.NoError(t, err)
	deposit := &etherman.Deposit{
		Metadata: common.Hex2Bytes("0x0"),
	}
	depositID, err := testStore.AddDeposit(ctx, deposit, tx)
	require.NoError(t, err)
	root := common.FromHex("0x88e652896cb1de5962a0173a222059f51e6b943a2ba6dfc9acbff051ceb1abb5")
	err = testStore.SetRoot(ctx, root, depositID, math.MaxInt32, tx)
	require.NoError(t, err)
	rRoot, err := testStore.GetRoot(ctx, 0, math.MaxInt32, tx)
	require.NoError(t, err)
	require.Equal(t, root, rRoot)
	require.NoError(t, testStore.Commit(ctx, tx))
}

func TestIncompleteL2GlobalExitRoot(t *testing.T) {
	ctx := context.Background()
	testStore, err := newStorageSettings(os.Getenv("ZKEVM_BRIDGE_SYNCDB_DATABASE"))
	require.NoError(t, err)
	tx, err := testStore.BeginDBTransaction(ctx)
	require.NoError(t, err)

	block := &etherman.Block{
		BlockNumber: 1,
		BlockHash:   common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f1"),
		NetworkID:   1,
	}

	blockID, err := testStore.AddBlock(ctx, block, tx)
	require.NoError(t, err)
	require.Equal(t, blockID, uint64(1))

	l2GER := &etherman.GlobalExitRoot{
		NetworkID:      1,
		BlockNumber:    1,
		BlockID:        1,
		GlobalExitRoot: common.HexToHash("0x29e885edaf8e4b51e1d2e05f9da28161d2fb4f6b1d53827d9b80a23cf2d7d9f1"),
	}

	err = testStore.AddGlobalExitRoot(ctx, l2GER, tx)
	require.NoError(t, err)

	_, err = testStore.GetLatestTrustedExitRoot(ctx, l2GER.NetworkID, tx)
	require.Error(t, err)
	require.NoError(t, testStore.Commit(ctx, tx))
}

type testStore interface {
	AddDeposit(ctx context.Context, deposit *etherman.Deposit, dbTx interface{}) (uint64, error)
	Rollback(ctx context.Context, dbTx interface{}) error
	BeginDBTransaction(ctx context.Context) (interface{}, error)
	AddBlock(ctx context.Context, block *etherman.Block, dbTx interface{}) (uint64, error)
	Commit(ctx context.Context, dbTx interface{}) error
	AddGlobalExitRoot(ctx context.Context, exitRoot *etherman.GlobalExitRoot, dbTx interface{}) error
	GetLatestL1SyncedExitRoot(ctx context.Context, dbTx interface{}) (*etherman.GlobalExitRoot, error)
	GetLatestExitRoot(ctx context.Context, networkID, destNetwork uint32, dbTx interface{}) (*etherman.GlobalExitRoot, error)
	AddTrustedGlobalExitRoot(_ context.Context, trustedExitRoot *etherman.GlobalExitRoot, dbTx interface{}) (bool, error)
	GetLatestTrustedExitRoot(ctx context.Context, networkID uint32, dbTx interface{}) (*etherman.GlobalExitRoot, error)
	QueryRowTesting(ctx context.Context, data string, dbTx interface{}) interface{}
	GetLastBlock(ctx context.Context, networkID uint32, dbTx interface{}) (*etherman.Block, error)
	GetPreviousBlock(ctx context.Context, networkID uint32, offset uint64, dbTx interface{}) (etherman.Block, error)
	// ExecTesting(ctx context.Context, data string) error
	SetRoot(ctx context.Context, root []byte, depositID uint64, network uint32, dbTx interface{}) error
	Set(ctx context.Context, key []byte, value [][]byte, depositID uint64, dbTx interface{}) error
	Get(ctx context.Context, key []byte, dbTx interface{}) ([][]byte, error)
	GetRoot(ctx context.Context, depositCnt, network uint32, dbTx interface{}) ([]byte, error)
	GetLastDepositCount(ctx context.Context, networkID uint32, dbTx interface{}) (uint32, error)
	GetDepositCountByRoot(ctx context.Context, root []byte, network uint32, dbTx interface{}) (uint32, error)
	AddClaim(ctx context.Context, claim *etherman.Claim, dbTx interface{}) error
	GetClaim(_ context.Context, depositCount, originNetworkID, networkID uint32, dbTx interface{}) (*etherman.Claim, error)
	GetClaims(ctx context.Context, destAddr string, limit, offset uint32, dbTx interface{}) ([]*etherman.Claim, error)
	GetDepositCount(ctx context.Context, destAddr string, dbTx interface{}) (uint64, error)
	GetDeposits(ctx context.Context, destAddr string, limit, offset uint32, dbTx interface{}) ([]*etherman.Deposit, error)
	GetDeposit(ctx context.Context, depositCounterUser, networkID uint32, dbTx interface{}) (*etherman.Deposit, error)
	GetNumberDeposits(ctx context.Context, networkID uint32, blockNumber uint64, dbTx interface{}) (uint32, error)
	GetClaimCount(ctx context.Context, destAddr string, dbTx interface{}) (uint64, error)
	GetTokenMetadata(ctx context.Context, networkID, destNet uint32, originalTokenAddr common.Address, dbTx interface{}) ([]byte, error)
	AddTokenWrapped(ctx context.Context, tokenWrapped *etherman.TokenWrapped, dbTx interface{}) error
	GetTokenWrapped(ctx context.Context, originalNetwork uint32, originalTokenAddress common.Address, dbTx interface{}) (*etherman.TokenWrapped, error)
}

func newStorageSettings(storageType string) (testStore, error) {
	if storageType == "postgres" {
		dbCfg := pgstorage.NewConfigFromEnv()
		err := pgstorage.InitOrReset(dbCfg)
		if err != nil {
			return nil, err
		}
		mt, err := pgstorage.NewPostgresStorage(dbCfg)
		return mt, err
	}
	return nil, fmt.Errorf("unknown storage type: %s", storageType)
}
