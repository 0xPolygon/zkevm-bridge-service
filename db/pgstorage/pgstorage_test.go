package pgstorage

import (
	"context"
	"math"
	"math/big"
	"testing"
	"time"

	ctmtypes "github.com/0xPolygon/zkevm-bridge-service/claimtxman/types"
	"github.com/0xPolygon/zkevm-bridge-service/etherman"
	"github.com/0xPolygon/zkevm-bridge-service/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetLeaves(t *testing.T) {
	data := `INSERT INTO sync.block
	(id, block_num, block_hash, network_id)
	VALUES(1, 1, decode('5C7831','hex'), 0);
	INSERT INTO sync.block
	(id, block_num, block_hash, network_id)
	VALUES(2, 2, decode('5C7832','hex'), 0);
	
	INSERT INTO mt.rollup_exit
	(leaf, rollup_id, root, block_id)
	VALUES(decode('A4BFA0908DC7B06D98DA4309F859023D6947561BC19BC00D77F763DEA1A0B9F5','hex'), 1, decode('42D3339FE8EB57770953423F20A029E778A707E8D58AAF110B40D5EB4DD25721','hex'), 1);
	INSERT INTO mt.rollup_exit
	(leaf, rollup_id, root, block_id)
	VALUES(decode('315FEE1AA202BF4A6BD0FDE560C89BE90B6E6E2AAF92DC5E8D118209ABC3410F','hex'), 2, decode('42D3339FE8EB57770953423F20A029E778A707E8D58AAF110B40D5EB4DD25721','hex'), 1);
	INSERT INTO mt.rollup_exit
	(leaf, rollup_id, root, block_id)
	VALUES(decode('B598CE65AA15C08DDA126A2985BA54F0559EAAC562BB43BA430C7344261FBC5D','hex'), 3, decode('42D3339FE8EB57770953423F20A029E778A707E8D58AAF110B40D5EB4DD25721','hex'), 1);
	INSERT INTO mt.rollup_exit
	(leaf, rollup_id, root, block_id)
	VALUES(decode('E6585BDF74B6A46B9EDE8B1B877E1232FB79EE93106C4DB8FFD49CF1685BF242','hex'), 4, decode('42D3339FE8EB57770953423F20A029E778A707E8D58AAF110B40D5EB4DD25721','hex'), 1);
	
	INSERT INTO mt.rollup_exit
	(leaf, rollup_id, root, block_id)
	VALUES(decode('A4BFA0908DC7B06D98DA4309F859023D6947561BC19BC00D77F763DEA1A0B9F6','hex'), 1, decode('42D3339FE8EB57770953423F20A029E778A707E8D58AAF110B40D5EB4DD25722','hex'), 2);
	INSERT INTO mt.rollup_exit
	(leaf, rollup_id, root, block_id)
	VALUES(decode('E6585BDF74B6A46B9EDE8B1B877E1232FB79EE93106C4DB8FFD49CF1685BF243','hex'), 4, decode('42D3339FE8EB57770953423F20A029E778A707E8D58AAF110B40D5EB4DD25722','hex'), 2);
	INSERT INTO mt.rollup_exit
	(leaf, rollup_id, root, block_id)
	VALUES(decode('B598CE65AA15C08DDA126A2985BA54F0559EAAC562BB43BA430C7344261FBC5E','hex'), 3, decode('42D3339FE8EB57770953423F20A029E778A707E8D58AAF110B40D5EB4DD25722','hex'), 2);
	INSERT INTO mt.rollup_exit
	(leaf, rollup_id, root, block_id)
	VALUES(decode('315FEE1AA202BF4A6BD0FDE560C89BE90B6E6E2AAF92DC5E8D118209ABC34110','hex'), 2, decode('42D3339FE8EB57770953423F20A029E778A707E8D58AAF110B40D5EB4DD25722','hex'), 1);
	`
	dbCfg := NewConfigFromEnv()
	ctx := context.Background()
	err := InitOrReset(ctx, dbCfg)
	require.NoError(t, err)

	store, err := NewPostgresStorage(ctx, dbCfg)
	require.NoError(t, err)
	_, err = store.Exec(ctx, data)
	require.NoError(t, err)

	leaves, err := store.GetLatestRollupExitLeaves(ctx, nil)
	require.NoError(t, err)

	for _, l := range leaves {
		log.Debugf("leaves: %+v", l)
	}
	assert.Equal(t, "0xa4bfa0908dc7b06d98da4309f859023d6947561bc19bc00d77f763dea1a0b9f6", leaves[0].Leaf.String())
	assert.Equal(t, uint64(5), leaves[0].ID)
	assert.Equal(t, uint64(2), leaves[0].BlockID)
	assert.Equal(t, uint32(1), leaves[0].RollupId)
	assert.Equal(t, "0x42d3339fe8eb57770953423f20a029e778a707e8d58aaf110b40d5eb4dd25722", leaves[0].Root.String())

	assert.Equal(t, "0x315fee1aa202bf4a6bd0fde560c89be90b6e6e2aaf92dc5e8d118209abc34110", leaves[1].Leaf.String())
	assert.Equal(t, uint64(8), leaves[1].ID)
	assert.Equal(t, uint64(1), leaves[1].BlockID)
	assert.Equal(t, uint32(2), leaves[1].RollupId)
	assert.Equal(t, "0x42d3339fe8eb57770953423f20a029e778a707e8d58aaf110b40d5eb4dd25722", leaves[1].Root.String())

	assert.Equal(t, "0xb598ce65aa15c08dda126a2985ba54f0559eaac562bb43ba430c7344261fbc5e", leaves[2].Leaf.String())
	assert.Equal(t, uint64(7), leaves[2].ID)
	assert.Equal(t, uint64(2), leaves[2].BlockID)
	assert.Equal(t, uint32(3), leaves[2].RollupId)
	assert.Equal(t, "0x42d3339fe8eb57770953423f20a029e778a707e8d58aaf110b40d5eb4dd25722", leaves[2].Root.String())

	assert.Equal(t, "0xe6585bdf74b6a46b9ede8b1b877e1232fb79ee93106c4db8ffd49cf1685bf243", leaves[3].Leaf.String())
	assert.Equal(t, uint64(6), leaves[3].ID)
	assert.Equal(t, uint64(2), leaves[3].BlockID)
	assert.Equal(t, uint32(4), leaves[3].RollupId)
	assert.Equal(t, "0x42d3339fe8eb57770953423f20a029e778a707e8d58aaf110b40d5eb4dd25722", leaves[3].Root.String())
}

