package migrations_test

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

type migrationTest0019 struct{}

func (m migrationTest0019) InsertData(db *sql.DB) error {
	return nil
}

func (m migrationTest0019) RunAssertsAfterMigrationUp(t *testing.T, db *sql.DB) {
	insertStatus := `INSERT INTO sync.status(network_id, percentage, remaining_blocks, synced)
		VALUES(0, 98, 4967295, false);`
	_, err := db.Exec(insertStatus)
	assert.NoError(t, err)
	insertStatus2 := `INSERT INTO sync.status(network_id, percentage, remaining_blocks, synced)
		VALUES(1, 2, 4294967295, false);`
	_, err = db.Exec(insertStatus2)
	assert.NoError(t, err)
	insertStatus3 := `INSERT INTO sync.status(network_id, percentage, remaining_blocks, synced)
		VALUES(2, 99, 32, true);`
	_, err = db.Exec(insertStatus3)
	assert.NoError(t, err)

	_, err = db.Exec(insertStatus3)
	assert.Error(t, err)

	var count uint32
	selectCount := "SELECT count(*) FROM sync.status"
	err = db.QueryRow(selectCount).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, uint32(3), count)
}

func (m migrationTest0019) RunAssertsAfterMigrationDown(t *testing.T, db *sql.DB) {
	var count uint32
	selectCount := "SELECT count(*) FROM sync.status"
	err := db.QueryRow(selectCount).Scan(&count)
	assert.Error(t, err)	
}

func TestMigration0019(t *testing.T) {
	runMigrationTest(t, 19, migrationTest0019{})
}
