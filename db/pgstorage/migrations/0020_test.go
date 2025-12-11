package migrations_test

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

type migrationTest0020 struct{}

func (m migrationTest0020) InsertData(db *sql.DB) error {
	block := "INSERT INTO sync.block (id, block_num, block_hash, network_id) VALUES(1, 1, decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C07','hex'), 0);"
	if _, err := db.Exec(block); err != nil {
		return err
	}
	return nil
}

func (m migrationTest0020) RunAssertsAfterMigrationUp(t *testing.T, db *sql.DB) {
	insertBackwardLET := `INSERT INTO sync.backward_let(block_id, previous_deposit_cnt, previous_root, new_deposit_cnt, new_root)
		VALUES(1, 2, decode('A9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C08','hex'), 1, decode('B9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C09','hex'));`
	_, err := db.Exec(insertBackwardLET)
	assert.NoError(t, err)
	insertSetClaim := `INSERT INTO sync.set_unset_claim(block_id, mainnet_flag, rollup_index, index, global_index, type)
		VALUES(1, true, 0, 11, '18446744073709551627', 'SET');`
	_, err = db.Exec(insertSetClaim)
	assert.NoError(t, err)
	insertUnSetClaim := `INSERT INTO sync.set_unset_claim(block_id, mainnet_flag, rollup_index, index, global_index, type)
		VALUES(1, false, 4294967295, 4294967295, '18446744073709551615', 'UNSET');`
	_, err = db.Exec(insertUnSetClaim)
	assert.NoError(t, err)

	var count uint32
	selectCount := "SELECT count(*) FROM sync.backward_let"
	err = db.QueryRow(selectCount).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), count)

	selectCount1 := "SELECT count(*) FROM sync.set_unset_claim WHERE type='SET'"
	err = db.QueryRow(selectCount1).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), count)

	selectCount2 := "SELECT count(*) FROM sync.set_unset_claim WHERE type='UNSET'"
	err = db.QueryRow(selectCount2).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), count)

	selectCount3 := "SELECT count(*) FROM sync.set_unset_claim"
	err = db.QueryRow(selectCount3).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, uint32(2), count)
}

func (m migrationTest0020) RunAssertsAfterMigrationDown(t *testing.T, db *sql.DB) {
	var count uint32
	selectCount := "SELECT count(*) FROM sync.backward_let"
	err := db.QueryRow(selectCount).Scan(&count)
	assert.Error(t, err)
	
	selectCount1 := "SELECT count(*) FROM sync.set_unset_claim"
	err = db.QueryRow(selectCount1).Scan(&count)
	assert.Error(t, err)
}

func TestMigration0020(t *testing.T) {
	runMigrationTest(t, 20, migrationTest0020{})
}