func TestIsRollupExitRoot(t *testing.T) {
	data := `INSERT INTO sync.block
	(id, block_num, block_hash, network_id)
	VALUES(1, 1, decode('5C7831','hex'), 0);
	
	INSERT INTO mt.rollup_exit
	(leaf, rollup_id, root, block_id)
	VALUES(decode('A4BFA0908DC7B06D98DA4309F859023D6947561BC19BC00D77F763DEA1A0B9F5','hex'), 1, decode('42D3339FE8EB57770953423F20A029E778A707E8D58AAF110B40D5EB4DD25721','hex'), 1);
	`
	root := common.HexToHash("0x42D3339FE8EB57770953423F20A029E778A707E8D58AAF110B40D5EB4DD25721")
	dbCfg := NewConfigFromEnv()
	ctx := context.Background()
	err := InitOrReset(ctx, dbCfg)
	require.NoError(t, err)

	store, err := NewPostgresStorage(ctx, dbCfg)
	require.NoError(t, err)

	exist, err := store.IsRollupExitRoot(ctx, root, nil)
	require.NoError(t, err)
	assert.Equal(t, false, exist)

	_, err = store.Exec(ctx, data)
	require.NoError(t, err)

	exist, err = store.IsRollupExitRoot(ctx, root, nil)
	require.NoError(t, err)

	assert.Equal(t, true, exist)
}

func createStore(t *testing.T) *PostgresStorage {
	dbCfg := NewConfigFromEnv()
	ctx := context.Background()
	err := InitOrReset(ctx, dbCfg)
	require.NoError(t, err)
	store, err := NewPostgresStorage(ctx, dbCfg)
	require.NoError(t, err)
	return store
}

