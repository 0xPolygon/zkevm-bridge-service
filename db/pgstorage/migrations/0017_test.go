package migrations_test

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

type migrationTest0017 struct{}

func (m migrationTest0017) InsertData(db *sql.DB) error {
	block := "INSERT INTO sync.block (id, block_num, block_hash, network_id) VALUES(1, 2803824, decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C07','hex'), 0);"
	if _, err := db.Exec(block); err != nil {
		return err
	}
	return nil
}

func (m migrationTest0017) RunAssertsAfterMigrationUp(t *testing.T, db *sql.DB) {
	insertClaim1 := `INSERT INTO sync.claim
	(network_id, "index", orig_net, orig_addr, amount, dest_addr, block_id, tx_hash, rollup_index, mainnet_flag, global_index)
	VALUES(1, 11, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, decode('BF2C816AB6F8A8F5F9DDA6EE97D433CC841E69B5669A5CDF499826FA4B99C179','hex'), 1, false, 4294967307);`
	_, err := db.Exec(insertClaim1)
	assert.NoError(t, err)
	insertClaim2 := `INSERT INTO sync.claim
	(network_id, "index", orig_net, orig_addr, amount, dest_addr, block_id, tx_hash, rollup_index, mainnet_flag)
	VALUES(1, 12, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, decode('BF2C816AB6F8A8F5F9DDA6EE97D433CC841E69B5669A5CDF499826FA4B99C179','hex'), 1, false);`
	_, err = db.Exec(insertClaim2)
	assert.NoError(t, err)
	selectClaim1 := "SELECT global_index FROM sync.claim WHERE network_id = 1 AND index = 11;"
	var globalIndex string
	err = db.QueryRow(selectClaim1).Scan(&globalIndex)
	assert.NoError(t, err)
	assert.Equal(t, "4294967307", globalIndex)
	selectClaim2 := "SELECT global_index FROM sync.claim WHERE network_id = 1 AND index = 12;"
	var globalIndex2 string
	err = db.QueryRow(selectClaim2).Scan(&globalIndex2)
	assert.NoError(t, err)	
	assert.Equal(t, "", globalIndex2)
}

func (m migrationTest0017) RunAssertsAfterMigrationDown(t *testing.T, db *sql.DB) {
	insertClaim1 := `INSERT INTO sync.claim
	(network_id, "index", orig_net, orig_addr, amount, dest_addr, block_id, tx_hash, rollup_index, mainnet_flag, global_index)
	VALUES(1, 13, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, decode('BF2C816AB6F8A8F5F9DDA6EE97D433CC841E69B5669A5CDF499826FA4B99C179','hex'), 1, false, 4294967307);`
	_, err := db.Exec(insertClaim1)
	assert.Error(t, err)
	insertClaim2 := `INSERT INTO sync.claim
	(network_id, "index", orig_net, orig_addr, amount, dest_addr, block_id, tx_hash, rollup_index, mainnet_flag)
	VALUES(1, 14, 0, decode('0000000000000000000000000000000000000000','hex'), '90000000000000000', decode('F39FD6E51AAD88F6F4CE6AB8827279CFFFB92266','hex'), 1, decode('BF2C816AB6F8A8F5F9DDA6EE97D433CC841E69B5669A5CDF499826FA4B99C179','hex'), 1, false);`
	_, err = db.Exec(insertClaim2)
	assert.NoError(t, err)
	selectCount := `SELECT count(*) FROM sync.claim;`
	var count uint64
	err = db.QueryRow(selectCount).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, uint64(3), count)
}

func TestMigration0017(t *testing.T) {
	runMigrationTest(t, 17, migrationTest0017{})
}
