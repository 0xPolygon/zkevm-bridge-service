package migrations_test

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

type migrationTest0023 struct{}

func (m migrationTest0023) InsertData(db *sql.DB) error {
	// Insert basic data required for testing indexes
	block := "INSERT INTO sync.block (id, block_num, block_hash, network_id) VALUES(1, 1, decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C07','hex'), 0);"
	if _, err := db.Exec(block); err != nil {
		return err
	}

	deposit := `INSERT INTO sync.deposit(id, leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, ready_for_claim, ignore)
		VALUES(1, 0, 0, 0, decode('A9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C08','hex'), '1000000000000000000', 1, decode('B9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C09','hex'), 1, 1, decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C10','hex'), decode('D9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C11','hex'), true, false);`
	if _, err := db.Exec(deposit); err != nil {
		return err
	}

	return nil
}

func (m migrationTest0023) RunAssertsAfterMigrationUp(t *testing.T, db *sql.DB) {
	// List of all indexes that should exist after migration 0023
	expectedIndexes := []struct {
		schema    string
		tableName string
		indexName string
	}{
		// sync.deposit indexes (8 indexes)
		{"sync", "deposit", "idx_deposit_network_id_block_id"},
		{"sync", "deposit", "idx_deposit_network_id_deposit_cnt"},
		{"sync", "deposit", "idx_deposit_dest_addr"},
		{"sync", "deposit", "idx_deposit_dest_net_ready_ignore"},
		{"sync", "deposit", "idx_deposit_metadata_lookup"},
		{"sync", "deposit", "idx_deposit_l2_claim_status"},
		{"sync", "deposit", "idx_deposit_id_network_dest"},
		{"sync", "deposit", "idx_deposit_orig_addr"},

		// sync.claim indexes
		{"sync", "claim", "idx_claim_mainnet_network"},
		{"sync", "claim", "idx_claim_rollup_network"},
		{"sync", "claim", "idx_claim_dest_addr"},
		{"sync", "claim", "idx_claim_network_id_index"},
		{"sync", "claim", "idx_claim_global_index_network"},

		// sync.block indexes
		{"sync", "block", "idx_block_network_id_block_num_desc"},
		{"sync", "block", "idx_block_num_network_id"},

		// sync.exit_root indexes
		{"sync", "exit_root", "idx_exit_root_allowed_block_id_network"},
		{"sync", "exit_root", "idx_exit_root_network_allowed_id_desc"},
		{"sync", "exit_root", "idx_exit_root_ger_lookup"},
		{"sync", "exit_root", "idx_exit_root_l2_ger"},
		{"sync", "exit_root", "idx_exit_root_exit_roots_array"},
		{"sync", "exit_root", "idx_exit_root_network_id_desc"},
		{"sync", "exit_root", "idx_exit_root_ger_network"},

		// sync.token_wrapped indexes
		{"sync", "token_wrapped", "idx_token_wrapped_orig_net_addr"},

		// mt.root indexes
		{"mt", "root", "idx_root_deposit_id_network"},
		{"mt", "root", "idx_root_network"},

		// mt.rollup_exit indexes
		{"mt", "rollup_exit", "idx_rollup_exit_root_rollup_id"},
		{"mt", "rollup_exit", "idx_rollup_exit_root"},
		{"mt", "rollup_exit", "idx_rollup_exit_rollup_id_id"},

		// sync.monitored_txs indexes
		{"sync", "monitored_txs", "idx_monitored_txs_deposit_id"},
		{"sync", "monitored_txs", "idx_monitored_txs_status_created"},
		{"sync", "monitored_txs", "idx_monitored_txs_group_id"},

		// sync.remove_exit_root indexes
		{"sync", "remove_exit_root", "idx_remove_exit_root_ger_network"},
		{"sync", "remove_exit_root", "idx_remove_exit_root_block_id"},

		// sync.backward_let indexes
		{"sync", "backward_let", "idx_backward_let_new_root"},
		{"sync", "backward_let", "idx_backward_let_previous_root"},
		{"sync", "backward_let", "idx_backward_let_block_id"},

		// sync.forward_let indexes
		{"sync", "forward_let", "idx_forward_let_new_root"},
		{"sync", "forward_let", "idx_forward_let_previous_root"},
		{"sync", "forward_let", "idx_forward_let_block_id"},

		// sync.set_unset_claim indexes
		{"sync", "set_unset_claim", "idx_set_unset_claim_index_rollup"},
		{"sync", "set_unset_claim", "idx_set_unset_claim_global_index"},
		{"sync", "set_unset_claim", "idx_set_unset_claim_type"},
		{"sync", "set_unset_claim", "idx_set_unset_claim_block_id"},

		// sync.deposit_backup indexes
		{"sync", "deposit_backup", "idx_deposit_backup_deposit_cnt"},
		{"sync", "deposit_backup", "idx_deposit_backup_network_id"},
	}

	// Verify all expected indexes exist
	for _, idx := range expectedIndexes {
		var indexExists bool
		checkIndex := `SELECT EXISTS (
			SELECT 1 FROM pg_indexes
			WHERE schemaname = $1
			AND tablename = $2
			AND indexname = $3
		)`
		err := db.QueryRow(checkIndex, idx.schema, idx.tableName, idx.indexName).Scan(&indexExists)
		assert.NoError(t, err, "Error checking index %s.%s.%s", idx.schema, idx.tableName, idx.indexName)
		assert.True(t, indexExists, "Index %s.%s.%s should exist after migration up", idx.schema, idx.tableName, idx.indexName)
	}

	// Verify specific index properties
	// Test GIN index on sync.exit_root
	var isGinIndex bool
	checkGinIndex := `SELECT EXISTS (
		SELECT 1 FROM pg_indexes
		WHERE schemaname = 'sync'
		AND tablename = 'exit_root'
		AND indexname = 'idx_exit_root_exit_roots_array'
		AND indexdef LIKE '%USING gin%'
	)`
	err := db.QueryRow(checkGinIndex).Scan(&isGinIndex)
	assert.NoError(t, err)
	assert.True(t, isGinIndex, "Index idx_exit_root_exit_roots_array should be a GIN index")

	// Verify indexes work with sample queries
	// Test deposit index
	var depositCount int
	err = db.QueryRow("SELECT COUNT(*) FROM sync.deposit WHERE network_id = 0 AND deposit_cnt = 1").Scan(&depositCount)
	assert.NoError(t, err)
	assert.Equal(t, 1, depositCount, "Should find 1 deposit with network_id=0 and deposit_cnt=1")

	// Test that queries execute without errors (indexes are being used)
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM sync.deposit WHERE ready_for_claim = true AND dest_net = 1)").Scan(&exists)
	assert.NoError(t, err)
}