func TestAddMonitoredTxs(t *testing.T) {
	store := createStore(t)
	ctx := context.Background()
	to := common.HexToAddress("0x123")
	groupID := uint64(1)
	monitoredTx := ctmtypes.MonitoredTx{
		To:        &to,
		Nonce:     1,
		Value:     nil,
		Data:      []byte("data"),
		Gas:       100,
		GasPrice:  nil,
		Status:    ctmtypes.MonitoredTxStatusCreated,
		History:   map[common.Hash]bool{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		GroupID:   &groupID,
	}
	err := store.AddClaimTx(ctx, monitoredTx, nil)
	require.NoError(t, err)
}

func TestAddMonitoredTxsGroup(t *testing.T) {
	store := createStore(t)
	ctx := context.Background()
	group := ctmtypes.MonitoredTxGroupDBEntry{
		GroupID:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := store.AddMonitoredTxsGroup(ctx, &group, nil)
	require.NoError(t, err)
	require.Equal(t, uint64(1), group.GroupID)
}

func TestGetPendingDepositsToClaim(t *testing.T) {
	data := `INSERT INTO sync.block
	(id, block_num, block_hash, network_id)
	VALUES(1, 1, decode('5C7831','hex'), 0);
	
	INSERT INTO sync.deposit
	(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim)
	VALUES(0, 0, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', 1, decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, 0, decode('CBE7A77275EE22780BB94EA900D42CEF88F5A2F0E1A7C76696556D7FF17767E6','hex'), decode('','hex'), 1, true);
	INSERT INTO sync.deposit
	(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim)
	VALUES(0, 0, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', 1, decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, 1, decode('6282FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE1F','hex'), decode('','hex'), 2, true);
	INSERT INTO sync.deposit
	(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim)
	VALUES(0, 0, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', 1, decode('F38FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, 2, decode('6282FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE1F','hex'), decode('','hex'), 3, true);
	INSERT INTO sync.claim
	(network_id, "index", orig_net, orig_addr, amount, dest_addr, block_id, tx_hash, rollup_index, mainnet_flag)
	VALUES(1, 0, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, decode('BF2C816AB6F8A8F5F9DDA6EE97D433CC841E69B5669A5CDF499826FA4B99C179','hex'), 0, true);
	`
	dbCfg := NewConfigFromEnv()
	ctx := context.Background()
	err := InitOrReset(ctx, dbCfg)
	require.NoError(t, err)

	store, err := NewPostgresStorage(ctx, dbCfg)
	require.NoError(t, err)

	_, err = store.Exec(ctx, data)
	require.NoError(t, err)
	deposits, totalCount, err := store.GetPendingDepositsToClaim(ctx, common.Address{}, 1, 0, 2, 0, -1, nil)
	require.NoError(t, err)
	assert.Equal(t, 2, len(deposits))
	assert.Equal(t, uint64(2), totalCount)
	assert.Equal(t, uint8(0), deposits[0].LeafType)
	assert.Equal(t, uint32(0), deposits[0].NetworkID)
	assert.Equal(t, uint32(0), deposits[0].OriginalNetwork)
	assert.Equal(t, common.Address{}, deposits[0].OriginalAddress)
	assert.Equal(t, big.NewInt(90000000000000000), deposits[0].Amount)
	assert.Equal(t, uint32(1), deposits[0].DestinationNetwork)
	assert.Equal(t, common.HexToAddress("0xF39FD6E51AAD88F6F4CE6AB8827279CFFFB92266"), deposits[0].DestinationAddress)
	assert.Equal(t, uint64(1), deposits[0].BlockID)
	assert.Equal(t, uint32(1), deposits[0].DepositCount)
	assert.Equal(t, common.HexToHash("0x6282FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE1F"), deposits[0].TxHash)
	assert.Equal(t, []byte{}, deposits[0].Metadata)
	assert.Equal(t, uint64(2), deposits[0].Id)
	assert.Equal(t, true, deposits[0].ReadyForClaim)
	assert.Equal(t, uint8(0), deposits[1].LeafType)
	assert.Equal(t, uint32(0), deposits[1].NetworkID)
	assert.Equal(t, uint32(0), deposits[1].OriginalNetwork)
	assert.Equal(t, common.Address{}, deposits[1].OriginalAddress)
	assert.Equal(t, big.NewInt(90000000000000000), deposits[1].Amount)
	assert.Equal(t, uint32(1), deposits[1].DestinationNetwork)
	assert.Equal(t, common.HexToAddress("0xF38FD6E51AAD88F6F4CE6AB8827279CFFFB92266"), deposits[1].DestinationAddress)
	assert.Equal(t, uint64(1), deposits[1].BlockID)
	assert.Equal(t, uint32(2), deposits[1].DepositCount)
	assert.Equal(t, common.HexToHash("0x6282FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE1F"), deposits[1].TxHash)
	assert.Equal(t, []byte{}, deposits[1].Metadata)
	assert.Equal(t, uint64(3), deposits[1].Id)
	assert.Equal(t, true, deposits[1].ReadyForClaim)

	deposits, totalCount, err = store.GetPendingDepositsToClaim(ctx, common.HexToAddress("0xF39FD6E51AAD88F6F4CE6AB8827279CFFFB92266"), 1, 0, 0, 0, 0, nil)
	require.NoError(t, err)
	assert.Equal(t, 1, len(deposits))
	assert.Equal(t, uint64(1), totalCount)
	assert.Equal(t, uint8(0), deposits[0].LeafType)
	assert.Equal(t, uint32(0), deposits[0].NetworkID)
	assert.Equal(t, uint32(0), deposits[0].OriginalNetwork)
	assert.Equal(t, common.Address{}, deposits[0].OriginalAddress)
	assert.Equal(t, big.NewInt(90000000000000000), deposits[0].Amount)
	assert.Equal(t, uint32(1), deposits[0].DestinationNetwork)
	assert.Equal(t, common.HexToAddress("0xF39FD6E51AAD88F6F4CE6AB8827279CFFFB92266"), deposits[0].DestinationAddress)
	assert.Equal(t, uint64(1), deposits[0].BlockID)
	assert.Equal(t, uint32(1), deposits[0].DepositCount)
	assert.Equal(t, common.HexToHash("0x6282FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE1F"), deposits[0].TxHash)
	assert.Equal(t, []byte{}, deposits[0].Metadata)
	assert.Equal(t, uint64(2), deposits[0].Id)
	assert.Equal(t, true, deposits[0].ReadyForClaim)

	err = store.IgnoreDeposit(ctx, 2, nil)
	assert.NoError(t, err)
	deposits, totalCount, err = store.GetPendingDepositsToClaim(ctx, common.Address{}, 1, 0, 2, 0, -1, nil)
	require.NoError(t, err)
	assert.Equal(t, 1, len(deposits))
	assert.Equal(t, uint64(1), totalCount)
	assert.Equal(t, uint8(0), deposits[0].LeafType)
	assert.Equal(t, uint32(0), deposits[0].NetworkID)
	assert.Equal(t, uint32(0), deposits[0].OriginalNetwork)
	assert.Equal(t, common.Address{}, deposits[0].OriginalAddress)
	assert.Equal(t, big.NewInt(90000000000000000), deposits[0].Amount)
	assert.Equal(t, uint32(1), deposits[0].DestinationNetwork)
	assert.Equal(t, common.HexToAddress("0xF38FD6E51AAD88F6F4CE6AB8827279CFFFB92266"), deposits[0].DestinationAddress)
	assert.Equal(t, uint64(1), deposits[0].BlockID)
	assert.Equal(t, uint32(2), deposits[0].DepositCount)
	assert.Equal(t, common.HexToHash("0x6282FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE1F"), deposits[0].TxHash)
	assert.Equal(t, []byte{}, deposits[0].Metadata)
	assert.Equal(t, uint64(3), deposits[0].Id)
	assert.Equal(t, true, deposits[0].ReadyForClaim)
}

func TestResetDeposits(t *testing.T) {
	data := `INSERT INTO sync.block
	(id, block_num, block_hash, network_id)
	VALUES(1, 1, decode('5C7831','hex'), 0);
	INSERT INTO sync.block
	(id, block_num, block_hash, network_id)
	VALUES(2, 1, decode('6C7831','hex'), 1);
	
	INSERT INTO sync.deposit
	(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim)
	VALUES(0, 0, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', 1, decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, 0, decode('CBE7A77275EE22780BB94EA900D42CEF88F5A2F0E1A7C76696556D7FF17767E6','hex'), decode('','hex'), 1, true);
	INSERT INTO sync.deposit
	(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim)
	VALUES(0, 0, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', 1, decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, 1, decode('6282FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE1F','hex'), decode('','hex'), 2, true);
	INSERT INTO sync.deposit
	(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim)
	VALUES(0, 0, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', 1, decode('F38FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, 2, decode('6282FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE1F','hex'), decode('','hex'), 3, true);
	INSERT INTO sync.claim
	(network_id, "index", orig_net, orig_addr, amount, dest_addr, block_id, tx_hash, rollup_index, mainnet_flag)
	VALUES(1, 0, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, decode('BF2C816AB6F8A8F5F9DDA6EE97D433CC841E69B5669A5CDF499826FA4B99C179','hex'), 0, true);
	INSERT INTO sync.claim
	(network_id, "index", orig_net, orig_addr, amount, dest_addr, block_id, tx_hash, rollup_index, mainnet_flag)
	VALUES(1, 1, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, decode('BF2C816AB6F8A8F5F9DDA6EE97D433CC841E69B5669A5CDF499826FA4B99C179','hex'), 0, true);

	INSERT INTO sync.deposit
	(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim)
	VALUES(0, 1, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', 0, decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 2, 0, decode('ABE7A77275EE22780BB94EA900D42CEF88F5A2F0E1A7C76696556D7FF17767E6','hex'), decode('','hex'), 4, true);
	INSERT INTO sync.deposit
	(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim)
	VALUES(0, 1, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', 0, decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 2, 1, decode('A282FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE1F','hex'), decode('','hex'), 5, true);
	INSERT INTO sync.deposit
	(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim)
	VALUES(0, 1, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', 0, decode('F38FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 2, 2, decode('B282FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE1F','hex'), decode('','hex'), 6, true);
	INSERT INTO sync.claim
	(network_id, "index", orig_net, orig_addr, amount, dest_addr, block_id, tx_hash, rollup_index, mainnet_flag)
	VALUES(0, 0, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 2, decode('CF2C816AB6F8A8F5F9DDA6EE97D433CC841E69B5669A5CDF499826FA4B99C179','hex'), 0, true);
	INSERT INTO sync.claim
	(network_id, "index", orig_net, orig_addr, amount, dest_addr, block_id, tx_hash, rollup_index, mainnet_flag)
	VALUES(0, 1, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 2, decode('DF2C816AB6F8A8F5F9DDA6EE97D433CC841E69B5669A5CDF499826FA4B99C179','hex'), 0, true);
	`
	dbCfg := NewConfigFromEnv()
	ctx := context.Background()
	err := InitOrReset(ctx, dbCfg)
	require.NoError(t, err)

	store, err := NewPostgresStorage(ctx, dbCfg)
	require.NoError(t, err)

	_, err = store.Exec(ctx, data)
	require.NoError(t, err)


	err = store.ResetDeposits(ctx, 1, 0, 1, nil)
	require.Error(t, err)
	d, err := store.GetNumberDeposits(ctx, 1, 1, nil)
	require.NoError(t, err)
	assert.Equal(t, uint32(3), d)
	err = store.ResetDeposits(ctx, 1, 1, 1, nil)
	require.NoError(t, err)
	d, err = store.GetNumberDeposits(ctx, 1, 1, nil)
	require.NoError(t, err)
	assert.Equal(t, uint32(1), d)

	// Verify that deposits were stored in order in deposit_backup table
	var depositCounts []uint32
	rows, err := store.Query(ctx, "SELECT deposit_cnt FROM sync.deposit_backup WHERE backward_let_id = 1 ORDER BY id ASC")
	require.NoError(t, err)
	defer rows.Close()

	for rows.Next() {
		var depositCnt uint32
		err = rows.Scan(&depositCnt)
		require.NoError(t, err)
		depositCounts = append(depositCounts, depositCnt)
	}
	require.NoError(t, rows.Err())

	// Check that we have 2 deposits (deposit_cnt 1 and 2)
	require.Len(t, depositCounts, 2)

	// Verify they are in ascending order
	assert.Equal(t, uint32(1), depositCounts[0], "First deposit should have deposit_cnt = 1")
	assert.Equal(t, uint32(2), depositCounts[1], "Second deposit should have deposit_cnt = 2")
}

func TestAddBackwardLET(t *testing.T) {
	data := `INSERT INTO sync.block
	(id, block_num, block_hash, network_id)
	VALUES(1, 1, decode('5C7831','hex'), 0);
	`
	dbCfg := NewConfigFromEnv()
	ctx := context.Background()
	err := InitOrReset(ctx, dbCfg)
	require.NoError(t, err)

	store, err := NewPostgresStorage(ctx, dbCfg)
	require.NoError(t, err)

	_, err = store.Exec(ctx, data)
	require.NoError(t, err)

	bLET := &etherman.BackwardLET{
		BlockID: 1,
    	PreviousDepositCount: 3,
    	PreviousRoot: common.HexToHash("0x6282FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE1F"),
    	NewDepositCount: 1,
    	NewRoot: common.HexToHash("0xAA82FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEECC"),
	}
	id, err := store.AddBackwardLET(ctx, bLET, nil)
	require.NoError(t, err)
	assert.NotZero(t, id)

	var count uint32
	selectCount := "SELECT count(*) FROM sync.backward_let"
	err = store.QueryRow(ctx, selectCount).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), count)
}

