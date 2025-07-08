package migrations_test

import (
	"database/sql"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type migrationTest0014 struct{}

const conflictingDepositNet1ToNet3 = `
INSERT INTO sync.claim (block_id, network_id, index, mainnet_flag, rollup_index, orig_addr, dest_addr, tx_hash) 
VALUES(69, 3, 0, false, 0, decode('00','hex'), decode('00','hex'), decode('00','hex'));
` // Rollup 1 to Rollup 3

func (m migrationTest0014) InsertData(db *sql.DB) error {
	block := "INSERT INTO sync.block (id, block_num, block_hash, parent_hash, network_id, received_at) VALUES(69, 2803824, decode('27474F16174BBE50C294FE13C190B92E42B2368A6D4AEB8A4A015F52816296C3','hex'), decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C07','hex'), 3, '0001-01-01 01:00:00.000');"
	if _, err := db.Exec(block); err != nil {
		return err
	}
	const originalDepositSQL = `
		INSERT INTO sync.claim (block_id, network_id, index, mainnet_flag, rollup_index, orig_addr, dest_addr, tx_hash)
		VALUES(69, 3, 0, true, 0, decode('00','hex'), decode('00','hex'), decode('00','hex'));
	` // L1 to Rollup 3
	if _, err := db.Exec(originalDepositSQL); err != nil {
		return err
	}
	_, err := db.Exec(conflictingDepositNet1ToNet3)
	if err == nil || !strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint \"claim_pkey\" (SQLSTATE 23505)") {
		return errors.New("should violate primary key")
	}

	return nil
}

func (m migrationTest0014) RunAssertsAfterMigrationUp(t *testing.T, db *sql.DB) {
	// check that original row still in there
	selectClaim := `SELECT block_id, network_id, index, mainnet_flag, rollup_index FROM sync.claim;`
	row := db.QueryRow(selectClaim)
	var (
		block_id, network_id, index, rollup_index int
		mainnet_flag                              bool
	)
	assert.NoError(t, row.Scan(&block_id, &network_id, &index, &mainnet_flag, &rollup_index))
	assert.Equal(t, 69, block_id)
	assert.Equal(t, 3, network_id)
	assert.Equal(t, 0, index)
	assert.Equal(t, true, mainnet_flag)
	assert.Equal(t, 0, rollup_index)

	// Add deposit that originally would have caused pkey violation
	_, err := db.Exec(conflictingDepositNet1ToNet3)
	assert.NoError(t, err)

	// Remove conflicting deposit so it's possible to run the migration down
	_, err = db.Exec("DELETE FROM sync.claim WHERE mainnet_flag = false;")
	assert.NoError(t, err)
}

func (m migrationTest0014) RunAssertsAfterMigrationDown(t *testing.T, db *sql.DB) {
	// check that original row still in there
	selectClaim := `SELECT block_id, network_id, index, mainnet_flag, rollup_index FROM sync.claim;`
	row := db.QueryRow(selectClaim)
	var (
		block_id, network_id, index, rollup_index int
		mainnet_flag                              bool
	)
	assert.NoError(t, row.Scan(&block_id, &network_id, &index, &mainnet_flag, &rollup_index))
	assert.Equal(t, 69, block_id)
	assert.Equal(t, 3, network_id)
	assert.Equal(t, 0, index)
	assert.Equal(t, true, mainnet_flag)
	assert.Equal(t, 0, rollup_index)
}

func TestMigration0014(t *testing.T) {
	runMigrationTest(t, 14, migrationTest0014{})
}
