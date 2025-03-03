package migrations_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

type migrationTest0016 struct{}

func (m migrationTest0016) InsertData(db *sql.DB) error {
	block := "INSERT INTO sync.block (id, block_num, block_hash, parent_hash, network_id, received_at) VALUES(69, 2803824, decode('27474F16174BBE50C294FE13C190B92E42B2368A6D4AEB8A4A016F52816296C3','hex'), decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C07','hex'), 0, '0001-01-01 01:00:00.000');"
	if _, err := db.Exec(block); err != nil {
		return err
	}
	block2 := "INSERT INTO sync.block (id, block_num, block_hash, parent_hash, network_id, received_at) VALUES(70, 2803824, decode('27474F16174BBE50C294FE13C190B92E42B2368A6D4AEB8A4A016F52816296C4','hex'), decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C08','hex'), 1, '0001-01-01 01:00:00.000');"
	if _, err := db.Exec(block2); err != nil {
		return err
	}
	return nil
}

func (m migrationTest0016) RunAssertsAfterMigrationUp(t *testing.T, db *sql.DB) {
	selectHashParent := `SELECT parent_hash FROM sync.block limit 1;`
	var hashParent common.Hash
	err := db.QueryRow(selectHashParent).Scan(&hashParent)
	assert.Error(t, err)

	selectReceivedAt := `SELECT received_at FROM sync.block limit 1;`
	var receivedAt time.Time
	err = db.QueryRow(selectReceivedAt).Scan(&receivedAt)
	assert.Error(t, err)

	selectCount := `SELECT count(*) FROM sync.block;`
	var count uint64
	err = db.QueryRow(selectCount).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, uint64(3), count)
}

func (m migrationTest0016) RunAssertsAfterMigrationDown(t *testing.T, db *sql.DB) {
	// Read values from the table
	selectHashParent := `SELECT parent_hash FROM sync.block where block_num='2803824' AND network_id = 0;`
	var hashParent common.Hash
	err := db.QueryRow(selectHashParent).Scan(&hashParent)
	assert.NoError(t, err)
	assert.Equal(t, common.Hash{}, hashParent)

	selectReceivedAt := `SELECT received_at FROM sync.block where block_num='2803824' AND network_id = 0;`
	var receivedAt time.Time
	err = db.QueryRow(selectReceivedAt).Scan(&receivedAt)
	assert.NoError(t, err)

	selectCount := `SELECT count(*) FROM sync.block;`
	var count uint64
	err = db.QueryRow(selectCount).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, uint64(3), count)
}

func TestMigration0016(t *testing.T) {
	runMigrationTest(t, 16, migrationTest0016{})
}
