package migrations_test

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

type migrationTest0018 struct{}

func (m migrationTest0018) InsertData(db *sql.DB) error {
	block := "INSERT INTO sync.block (id, block_num, block_hash, network_id) VALUES(1, 2803824, decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C07','hex'), 0);"
	if _, err := db.Exec(block); err != nil {
		return err
	}
	insertDeposit := `INSERT INTO sync.deposit(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim)
		VALUES(0, 42949672, 42949672, decode('0000000000000000000000000000000000000000','hex'), '10000000000000000000', 42949672, decode('C949254D682D8C9AD5682521675B8F43B102AEC4','hex'), 1, 42949672, decode('C2D6575EA98EB55E36B5AC6E11196800362594458A4B3143DB50E4995CB2422E','hex'), decode('','hex'), 92233720368547758, true);`
	if _, err := db.Exec(insertDeposit); err != nil {
		return err
	}
	return nil
}

func (m migrationTest0018) RunAssertsAfterMigrationUp(t *testing.T, db *sql.DB) {
	selectIgnore := "SELECT ignore FROM sync.deposit WHERE id = 92233720368547758;"
	var ignore bool
	err := db.QueryRow(selectIgnore).Scan(&ignore)
	assert.NoError(t, err)	
	assert.Equal(t, false, ignore)

	insertDeposit := `INSERT INTO sync.deposit(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim, ignore)
		VALUES(0, 4294967295, 4294967295, decode('0000000000000000000000000000000000000000','hex'), '10000000000000000000', 4294967295, decode('C949254D682D8C9AD5682521675B8F43B102AEC4','hex'), 1, 4294967295, decode('C2D6575EA98EB55E36B5AC6E11196800362594458A4B3143DB50E4995CB2422E','hex'), decode('','hex'), 9223372036854775807, true, false);`
	_, err = db.Exec(insertDeposit)
	assert.NoError(t, err)

	selectIgnore2 := "SELECT ignore FROM sync.deposit WHERE id = 9223372036854775807;"
	err = db.QueryRow(selectIgnore2).Scan(&ignore)
	assert.NoError(t, err)	
	assert.Equal(t, false, ignore)

	insertDeposit2 := `INSERT INTO sync.deposit(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim, ignore)
		VALUES(0, 429496729, 429496729, decode('0000000000000000000000000000000000000000','hex'), '10000000000000000000', 429496729, decode('C949254D682D8C9AD5682521675B8F43B102AEC4','hex'), 1, 429496729, decode('C2D6575EA98EB55E36B5AC6E11196800362594458A4B3143DB50E4995CB2422E','hex'), decode('','hex'), 922337203685477580, true, true);`
	_, err = db.Exec(insertDeposit2)
	assert.NoError(t, err)
	
	selectIgnore3 := "SELECT ignore FROM sync.deposit WHERE id = 922337203685477580;"
	err = db.QueryRow(selectIgnore3).Scan(&ignore)
	assert.NoError(t, err)	
	assert.Equal(t, true, ignore)
}

func (m migrationTest0018) RunAssertsAfterMigrationDown(t *testing.T, db *sql.DB) {
	insertDeposit := `INSERT INTO sync.deposit(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim, ignore)
		VALUES(0, 4294967, 4294967, decode('0000000000000000000000000000000000000000','hex'), '10000000000000000000', 4294967, decode('C949254D682D8C9AD5682521675B8F43B102AEC4','hex'), 1, 4294967, decode('C2D6575EA98EB55E36B5AC6E11196800362594458A4B3143DB50E4995CB2422E','hex'), decode('','hex'), 922337203685477, true, true);`
	_, err := db.Exec(insertDeposit)
	assert.Error(t, err)

	insertDeposit2 := `INSERT INTO sync.deposit(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim)
		VALUES(0, 4294967, 4294967, decode('0000000000000000000000000000000000000000','hex'), '10000000000000000000', 4294967, decode('C949254D682D8C9AD5682521675B8F43B102AEC4','hex'), 1, 4294967, decode('C2D6575EA98EB55E36B5AC6E11196800362594458A4B3143DB50E4995CB2422E','hex'), decode('','hex'), 922337203685477, true);`
	_, err = db.Exec(insertDeposit2)
	assert.NoError(t, err)

	selectCount := `SELECT count(*) FROM sync.deposit;`
	var count uint64
	err = db.QueryRow(selectCount).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, uint64(4), count)
}

func TestMigration0018(t *testing.T) {
	runMigrationTest(t, 18, migrationTest0018{})
}