func (m migrationTest0023) RunAssertsAfterMigrationDown(t *testing.T, db *sql.DB) {
	// List of ALL indexes that should NOT exist after migration down
	// This must match exactly the indexes created in the UP migration
	indexesToCheck := []struct {
		schema    string
		tableName string
		indexName string
	}{
		// sync.deposit indexes (8 indexes)
		{"sync", "deposit", "idx_deposit_network_id_block_id"},
		{"sync", "deposit", "idx_deposit_network_id_deposit_cnt"},
		{"sync", "deposit", "idx_deposit_dest_addr"},
		{"sync", "deposit", "idx_deposit_dest_net_ready_ignore"},
		{"sync", "deposit", "idx_deposit_metadata_lookup"},
		{"sync", "deposit", "idx_deposit_l2_claim_status"},
		{"sync", "deposit", "idx_deposit_id_network_dest"},
		{"sync", "deposit", "idx_deposit_orig_addr"},

		// sync.claim indexes (5 indexes)
		{"sync", "claim", "idx_claim_mainnet_network"},
		{"sync", "claim", "idx_claim_rollup_network"},
		{"sync", "claim", "idx_claim_dest_addr"},
		{"sync", "claim", "idx_claim_network_id_index"},
		{"sync", "claim", "idx_claim_global_index_network"},

		// sync.block indexes (2 indexes)
		{"sync", "block", "idx_block_network_id_block_num_desc"},
		{"sync", "block", "idx_block_num_network_id"},

		// sync.exit_root indexes (7 indexes)
		{"sync", "exit_root", "idx_exit_root_allowed_block_id_network"},
		{"sync", "exit_root", "idx_exit_root_network_allowed_id_desc"},
		{"sync", "exit_root", "idx_exit_root_ger_lookup"},
		{"sync", "exit_root", "idx_exit_root_l2_ger"},
		{"sync", "exit_root", "idx_exit_root_exit_roots_array"},
		{"sync", "exit_root", "idx_exit_root_network_id_desc"},
		{"sync", "exit_root", "idx_exit_root_ger_network"},

		// sync.token_wrapped indexes (1 index)
		{"sync", "token_wrapped", "idx_token_wrapped_orig_net_addr"},

		// mt.root indexes (2 indexes)
		{"mt", "root", "idx_root_deposit_id_network"},
		{"mt", "root", "idx_root_network"},

		// mt.rollup_exit indexes (3 indexes)
		{"mt", "rollup_exit", "idx_rollup_exit_root_rollup_id"},
		{"mt", "rollup_exit", "idx_rollup_exit_root"},
		{"mt", "rollup_exit", "idx_rollup_exit_rollup_id_id"},

		// sync.monitored_txs indexes (3 indexes)
		{"sync", "monitored_txs", "idx_monitored_txs_deposit_id"},
		{"sync", "monitored_txs", "idx_monitored_txs_status_created"},
		{"sync", "monitored_txs", "idx_monitored_txs_group_id"},

		// sync.remove_exit_root indexes (2 indexes)
		{"sync", "remove_exit_root", "idx_remove_exit_root_ger_network"},
		{"sync", "remove_exit_root", "idx_remove_exit_root_block_id"},

		// sync.backward_let indexes (3 indexes)
		{"sync", "backward_let", "idx_backward_let_new_root"},
		{"sync", "backward_let", "idx_backward_let_previous_root"},
		{"sync", "backward_let", "idx_backward_let_block_id"},

		// sync.forward_let indexes (3 indexes)
		{"sync", "forward_let", "idx_forward_let_new_root"},
		{"sync", "forward_let", "idx_forward_let_previous_root"},
		{"sync", "forward_let", "idx_forward_let_block_id"},

		// sync.set_unset_claim indexes (4 indexes)
		{"sync", "set_unset_claim", "idx_set_unset_claim_index_rollup"},
		{"sync", "set_unset_claim", "idx_set_unset_claim_global_index"},
		{"sync", "set_unset_claim", "idx_set_unset_claim_type"},
		{"sync", "set_unset_claim", "idx_set_unset_claim_block_id"},

		// sync.deposit_backup indexes (2 indexes)
		{"sync", "deposit_backup", "idx_deposit_backup_deposit_cnt"},
		{"sync", "deposit_backup", "idx_deposit_backup_network_id"},
	}
	// Total: 45 indexes

	for _, idx := range indexesToCheck {
		var indexExists bool
		checkIndex := `SELECT EXISTS (
			SELECT 1 FROM pg_indexes
			WHERE schemaname = $1
			AND tablename = $2
			AND indexname = $3
		)`
		err := db.QueryRow(checkIndex, idx.schema, idx.tableName, idx.indexName).Scan(&indexExists)
		assert.NoError(t, err, "Error checking index %s.%s.%s", idx.schema, idx.tableName, idx.indexName)
		assert.False(t, indexExists, "Index %s.%s.%s should NOT exist after migration down", idx.schema, idx.tableName, idx.indexName)
	}

	// Verify data still exists (only indexes were removed, not data)
	var depositCount int
	err := db.QueryRow("SELECT COUNT(*) FROM sync.deposit").Scan(&depositCount)
	assert.NoError(t, err)
	assert.Equal(t, 1, depositCount, "Deposit data should still exist after migration down")
}

func TestMigration0023(t *testing.T) {
	runMigrationTest(t, 23, migrationTest0023{})
}