func TestAddForwardLET(t *testing.T) {
	data := `INSERT INTO sync.block
	(id, block_num, block_hash, network_id)
	VALUES(1, 1, decode('5C7831','hex'), 0);
	`
	dbCfg := NewConfigFromEnv()
	ctx := context.Background()
	err := InitOrReset(ctx, dbCfg)
	require.NoError(t, err)

	store, err := NewPostgresStorage(ctx, dbCfg)
	require.NoError(t, err)

	_, err = store.Exec(ctx, data)
	require.NoError(t, err)

	fLET := &etherman.ForwardLET{
		BlockID: 1,
    	PreviousDepositCount: 3,
    	PreviousRoot: common.HexToHash("0x6282FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE1F"),
    	NewDepositCount: 1,
    	NewRoot: common.HexToHash("0xAA82FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEECC"),
		NewRawLeaves: []byte("Random bytes to fill the db field"),
	}
	err = store.AddForwardLET(ctx, fLET, nil)
	require.NoError(t, err)

	var count uint32
	selectCount := "SELECT count(*) FROM sync.forward_let"
	err = store.QueryRow(ctx, selectCount).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), count)
}

func TestAddSetUnsetClaim(t *testing.T) {
	data := `INSERT INTO sync.block
	(id, block_num, block_hash, network_id)
	VALUES(1, 1, decode('5C7831','hex'), 0);
	`
	dbCfg := NewConfigFromEnv()
	ctx := context.Background()
	err := InitOrReset(ctx, dbCfg)
	require.NoError(t, err)

	store, err := NewPostgresStorage(ctx, dbCfg)
	require.NoError(t, err)

	_, err = store.Exec(ctx, data)
	require.NoError(t, err)

	globalIndex, _ := big.NewInt(0).SetString("18446744073709551615", 0)
	setClaim := &etherman.SetClaim{
		BlockID: 1,
    	MainnetFlag: false,
    	RollupIndex: math.MaxUint32,
    	Index: math.MaxUint32,
    	GlobalIndex: globalIndex,
	}
	err = store.AddSetClaim(ctx, setClaim, nil)
	require.NoError(t, err)

	globalIndex, _ = big.NewInt(0).SetString("18446744073709551627", 0)
	unsetClaim := &etherman.UnsetClaim{
		BlockID: 1,
    	MainnetFlag: true,
    	RollupIndex: 0,
    	Index: 11,
    	GlobalIndex: globalIndex,
	}
	err = store.AddUnsetClaim(ctx, unsetClaim, nil)
	require.NoError(t, err)

	var count uint32
	selectCount := "SELECT count(*) FROM sync.set_unset_claim"
	err = store.QueryRow(ctx, selectCount).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, uint32(2), count)
}

