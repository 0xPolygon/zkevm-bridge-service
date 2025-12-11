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


	err = store.ResetDeposits(ctx, 1, 0, nil)
	require.Error(t, err)
	d, err := store.GetNumberDeposits(ctx, 1, 1, nil)
	require.NoError(t, err)
	assert.Equal(t, uint32(3), d)
	err = store.ResetDeposits(ctx, 1, 1, nil)
	require.NoError(t, err)
	d, err = store.GetNumberDeposits(ctx, 1, 1, nil)
	require.NoError(t, err)
	assert.Equal(t, uint32(1), d)
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
	err = store.AddBackwardLET(ctx, bLET, nil)
	require.NoError(t, err)

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