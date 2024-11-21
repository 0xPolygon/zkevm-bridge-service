package migrations_test

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

type migrationTest0013 struct{}

func (m migrationTest0013) InsertData(db *sql.DB) error {
	block := "INSERT INTO sync.block (id, block_num, block_hash, parent_hash, network_id, received_at) VALUES(2, 2803824, decode('27474F16174BBE50C294FE13C190B92E42B2368A6D4AEB8A4A015F52816296C3','hex'), decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C07','hex'), 1, '0001-01-01 01:00:00.000');"
	if _, err := db.Exec(block); err != nil {
		return err
	}
	return nil
}

func (m migrationTest0013) RunAssertsAfterMigrationUp(t *testing.T, db *sql.DB) {
	insertDeposit := "INSERT INTO sync.deposit(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim) " +
		"VALUES(0, 4294967295, 4294967295, decode('0000000000000000000000000000000000000000','hex'), '10000000000000000000', 4294967295, decode('C949254D682D8C9AD5682521675B8F43B102AEC4','hex'), 2, 4294967295, decode('C2D6575EA98EB55E36B5AC6E11196800362594458A4B3143DB50E4995CB2422E','hex'), decode('','hex'), 9223372036854775807, true);"
	if _, err := db.Exec(insertDeposit); err != nil {
		assert.NoError(t, err)
	}
	insertClaim := "INSERT INTO sync.Claim (network_id, index, orig_net, orig_addr, amount, dest_addr, block_id, tx_hash) VALUES(4294967295, 4294967295, 4294967295, decode('0000000000000000000000000000000000000000','hex'), '300000000000000000', decode('14567C0DCF79C20FE1A21E36EC975D1775A1905C','hex'), 2, decode('A9505DB7D7EDD08947F12F2B1F7898148FFB43D80BCB977B78161EF14173D575','hex'));"
	if _, err := db.Exec(insertClaim); err != nil {
		assert.NoError(t, err)
	}
	insertTokenWrapped := "INSERT INTO sync.token_wrapped (network_id, orig_net, orig_token_addr, wrapped_token_addr,  block_id) " +
		"VALUES(4294967295, 4294967295, decode('0000000000000000000000000000000000000000','hex'),decode('0000000000000000000000000000000000000000','hex'), 2);"
	if _, err := db.Exec(insertTokenWrapped); err != nil {
		assert.NoError(t, err)
	}
	block := "INSERT INTO sync.block (id, block_num, block_hash, parent_hash, network_id, received_at) VALUES(9223372036854775807, 4294967295, decode('C2D6575EA98EB55E36B5AC6E11196800362594458A4B3143DB50E4995CB2422E','hex'), decode('C2D6575EA98EB55E36B5AC6E11196800362594458A4B3143DB50E4995CB2422E','hex'), 4294967295, '0001-01-01 01:00:00.000');"
	if _, err := db.Exec(block); err != nil {
		assert.NoError(t, err)
	}
	exitRoot := "INSERT INTO sync.exit_root(id, block_id, global_exit_root, exit_roots, network_id) VALUES(9223372036854775807, 9223372036854775807, decode('B881611B39DC5EAFF3AF06DCA56A0AB9997EF9F72FA2B34BCC80F1CCDC4242CD','hex'), '{decode(''5C7865386436396433336461383039616339653861323963666632373264663630656461373461646330626164653230313733393639636464656333643531616232'',''hex''),decode(''5C7832376165356261303864373239316339366338636264646363313438626634386136643638633739373462393433353666353337353465663631373164373537'',''hex'')}', 4294967295);"
	if _, err := db.Exec(exitRoot); err != nil {
		assert.NoError(t, err)
	}
	root := "INSERT INTO mt.root(root, network, deposit_id) VALUES(decode('E8D69D33DA809AC9E8A29CFF272DF60EDA74ADC0BADE20173969CDDEC3D51AB2','hex'), 4294967295, 9223372036854775807);"
	if _, err := db.Exec(root); err != nil {
		assert.NoError(t, err)
	}
	// Remove values for down migration
	_, err := db.Exec("DELETE FROM sync.claim;")
	assert.NoError(t, err)
	_, err = db.Exec("DELETE FROM mt.root;")
	assert.NoError(t, err)
	_, err = db.Exec("DELETE FROM sync.deposit;")
	assert.NoError(t, err)
	_, err = db.Exec("DELETE FROM sync.token_wrapped;")
	assert.NoError(t, err)
	_, err = db.Exec("DELETE FROM sync.block WHERE id = 9223372036854775807;")
	assert.NoError(t, err)
	_, err = db.Exec("DELETE FROM sync.exit_root;")
	assert.NoError(t, err)
}