func TestAddSyncStatus(t *testing.T) {
	store := createStore(t)
	ctx := context.Background()
	status1 := etherman.SyncStatus{
		NetworkID:       0,
		Percentage:      50,
		RemainingBlocks: 1000,
		Synced:          false,
	}
	err := store.AddSyncStatus(ctx, status1, nil)
	require.NoError(t, err)
	status1bis := etherman.SyncStatus{
		NetworkID:       0,
		Percentage:      51,
		RemainingBlocks: 1000,
		Synced:          false,
	}
	err = store.AddSyncStatus(ctx, status1bis, nil)
	require.NoError(t, err)
	status2 := etherman.SyncStatus{
		NetworkID:       1,
		Percentage:      99,
		RemainingBlocks: 23,
		Synced:          false,
	}
	err = store.AddSyncStatus(ctx, status2, nil)
	require.NoError(t, err)
	status3 := etherman.SyncStatus{
		NetworkID:       2,
		Percentage:      100,
		RemainingBlocks: 0,
		Synced:          true,
	}
	err = store.AddSyncStatus(ctx, status3, nil)
	require.NoError(t, err)
	status, err := store.GetSyncStatus(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, 3, len(status))
	assert.Equal(t, uint32(0), status[0].NetworkID)
	assert.Equal(t, uint32(51), status[0].Percentage)
	assert.Equal(t, uint64(1000), status[0].RemainingBlocks)
	assert.Equal(t, false, status[0].Synced)
	assert.Equal(t, uint32(1), status[1].NetworkID)
	assert.Equal(t, uint32(99), status[1].Percentage)
	assert.Equal(t, uint64(23), status[1].RemainingBlocks)
	assert.Equal(t, false, status[1].Synced)
	assert.Equal(t, uint32(2), status[2].NetworkID)
	assert.Equal(t, uint32(100), status[2].Percentage)
	assert.Equal(t, uint64(0), status[2].RemainingBlocks)
	assert.Equal(t, true, status[2].Synced)
}

