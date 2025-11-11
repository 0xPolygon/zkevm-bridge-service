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

	var count uint32
	selectCount := "SELECT count(*) FROM sync.backward_let"
	err = db.QueryRow(selectCount).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), count)
}

func (m migrationTest0020) RunAssertsAfterMigrationDown(t *testing.T, db *sql.DB) {
	var count uint32
	selectCount := "SELECT count(*) FROM sync.backward_let"
	err := db.QueryRow(selectCount).Scan(&count)
	assert.Error(t, err)	
}

func TestMigration0020(t *testing.T) {
	runMigrationTest(t, 20, migrationTest0020{})
}
