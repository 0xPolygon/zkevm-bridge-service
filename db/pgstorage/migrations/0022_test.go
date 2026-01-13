package migrations_test

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

type migrationTest0022 struct{}

func (m migrationTest0022) InsertData(db *sql.DB) error {
	block := "INSERT INTO sync.block (id, block_num, block_hash, network_id) VALUES(1, 1, decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C07','hex'), 0);"
	if _, err := db.Exec(block); err != nil {
		return err
	}
	return nil
}

func (m migrationTest0022) RunAssertsAfterMigrationUp(t *testing.T, db *sql.DB) {
	insertDepositBackup := `INSERT INTO sync.deposit_backup(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, ready_for_claim, backward_let_id)
		VALUES(0, 1, 0, decode('A9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C08','hex'), '1000000000000000000', 1, decode('B9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C09','hex'), 1, 5, decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C10','hex'), decode('D9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C11','hex'), true, 1);`
	_, err := db.Exec(insertDepositBackup)
	assert.NoError(t, err)

	var count uint32
	selectCount := "SELECT count(*) FROM sync.deposit_backup"
	err = db.QueryRow(selectCount).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), count)

	// Verify the data was inserted correctly
	var leafType, networkID, origNet, destNet, depositCnt, backwardLetID int64
	var amount string
	var readyForClaim bool
	selectData := "SELECT leaf_type, network_id, orig_net, amount, dest_net, deposit_cnt, ready_for_claim, backward_let_id FROM sync.deposit_backup WHERE id = 1"
	err = db.QueryRow(selectData).Scan(&leafType, &networkID, &origNet, &amount, &destNet, &depositCnt, &readyForClaim, &backwardLetID)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), leafType)
	assert.Equal(t, int64(1), networkID)
	assert.Equal(t, int64(0), origNet)
	assert.Equal(t, "1000000000000000000", amount)
	assert.Equal(t, int64(1), destNet)
	assert.Equal(t, int64(5), depositCnt)
	assert.Equal(t, true, readyForClaim)
	assert.Equal(t, int64(1), backwardLetID)
}

func (m migrationTest0022) RunAssertsAfterMigrationDown(t *testing.T, db *sql.DB) {
	var count uint32
	selectCount := "SELECT count(*) FROM sync.deposit_backup"
	err := db.QueryRow(selectCount).Scan(&count)
	assert.Error(t, err)
}

func TestMigration0022(t *testing.T) {
	runMigrationTest(t, 22, migrationTest0022{})
}