func TestDeleteClaimByGlobalIndex(t *testing.T) {
	data := `INSERT INTO sync.block
	(id, block_num, block_hash, network_id)
	VALUES(1, 1, decode('5C7831','hex'), 0);
	
	INSERT INTO sync.deposit
	(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim)
	VALUES(0, 0, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', 1, decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, 0, decode('CBE7A77275EE22780BB94EA900D42CEF88F5A2F0E1A7C76696556D7FF17767E6','hex'), decode('','hex'), 1, true);
	INSERT INTO sync.deposit
	(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim)
	VALUES(0, 0, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', 1, decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, 1, decode('6282FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE1F','hex'), decode('','hex'), 2, true);
	INSERT INTO sync.deposit
	(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim)
	VALUES(0, 0, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', 1, decode('F38FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, 2, decode('6282FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE1F','hex'), decode('','hex'), 3, true);
	INSERT INTO sync.claim
	(network_id, "index", orig_net, orig_addr, amount, dest_addr, block_id, tx_hash, rollup_index, mainnet_flag, global_index)
	VALUES(1, 0, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, decode('BF2C816AB6F8A8F5F9DDA6EE97D433CC841E69B5669A5CDF499826FA4B99C179','hex'), 0, true, 12331);
	INSERT INTO sync.claim
	(network_id, "index", orig_net, orig_addr, amount, dest_addr, block_id, tx_hash, rollup_index, mainnet_flag, global_index)
	VALUES(1, 1, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, decode('BF2C816AB6F8A8F5F9DDA6EE97D433CC841E69B5669A5CDF499826FA4B99C179','hex'), 0, true, 18446744073709551616);
	INSERT INTO sync.claim
	(network_id, "index", orig_net, orig_addr, amount, dest_addr, block_id, tx_hash, rollup_index, mainnet_flag, global_index)
	VALUES(1, 2, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, decode('BF2C816AB6F8A8F5F9DDA6EE97D433CC841E69B5669A5CDF499826FA4B99C179','hex'), 0, true, 5432312);
	`
	dbCfg := NewConfigFromEnv()
	ctx := context.Background()
	err := InitOrReset(ctx, dbCfg)
	require.NoError(t, err)

	store, err := NewPostgresStorage(ctx, dbCfg)
	require.NoError(t, err)

	_, err = store.Exec(ctx, data)
	require.NoError(t, err)
	
	globalIndex, _ := big.NewInt(0).SetString("18446744073709551616", 0)
	err = store.DeleteClaimByGlobalIndex(ctx, globalIndex, 1, nil)
	require.NoError(t, err)

	var count uint32
	selectCount := "SELECT count(*) FROM sync.claim"
	err = store.QueryRow(ctx, selectCount).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, uint32(2), count)
}

