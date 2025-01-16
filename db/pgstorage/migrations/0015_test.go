package migrations_test

import (
	"database/sql"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

type migrationTest0015 struct{}

func (m migrationTest0015) InsertData(db *sql.DB) error {
	block := "INSERT INTO sync.block (id, block_num, block_hash, parent_hash, network_id, received_at) VALUES(69, 2803824, decode('27474F16174BBE50C294FE13C190B92E42B2368A6D4AEB8A4A015F52816296C3','hex'), decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C07','hex'), 0, '0001-01-01 01:00:00.000');"
	if _, err := db.Exec(block); err != nil {
		return err
	}
	block2 := "INSERT INTO sync.block (id, block_num, block_hash, parent_hash, network_id, received_at) VALUES(70, 2803824, decode('27474F16174BBE50C294FE13C190B92E42B2368A6D4AEB8A4A015F52816296C4','hex'), decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C08','hex'), 1, '0001-01-01 01:00:00.000');"
	if _, err := db.Exec(block2); err != nil {
		return err
	}
	const gerSQL = `
		INSERT INTO sync.exit_root
		(id, block_id, global_exit_root, exit_roots, network_id)
		VALUES(1, 69, decode('717E05DE47A87A7D1679E183F1C224150675F6302B7DA4EAAB526B2B91AE0761','hex'), '{decode(''5C7830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030'',''hex''),decode(''5C7832376165356261303864373239316339366338636264646363313438626634386136643638633739373462393433353666353337353465663631373164373537'',''hex'')}', 0);
		INSERT INTO sync.exit_root
		(id, block_id, global_exit_root, exit_roots, network_id)
		VALUES(2, 70, decode('717E05DE47A87A7D1679E183F1C224150675F6302B7DA4EAAB526B2B91AE0761','hex'), '{decode(''5C7830303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030'',''hex''),decode(''5C7832376165356261303864373239316339366338636264646363313438626634386136643638633739373462393433353666353337353465663631373164373537'',''hex'')}', 1);
	`
	if _, err := db.Exec(gerSQL); err != nil {
		return err
	}
	return nil
}

func (m migrationTest0015) RunAssertsAfterMigrationUp(t *testing.T, db *sql.DB) {
	selectGER := `SELECT allowed FROM sync.exit_root limit 1;`
	var allowed bool
	err := db.QueryRow(selectGER).Scan(&allowed)
	assert.NoError(t, err)

	const gerSQL = `
		INSERT INTO sync.exit_root
		(id, block_id, global_exit_root, exit_roots, network_id, allowed)
		VALUES(3, 69, decode('26C7509B1B6E04162FD30850A046A59A725F31FBF2CA43C0A7A015D667F3CFFD','hex'), '{decode(''5C7835626130303233323962353363313161326631646665393062313165303331373731383432303536636632313235623433646138313033633139396463643766'',''hex''),decode(''5C7832376165356261303864373239316339366338636264646363313438626634386136643638633739373462393433353666353337353465663631373164373537'',''hex'')}', 0, true);
		INSERT INTO sync.exit_root
		(id, block_id, global_exit_root, exit_roots, network_id, allowed)
		VALUES(4, 70, decode('26C7509B1B6E04162FD30850A046A59A725F31FBF2CA43C0A7A015D667F3CFFD','hex'), '{decode(''5C7835626130303233323962353363313161326631646665393062313165303331373731383432303536636632313235623433646138313033633139396463643766'',''hex''),decode(''5C7832376165356261303864373239316339366338636264646363313438626634386136643638633739373462393433353666353337353465663631373164373537'',''hex'')}', 1, false);`
	_, err = db.Exec(gerSQL)
	assert.NoError(t, err)

	selectGER = `SELECT allowed FROM sync.exit_root where id = 1;`
	err = db.QueryRow(selectGER).Scan(&allowed)
	assert.NoError(t, err)
	assert.Equal(t, true, allowed)

	selectGER = `SELECT allowed FROM sync.exit_root where id = 3;`
	err = db.QueryRow(selectGER).Scan(&allowed)
	assert.NoError(t, err)
	assert.Equal(t, true, allowed)

	selectGER = `SELECT allowed FROM sync.exit_root where id = 4;`
	err = db.QueryRow(selectGER).Scan(&allowed)
	assert.NoError(t, err)
	assert.Equal(t, false, allowed)
}

func (m migrationTest0015) RunAssertsAfterMigrationDown(t *testing.T, db *sql.DB) {
	var (
		id             uint64
		blockID        uint64
		globalExitRoot common.Hash
		exitRoots      [][]byte
		networkID      uint32
		allowed        bool
	)
	selectGER := `SELECT allowed FROM sync.exit_root limit 1;`
	err := db.QueryRow(selectGER).Scan(&allowed)
	assert.Error(t, err)

	selectGER = `SELECT id, block_id, global_exit_root, exit_roots, network_id FROM sync.exit_root where id = 4;`
	err = db.QueryRow(selectGER).Scan(&id, &blockID, &globalExitRoot, pq.Array(&exitRoots), &networkID)
	assert.NoError(t, err)
	assert.Equal(t, uint64(4), id)
	assert.Equal(t, uint64(70), blockID)
	assert.Equal(t, common.HexToHash("0x26c7509b1b6e04162fd30850a046a59a725f31fbf2ca43c0a7a015d667f3cffd"), globalExitRoot)
	assert.NotEqual(t, 0, len(exitRoots))
	assert.Equal(t, uint32(1), networkID)
}

func TestMigration0015(t *testing.T) {
	runMigrationTest(t, 15, migrationTest0015{})
}