func (m migrationTest0013) RunAssertsAfterMigrationDown(t *testing.T, db *sql.DB) {
	insertDeposit := "INSERT INTO sync.deposit(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim) " +
		"VALUES(0, 4294967295, 4294967295, decode('0000000000000000000000000000000000000000','hex'), '10000000000000000000', 4294967295, decode('C949254D682D8C9AD5682521675B8F43B102AEC4','hex'), 2, 4294967295, decode('C2D6575EA98EB55E36B5AC6E11196800362594458A4B3143DB50E4995CB2422E','hex'), decode('','hex'), 1, true);"
	if _, err := db.Exec(insertDeposit); err != nil {
		assert.Error(t, err)
	}
	insertClaim := "INSERT INTO sync.Claim (network_id, index, orig_net, orig_addr, amount, dest_addr, block_id, tx_hash) VALUES(4294967295, 4294967295, 4294967295, decode('0000000000000000000000000000000000000000','hex'), '300000000000000000', decode('14567C0DCF79C20FE1A21E36EC975D1775A1905C','hex'), 2, decode('A9505DB7D7EDD08947F12F2B1F7898148FFB43D80BCB977B78161EF14173D575','hex'));"
	if _, err := db.Exec(insertClaim); err != nil {
		assert.Error(t, err)
	}
	insertTokenWrapped := "INSERT INTO sync.token_wrapped (network_id, orig_net, orig_token_addr, wrapped_token_addr,  block_id) " +
		"VALUES(4294967295, 4294967295, decode('0000000000000000000000000000000000000000','hex'),decode('0000000000000000000000000000000000000000','hex'), 2);"
	if _, err := db.Exec(insertTokenWrapped); err != nil {
		assert.Error(t, err)
	}
	block := "INSERT INTO sync.block (id, block_num, block_hash, parent_hash, network_id, received_at) VALUES(5, 4294967295, decode('C2D6575EA98EB55E36B5AC6E11196800362594458A4B3143DB50E4995CB2422E','hex'), decode('C2D6575EA98EB55E36B5AC6E11196800362594458A4B3143DB50E4995CB2422E','hex'), 4294967295, '0001-01-01 01:00:00.000');"
	if _, err := db.Exec(block); err != nil {
		assert.Error(t, err)
	}
	exitRoot := "INSERT INTO sync.exit_root(id, block_id, global_exit_root, exit_roots, network_id) VALUES(2, 2, decode('B881611B39DC5EAFF3AF06DCA56A0AB9997EF9F72FA2B34BCC80F1CCDC4242CD','hex'), '{decode(''5C7865386436396433336461383039616339653861323963666632373264663630656461373461646330626164653230313733393639636464656333643531616232'',''hex''),decode(''5C7832376165356261303864373239316339366338636264646363313438626634386136643638633739373462393433353666353337353465663631373164373537'',''hex'')}', 4294967295);"
	if _, err := db.Exec(exitRoot); err != nil {
		assert.Error(t, err)
	}
	root := "INSERT INTO mt.root(root, network, deposit_id) VALUES(decode('E8D69D33DA809AC9E8A29CFF272DF60EDA74ADC0BADE20173969CDDEC3D51AB2','hex'), 4294967295, 1);"
	if _, err := db.Exec(root); err != nil {
		assert.Error(t, err)
	}
}

func TestMigration0013(t *testing.T) {
	runMigrationTest(t, 13, migrationTest0013{})
}