func TestGetAndDeleteOrphanDepositBackups(t *testing.T) {
	dbCfg := NewConfigFromEnv()
	ctx := context.Background()
	err := InitOrReset(ctx, dbCfg)
	require.NoError(t, err)

	store, err := NewPostgresStorage(ctx, dbCfg)
	require.NoError(t, err)

	// Setup: Create blocks
	// Blocks 1-6: Original blocks where deposits were synced (deposit_cnt 0-5)
	// Block 7: First backward_let event
	// Block 8: Second backward_let event
	blocksData := `
	INSERT INTO sync.block (id, block_num, block_hash, network_id)
	VALUES(1, 100, decode('AA01033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C01','hex'), 1),
	      (2, 101, decode('AA02033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C02','hex'), 1),
	      (3, 102, decode('AA03033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C03','hex'), 1),
	      (4, 103, decode('AA04033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C04','hex'), 1),
	      (5, 104, decode('AA05033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C05','hex'), 1),
	      (6, 105, decode('AA06033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C06','hex'), 1),
	      (7, 200, decode('AAAA033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C07','hex'), 1),
	      (8, 201, decode('BBBB033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C08','hex'), 1);
	`
	_, err = store.Exec(ctx, blocksData)
	require.NoError(t, err)

	// Simulate: Originally we had 10 deposits (deposit_cnt 0-9)
	// First backward_let: previous_deposit_cnt=9, new_deposit_cnt=7 (deleted deposits 8,9)
	firstBackwardLET := &etherman.BackwardLET{
		BlockID:              7,
		PreviousDepositCount: 9,
		PreviousRoot:         common.HexToHash("0x1111FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE1F"),
		NewDepositCount:      7,
		NewRoot:              common.HexToHash("0x2222FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE2F"),
	}
	backwardLetID1, err := store.AddBackwardLET(ctx, firstBackwardLET, nil)
	require.NoError(t, err)
	require.Equal(t, uint64(1), backwardLetID1)

	// Second backward_let: previous_deposit_cnt=7, new_deposit_cnt=2 (deleted deposits 3,4,5,6,7)
	secondBackwardLET := &etherman.BackwardLET{
		BlockID:              8,
		PreviousDepositCount: 7,
		PreviousRoot:         common.HexToHash("0x2222FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE2F"),
		NewDepositCount:      2,
		NewRoot:              common.HexToHash("0x3333FACE883070640F802CE8A2C42593AA18D3A691C61BA006EC477D6E5FEE3F"),
	}
	backwardLetID2, err := store.AddBackwardLET(ctx, secondBackwardLET, nil)
	require.NoError(t, err)
	require.Equal(t, uint64(2), backwardLetID2)

	// Insert deposit_backup entries for both backward_let events
	// Deposits backed up by first backward_let (backward_let_id=1)
	// These deposit_backups has the block_id of the original deposit. I mean the block_id where the deposit was synced. Not the block_id of the block where the backward_let event was synced
	depositBackupData1 := `
	INSERT INTO sync.deposit_backup(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, ready_for_claim, backward_let_id)
	VALUES(0, 1, 0, decode('A9B5033799ADF3739383A0489EFBE8A0D4D5E447','hex'), '8000000000000000000', 0, decode('B9B5033799ADF3739383A0489EFBE8A0D4D5E458','hex'), 1, 0, decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C10','hex'), decode('','hex'), true, 1),
	      (0, 1, 0, decode('A9B5033799ADF3739383A0489EFBE8A0D4D5E448','hex'), '9000000000000000000', 0, decode('B9B5033799ADF3739383A0489EFBE8A0D4D5E459','hex'), 2, 1, decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C11','hex'), decode('','hex'), true, 1);
	`
	_, err = store.Exec(ctx, depositBackupData1)
	require.NoError(t, err)

	// Deposits backed up by second backward_let (backward_let_id=2)
	depositBackupData2 := `
	INSERT INTO sync.deposit_backup(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, ready_for_claim, backward_let_id)
	VALUES(0, 1, 0, decode('A9B5033799ADF3739383A0489EFBE8A0D4D5E546','hex'), '3000000000000000000', 0, decode('B9B5033799ADF3739383A0489EFBE8A0D4D5E481','hex'), 3, 2, decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C12','hex'), decode('','hex'), true, 2);

	INSERT INTO sync.deposit_backup(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, ready_for_claim, backward_let_id)
	VALUES(0, 1, 0, decode('A9B5033799ADF3739383A0489EFBE8A0D4D5E547','hex'), '4000000000000000000', 0, decode('B9B5033799ADF3739383A0489EFBE8A0D4D5E482','hex'), 4, 3, decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C13','hex'), decode('','hex'), true, 2);

	INSERT INTO sync.deposit_backup(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, ready_for_claim, backward_let_id)
	VALUES(0, 1, 0, decode('A9B5033799ADF3739383A0489EFBE8A0D4D5E548','hex'), '5000000000000000000', 0, decode('B9B5033799ADF3739383A0489EFBE8A0D4D5E483','hex'), 5, 4, decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C14','hex'), decode('','hex'), false, 2);

	INSERT INTO sync.deposit_backup(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, ready_for_claim, backward_let_id)
	VALUES(0, 1, 0, decode('A9B5033799ADF3739383A0489EFBE8A0D4D5E549','hex'), '6000000000000000000', 0, decode('B9B5033799ADF3739383A0489EFBE8A0D4D5E484','hex'), 6, 5, decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C15','hex'), decode('','hex'), true, 2);
	`
	_, err = store.Exec(ctx, depositBackupData2)
	require.NoError(t, err)

	// Verify initial state: should have 6 deposit_backup entries (2 from first backward_let + 4 from second)
	var initialCount uint32
	err = store.QueryRow(ctx, "SELECT count(*) FROM sync.deposit_backup").Scan(&initialCount)
	require.NoError(t, err)
	assert.Equal(t, uint32(6), initialCount, "Should have 6 deposit_backup entries initially")

	// Verify we have 2 backward_let entries
	var backwardLetCount uint32
	err = store.QueryRow(ctx, "SELECT count(*) FROM sync.backward_let").Scan(&backwardLetCount)
	require.NoError(t, err)
	assert.Equal(t, uint32(2), backwardLetCount, "Should have 2 backward_let entries")

	// Now delete block 8 which should trigger CASCADE delete on second backward_let
	// This will make the 4 deposits associated with backward_let_id=2 orphans
	deleteBlock := "DELETE FROM sync.block WHERE id = 8"
	_, err = store.Exec(ctx, deleteBlock)
	require.NoError(t, err)

	// Verify backward_let with id=2 was deleted (CASCADE)
	err = store.QueryRow(ctx, "SELECT count(*) FROM sync.backward_let").Scan(&backwardLetCount)
	require.NoError(t, err)
	assert.Equal(t, uint32(1), backwardLetCount, "Should have only 1 backward_let entry after cascade delete")

	// Verify the remaining backward_let is id=1
	var remainingBackwardLetID uint64
	err = store.QueryRow(ctx, "SELECT id FROM sync.backward_let").Scan(&remainingBackwardLetID)
	require.NoError(t, err)
	assert.Equal(t, uint64(1), remainingBackwardLetID, "Remaining backward_let should have id=1")

	// Now the 4 deposits with backward_let_id=2 are orphans
	// Run GetAndDeleteOrphanDepositBackups to find and delete them
	orphanDeposits, err := store.GetAndDeleteOrphanDepositBackups(ctx, nil)
	require.NoError(t, err)

	// Verify that 4 orphan deposits were returned
	assert.Equal(t, 4, len(orphanDeposits), "Should return 4 orphan deposits")

	// Verify the orphan deposits have the expected amounts (3000-6000 ETH)
	expectedAmounts := map[string]bool{
		"3000000000000000000": false,
		"4000000000000000000": false,
		"5000000000000000000": false,
		"6000000000000000000": false,
	}
	for i, dep := range orphanDeposits {
		amountStr := dep.Amount.String()
		if _, exists := expectedAmounts[amountStr]; exists {
			expectedAmounts[amountStr] = true
		}
		assert.Equal(t, uint32(1), dep.NetworkID, "NetworkID should be 1")
		assert.Equal(t, uint32(i)+2, dep.DepositCount, "BlockID should be different") // nolint:gosec
		assert.Equal(t, uint32(0), dep.DestinationNetwork, "DestinationNetwork should be 0")
	}

	// Verify all expected amounts were found
	for amount, found := range expectedAmounts {
		assert.True(t, found, "Should have found orphan deposit with amount %s", amount)
	}

	// Verify that orphan deposits were deleted: should have only 2 deposit_backup entries remaining
	var finalCount uint32
	err = store.QueryRow(ctx, "SELECT count(*) FROM sync.deposit_backup").Scan(&finalCount)
	require.NoError(t, err)
	assert.Equal(t, uint32(2), finalCount, "Should have only 2 valid deposit_backup entries remaining")

	// Verify the remaining deposits have valid backward_let_id=1
	var remainingBackwardLetIDs []int64
	rows, err := store.Query(ctx, "SELECT backward_let_id FROM sync.deposit_backup ORDER BY deposit_cnt")
	require.NoError(t, err)
	defer rows.Close()

	for rows.Next() {
		var backwardLetID int64
		err = rows.Scan(&backwardLetID)
		require.NoError(t, err)
		remainingBackwardLetIDs = append(remainingBackwardLetIDs, backwardLetID)
	}
	assert.Equal(t, []int64{1, 1}, remainingBackwardLetIDs, "All remaining deposits should reference backward_let_id=1")

	// Verify deposit_cnt of remaining deposits (should be 0 and 1)
	var depositCounts []uint32
	rows2, err := store.Query(ctx, "SELECT deposit_cnt FROM sync.deposit_backup ORDER BY deposit_cnt")
	require.NoError(t, err)
	defer rows2.Close()

	for rows2.Next() {
		var depositCnt uint32
		err = rows2.Scan(&depositCnt)
		require.NoError(t, err)
		depositCounts = append(depositCounts, depositCnt)
	}
	assert.Equal(t, []uint32{0, 1}, depositCounts, "Remaining deposits should have deposit_cnt 0 and 1")

	// Test that running GetAndDeleteOrphanDepositBackups again returns no orphans
	orphanDeposits2, err := store.GetAndDeleteOrphanDepositBackups(ctx, nil)
	require.NoError(t, err)
	assert.Equal(t, 0, len(orphanDeposits2), "Should return 0 orphan deposits on second run")

	// Verify count is still 2
	err = store.QueryRow(ctx, "SELECT count(*) FROM sync.deposit_backup").Scan(&finalCount)
	require.NoError(t, err)
	assert.Equal(t, uint32(2), finalCount, "Should still have 2 deposit_backup entries")
}