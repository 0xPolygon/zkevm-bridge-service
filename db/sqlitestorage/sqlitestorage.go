package sqlitestorage

import (
	"context"
	"errors"
	"fmt"
	"math"
	"database/sql"
	"math/big"
	"strings"
	"time"

	ctmtypes "github.com/0xPolygonHermez/zkevm-bridge-service/claimtxman/types"
	"github.com/0xPolygonHermez/zkevm-bridge-service/etherman"
	"github.com/0xPolygonHermez/zkevm-bridge-service/log"
	"github.com/0xPolygonHermez/zkevm-bridge-service/utils/gerror"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/mattn/go-sqlite3"
)

// DBStorage implements the Storage interface
type DBStorage struct {
	DB *sql.DB
}

// getExecQuerier determines which execQuerier to use, dbTx or the main pgxpool
func (st *DBStorage) getExecQuerier(dbTx interface{}) execQuerier {
	if dbTx != nil {
		DBTx := dbTx.(*sql.Tx)
		return DBTx
	}
	return st.DB
}

// NewSQLiteStorage creates a new Storage DB
func NewSQLiteStorage(cfg Config) (*DBStorage, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?_txlock=exclusive&_foreign_keys=on&_journal_mode=WAL", cfg.DBFile))
	if err != nil {
		return nil, err
	}
	return &DBStorage{DB: db}, nil
}

// Rollback rollbacks a db transaction.
func (st *DBStorage) Rollback(_ context.Context, dbTx interface{}) error {
	if dbTx != nil {
		DBTx := dbTx.(*sql.Tx)
		return DBTx.Rollback()
	}

	return gerror.ErrNilDBTransaction
}

// Commit commits a db transaction.
func (st *DBStorage) Commit(_ context.Context, dbTx interface{}) error {
	if dbTx != nil {
		DBTx := dbTx.(*sql.Tx)
		return DBTx.Commit()
	}
	return gerror.ErrNilDBTransaction
}

// BeginDBTransaction starts a transaction block.
func (st *DBStorage) BeginDBTransaction(_ context.Context) (interface{}, error) {
	return st.DB.Begin()
}

// GetLastBlock gets the last block.
func (st *DBStorage) GetLastBlock(_ context.Context, networkID uint32, dbTx interface{}) (*etherman.Block, error) {
	var block etherman.Block
	const getLastBlockSQL = "SELECT id, block_num, block_hash, network_id FROM block where network_id = ? ORDER BY block_num DESC LIMIT 1"

	e := st.getExecQuerier(dbTx)
	err := e.QueryRow(getLastBlockSQL, networkID).Scan(&block.ID, &block.BlockNumber, &block.BlockHash, &block.NetworkID)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, gerror.ErrStorageNotFound
	}

	return &block, err
}

// AddBlock adds a new block to the storage.
func (st *DBStorage) AddBlock(_ context.Context, block *etherman.Block, dbTx interface{}) (uint64, error) {
	e := st.getExecQuerier(dbTx)
	insertQuery := `INSERT INTO block (block_num, block_hash, network_id)
		VALUES (?, ?, ?)
		ON CONFLICT (block_hash) DO NOTHING;`
	_, err := e.Exec(insertQuery, block.BlockNumber, block.BlockHash, block.NetworkID)
	if err != nil {
		return 0, fmt.Errorf("error inserting block. Error: %s", err.Error())
	}

	// Retrieve the block ID
	selectQuery := "SELECT id FROM block WHERE block_hash = ?;"
	var blockID uint64
	err = e.QueryRow(selectQuery, block.BlockHash).Scan(&blockID)
	if err != nil {
		return 0, fmt.Errorf("error getting blockID after inserting the block. Error: %s", err.Error())
	}

	return blockID, err
}

// AddGlobalExitRoot adds a new ExitRoot to the db.
func (st *DBStorage) AddGlobalExitRoot(_ context.Context, exitRoot *etherman.GlobalExitRoot, dbTx interface{}) error {
	const addExitRootSQL = "INSERT INTO exit_root (block_id, global_exit_root, mainnet_exit_root, rollups_exit_root, network_id, allowed) VALUES (?, ?, ?, ?, ?, true)"
	exitRoots := [][]byte{}
	if len(exitRoot.ExitRoots) != 0 {
		exitRoots = [][]byte{exitRoot.ExitRoots[0][:], exitRoot.ExitRoots[1][:]}
	}
	e := st.getExecQuerier(dbTx)
	_, err := e.Exec(addExitRootSQL, exitRoot.BlockID, exitRoot.GlobalExitRoot, exitRoots[0], exitRoots[1], exitRoot.NetworkID)
	return err
}

// AddDeposit adds new deposit to the storage.
func (st *DBStorage) AddDeposit(_ context.Context, deposit *etherman.Deposit, dbTx interface{}) (uint64, error) {
	const addDepositSQL = "INSERT INTO deposit (leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	e := st.getExecQuerier(dbTx)
	result, err := e.Exec(addDepositSQL, deposit.LeafType, deposit.NetworkID, deposit.OriginalNetwork, deposit.OriginalAddress, deposit.Amount.String(), deposit.DestinationNetwork, deposit.DestinationAddress, deposit.BlockID, deposit.DepositCount, deposit.TxHash, deposit.Metadata)
	if err != nil {
		return 0, err
	}
	depositID, err := result.LastInsertId()
	return uint64(depositID), err
}

// AddClaim adds new claim to the storage.
func (st *DBStorage) AddClaim(_ context.Context, claim *etherman.Claim, dbTx interface{}) error {
	const addClaimSQL = "INSERT INTO claim (network_id, claim_index, orig_net, orig_addr, amount, dest_addr, block_id, tx_hash, rollup_index, mainnet_flag) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	e := st.getExecQuerier(dbTx)
	_, err := e.Exec(addClaimSQL, claim.NetworkID, claim.Index, claim.OriginalNetwork, claim.OriginalAddress, claim.Amount.String(), claim.DestinationAddress, claim.BlockID, claim.TxHash, claim.RollupIndex, claim.MainnetFlag)
	return err
}

// GetTokenMetadata gets the metadata of the dedicated token.
func (st *DBStorage) GetTokenMetadata(_ context.Context, networkID, destNet uint32, originalTokenAddr common.Address, dbTx interface{}) ([]byte, error) {
	var metadata []byte
	const getMetadataSQL = "SELECT metadata from deposit WHERE network_id = ? AND orig_addr = ? AND dest_net = ? AND metadata IS NOT NULL LIMIT 1"
	e := st.getExecQuerier(dbTx)
	err := e.QueryRow(getMetadataSQL, networkID, originalTokenAddr, destNet).Scan(&metadata)
	return metadata, err
}

// AddTokenWrapped adds new wrapped token to the storage.
func (st *DBStorage) AddTokenWrapped(ctx context.Context, tokenWrapped *etherman.TokenWrapped, dbTx interface{}) error {
	metadata, err := st.GetTokenMetadata(ctx, tokenWrapped.OriginalNetwork, tokenWrapped.NetworkID, tokenWrapped.OriginalTokenAddress, dbTx)
	var tokenMetadata *etherman.TokenMetadata
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		// if err == sql.ErrNoRows, this is due to missing the related deposit in the opposite network in fast sync mode.
		// ref: https://github.com/0xPolygonHermez/zkevm-bridge-service/issues/230
		tokenMetadata = &etherman.TokenMetadata{}
	} else {
		tokenMetadata, err = getDecodedToken(metadata)
		if err != nil {
			return err
		}
	}

	const addTokenWrappedSQL = "INSERT INTO token_wrapped (network_id, orig_net, orig_token_addr, wrapped_token_addr, block_id, name, symbol, decimals) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	e := st.getExecQuerier(dbTx)
	_, err = e.Exec(addTokenWrappedSQL, tokenWrapped.NetworkID, tokenWrapped.OriginalNetwork, tokenWrapped.OriginalTokenAddress, tokenWrapped.WrappedTokenAddress, tokenWrapped.BlockID, tokenMetadata.Name, tokenMetadata.Symbol, tokenMetadata.Decimals)
	return err
}

// Reset resets the state to a block for the given DB tx.
func (st *DBStorage) Reset(_ context.Context, blockNumber uint64, networkID uint32, dbTx interface{}) error {
	const resetSQL = "DELETE FROM block WHERE block_num > ? AND network_id = ?"
	e := st.getExecQuerier(dbTx)
	_, err := e.Exec(resetSQL, blockNumber, networkID)
	return err
}

// GetPreviousBlock gets the offset previous L1 block respect to latest.
func (st *DBStorage) GetPreviousBlock(_ context.Context, networkID uint32, offset uint64, dbTx interface{}) (etherman.Block, error) {
	var block etherman.Block
	const getPreviousBlockSQL = "SELECT block_num, block_hash, network_id FROM block WHERE network_id = ? ORDER BY block_num DESC LIMIT 1 OFFSET ?"
	e := st.getExecQuerier(dbTx)
	err := e.QueryRow(getPreviousBlockSQL, networkID, offset).Scan(&block.BlockNumber, &block.BlockHash, &block.NetworkID)
	if errors.Is(err, sql.ErrNoRows) {
		return etherman.Block{}, gerror.ErrStorageNotFound
	}
	return block, err
}

// GetNumberDeposits gets the number of  deposits.
func (st *DBStorage) GetNumberDeposits(_ context.Context, networkID uint32, blockNumber uint64, dbTx interface{}) (uint32, error) {
	var nDeposits int64
	const getNumDepositsSQL = "SELECT coalesce(MAX(deposit_cnt), -1) FROM deposit as d INNER JOIN block as b ON d.network_id = b.network_id AND d.block_id = b.id WHERE d.network_id = ? AND b.block_num <= ?"
	err := st.getExecQuerier(dbTx).QueryRow(getNumDepositsSQL, networkID, blockNumber).Scan(&nDeposits)
	return uint32(nDeposits + 1), err // nolint:gosec
}

// AddTrustedGlobalExitRoot adds new global exit root which comes from the trusted sequencer.
func (st *DBStorage) AddTrustedGlobalExitRoot(_ context.Context, trustedExitRoot *etherman.GlobalExitRoot, dbTx interface{}) (bool, error) {
	const addTrustedGerSQL = `
		INSERT INTO exit_root (block_id, global_exit_root, mainnet_exit_root, rollups_exit_root, network_id, allowed)
		VALUES (0, ?, ?, ?, ?, true)
		ON CONFLICT ON CONSTRAINT UC DO NOTHING;`
	res, err := st.getExecQuerier(dbTx).Exec(addTrustedGerSQL, trustedExitRoot.GlobalExitRoot, trustedExitRoot.ExitRoots[0], trustedExitRoot.ExitRoots[1], trustedExitRoot.NetworkID)
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}

// GetClaim gets a specific claim from the storage.
func (st *DBStorage) GetClaim(_ context.Context, depositCount, originNetworkID, networkID uint32, dbTx interface{}) (*etherman.Claim, error) {
	var (
		claim  etherman.Claim
		amount string
	)
	// origin rollup ID is calculated as follows:
	// // if mainnet_flag: 0
	// // else: rollup_index + 1
	// destination rollup ID == network_id: network that has received the claim, therefore, the destination rollupID of the claim

	const getClaimSQLOriginMainnet = `
	SELECT claim_index, orig_net, orig_addr, amount, dest_addr, block_id, network_id, tx_hash, rollup_index 
	FROM claim 
	WHERE claim_index = ? AND mainnet_flag AND network_id = ?;
	`

	const getClaimSQLOriginRollup = `
	SELECT claim_index, orig_net, orig_addr, amount, dest_addr, block_id, network_id, tx_hash, rollup_index 
	FROM claim 
	WHERE claim_index = ? AND NOT mainnet_flag AND rollup_index + 1 = ? AND network_id = ?;
	`
	var row *sql.Row
	if originNetworkID == 0 {
		claim.MainnetFlag = true
		row = st.getExecQuerier(dbTx).QueryRow(getClaimSQLOriginMainnet, depositCount, networkID)
	} else {
		row = st.getExecQuerier(dbTx).QueryRow(getClaimSQLOriginRollup, depositCount, originNetworkID, networkID)
	}

	err := row.Scan(&claim.Index, &claim.OriginalNetwork, &claim.OriginalAddress, &amount, &claim.DestinationAddress, &claim.BlockID, &claim.NetworkID, &claim.TxHash, &claim.RollupIndex)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, gerror.ErrStorageNotFound
	}
	claim.Amount, _ = new(big.Int).SetString(amount, 10) //nolint:mnd
	return &claim, err
}

// GetDeposit gets a specific deposit from the storage.
func (st *DBStorage) GetDeposit(_ context.Context, depositCounterUser, networkID uint32, dbTx interface{}) (*etherman.Deposit, error) {
	var (
		deposit etherman.Deposit
		amount  string
	)
	const getDepositSQL = "SELECT d.id, leaf_type, orig_net, orig_addr, amount, dest_net, dest_addr, deposit_cnt, block_id, b.block_num, d.network_id, tx_hash, metadata, ready_for_claim FROM deposit as d INNER JOIN block as b ON d.network_id = b.network_id AND d.block_id = b.id WHERE d.network_id = ? AND deposit_cnt = ?"
	err := st.getExecQuerier(dbTx).QueryRow(getDepositSQL, networkID, depositCounterUser).Scan(&deposit.Id, &deposit.LeafType, &deposit.OriginalNetwork, &deposit.OriginalAddress, &amount, &deposit.DestinationNetwork, &deposit.DestinationAddress, &deposit.DepositCount, &deposit.BlockID, &deposit.BlockNumber, &deposit.NetworkID, &deposit.TxHash, &deposit.Metadata, &deposit.ReadyForClaim)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, gerror.ErrStorageNotFound
	}
	deposit.Amount, _ = new(big.Int).SetString(amount, 10) //nolint:mnd

	return &deposit, err
}

// GetLatestExitRoot gets the latest global exit root.
func (st *DBStorage) GetLatestExitRoot(ctx context.Context, networkID, destNetwork uint32, dbTx interface{}) (*etherman.GlobalExitRoot, error) {
	if networkID == 0 {
		return st.GetLatestTrustedExitRoot(ctx, destNetwork, dbTx)
	}

	return st.GetLatestL1SyncedExitRoot(ctx, dbTx)
}

// GetLatestL1SyncedExitRoot gets the latest L1 synced global exit root.
func (st *DBStorage) GetLatestL1SyncedExitRoot(_ context.Context, dbTx interface{}) (*etherman.GlobalExitRoot, error) {
	var (
		ger       etherman.GlobalExitRoot
		mainnetExitRoot, rollupsExitRoot common.Hash
	)
	const getLatestL1SyncedExitRootSQL = "SELECT block_id, global_exit_root, mainnet_exit_root, rollups_exit_root FROM exit_root WHERE allowed = true AND block_id > 0 AND network_id = 0 ORDER BY id DESC LIMIT 1"
	err := st.getExecQuerier(dbTx).QueryRow(getLatestL1SyncedExitRootSQL).Scan(&ger.BlockID, &ger.GlobalExitRoot, &mainnetExitRoot, &rollupsExitRoot)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &ger, gerror.ErrStorageNotFound
		}
		return nil, err
	}
	ger.ExitRoots = []common.Hash{mainnetExitRoot, rollupsExitRoot}
	return &ger, nil
}

// GetL1ExitRootByGER gets the latest L1 synced global exit root.
func (st *DBStorage) GetL1ExitRootByGER(_ context.Context, ger common.Hash, dbTx interface{}) (*etherman.GlobalExitRoot, error) {
	var (
		gerData   etherman.GlobalExitRoot
		mainnetExitRoot, rollupsExitRoot common.Hash

	)
	const getSyncedExitRootSQL = "SELECT block_id, global_exit_root, mainnet_exit_root, rollups_exit_root FROM exit_root WHERE allowed = true AND block_id > 0 AND global_exit_root = ? AND network_id = 0 ORDER BY id DESC LIMIT 1"
	err := st.getExecQuerier(dbTx).QueryRow(getSyncedExitRootSQL, ger).Scan(&gerData.BlockID, &gerData.GlobalExitRoot, &mainnetExitRoot, &rollupsExitRoot)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &gerData, gerror.ErrStorageNotFound
		}
		return nil, err
	}
	gerData.ExitRoots = []common.Hash{mainnetExitRoot, rollupsExitRoot}
	return &gerData, nil
}

// GetL2ExitRootsByGER gets the global exit roots in all the L2 networks.
func (st *DBStorage) GetL2ExitRootsByGER(_ context.Context, ger common.Hash, dbTx interface{}) ([]etherman.GlobalExitRoot, error) {
	const getL2ExitRootsByGERSQL = "SELECT block_id, global_exit_root, network_id, id FROM exit_root WHERE allowed = true AND block_id > 0 AND global_exit_root = ? AND network_id != 0 AND mainnet_exit_root = X'0000000000000000000000000000000000000000000000000000000000000000' AND rollups_exit_root = X'0000000000000000000000000000000000000000000000000000000000000000' ORDER BY id DESC"
	rows, err := st.getExecQuerier(dbTx).Query(getL2ExitRootsByGERSQL, ger)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, gerror.ErrStorageNotFound
		}
		return nil, err
	}

	var gers []etherman.GlobalExitRoot

	for rows.Next() {
		var gerData etherman.GlobalExitRoot
		err = rows.Scan(&gerData.BlockID, &gerData.GlobalExitRoot, &gerData.NetworkID, &gerData.ID)
		if err != nil {
			return nil, err
		}

		gers = append(gers, gerData)
	}
	return gers, nil
}

// UpdateL2GER updates an L2 Ger in the storage.
func (st *DBStorage) UpdateL2GER(_ context.Context, ger etherman.GlobalExitRoot, dbTx interface{}) error {
	const updateL2GERSQL = `UPDATE exit_root
		SET mainnet_exit_root = ?, rollups_exit_root = ?
		WHERE id = ?`
	_, err := st.getExecQuerier(dbTx).Exec(updateL2GERSQL, ger.ExitRoots[0], ger.ExitRoots[1], ger.ID)
	return err
}

// GetLatestTrustedExitRoot gets the latest trusted global exit root.
func (st *DBStorage) GetLatestTrustedExitRoot(ctx context.Context, networkID uint32, dbTx interface{}) (*etherman.GlobalExitRoot, error) {
	var (
		ger       etherman.GlobalExitRoot
		mainnetExitRoot, rollupsExitRoot common.Hash
	)
	const getLatestTrustedExitRootSQL = "SELECT global_exit_root, mainnet_exit_root, rollups_exit_root, block_id, network_id, id FROM exit_root WHERE network_id = ? AND allowed = true ORDER BY id DESC LIMIT 1"
	err := st.getExecQuerier(dbTx).QueryRow(getLatestTrustedExitRootSQL, networkID).Scan(&ger.GlobalExitRoot, &mainnetExitRoot, &rollupsExitRoot, &ger.BlockID, &ger.NetworkID, &ger.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, gerror.ErrStorageNotFound
		}
		return nil, err
	}
	if mainnetExitRoot != (common.Hash{}) && rollupsExitRoot != (common.Hash{}) { //nolint:mnd
		ger.ExitRoots = []common.Hash{mainnetExitRoot, rollupsExitRoot}
	} else {
		// Query to look for the missing values
		l1GER, err := st.GetL1ExitRootByGER(ctx, ger.GlobalExitRoot, dbTx)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				log.Error("Missing L1Ger for the L2Ger entry")
				return nil, gerror.ErrStorageNotFound
			}
			return nil, err
		}
		ger.ExitRoots = l1GER.ExitRoots
	}
	return &ger, nil
}

// GetTokenWrapped gets a specific wrapped token.
func (st *DBStorage) GetTokenWrapped(ctx context.Context, originalNetwork uint32, originalTokenAddress common.Address, dbTx interface{}) (*etherman.TokenWrapped, error) {
	const getWrappedTokenSQL = "SELECT network_id, orig_net, orig_token_addr, wrapped_token_addr, block_id, name, symbol, decimals FROM token_wrapped WHERE orig_net = ? AND orig_token_addr = ?"

	var token etherman.TokenWrapped
	err := st.getExecQuerier(dbTx).QueryRow(getWrappedTokenSQL, originalNetwork, originalTokenAddress).Scan(&token.NetworkID, &token.OriginalNetwork, &token.OriginalTokenAddress, &token.WrappedTokenAddress, &token.BlockID, &token.Name, &token.Symbol, &token.Decimals)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, gerror.ErrStorageNotFound
	}

	// this is due to missing the related deposit in the opposite network in fast sync mode.
	// ref: https://github.com/0xPolygonHermez/zkevm-bridge-service/issues/230
	if token.Symbol == "" {
		metadata, err := st.GetTokenMetadata(ctx, token.OriginalNetwork, token.NetworkID, token.OriginalTokenAddress, dbTx)
		var tokenMetadata *etherman.TokenMetadata
		if err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
		} else {
			tokenMetadata, err = getDecodedToken(metadata)
			if err != nil {
				return nil, err
			}
			updateWrappedTokenSQL := "UPDATE token_wrapped SET name = ?, symbol = ?, decimals = ?  WHERE orig_net = ? AND orig_token_addr = ?" //nolint: gosec
			_, err = st.getExecQuerier(dbTx).Exec(updateWrappedTokenSQL, originalNetwork, originalTokenAddress, tokenMetadata.Name, tokenMetadata.Symbol, tokenMetadata.Decimals)
			if err != nil {
				return nil, err
			}
			token.Name, token.Symbol, token.Decimals = tokenMetadata.Name, tokenMetadata.Symbol, tokenMetadata.Decimals
		}
	}
	return &token, err
}

// GetDepositCountByRoot gets the deposit count by the root.
func (st *DBStorage) GetDepositCountByRoot(_ context.Context, root []byte, network uint32, dbTx interface{}) (uint32, error) {
	var depositCount uint32
	const getDepositCountByRootSQL = "SELECT deposit.deposit_cnt FROM root INNER JOIN deposit ON deposit.id = root.deposit_id WHERE root.root = ? AND root.network = ?"
	err := st.getExecQuerier(dbTx).QueryRow(getDepositCountByRootSQL, root, network).Scan(&depositCount)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, gerror.ErrStorageNotFound
	}
	return depositCount, nil
}

// CheckIfRootExists checks that the root exists on the db.
func (st *DBStorage) CheckIfRootExists(_ context.Context, root []byte, network uint32, dbTx interface{}) (bool, error) {
	var count uint
	const getDepositCountByRootSQL = "SELECT count(*) FROM root WHERE root = ? AND network = ?"
	err := st.getExecQuerier(dbTx).QueryRow(getDepositCountByRootSQL, root, network).Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return false, gerror.ErrStorageNotFound
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// GetRoot gets root by the deposit count from the merkle tree.
func (st *DBStorage) GetRoot(_ context.Context, depositCnt, network uint32, dbTx interface{}) ([]byte, error) {
	var root []byte
	const getRootByDepositCntSQL = "SELECT root FROM root inner join deposit on root.deposit_id = deposit.id WHERE deposit.deposit_cnt = ? AND network = ?"
	err := st.getExecQuerier(dbTx).QueryRow(getRootByDepositCntSQL, depositCnt, network).Scan(&root)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, gerror.ErrStorageNotFound
	}
	return root, err
}

// SetRoot store the root with deposit count to the storage.
func (st *DBStorage) SetRoot(_ context.Context, root []byte, depositID uint64, network uint32, dbTx interface{}) error {
	const setRootSQL = "INSERT INTO root (root, deposit_id, network) VALUES (?, ?, ?);"
	_, err := st.getExecQuerier(dbTx).Exec(setRootSQL, root, depositID, network)
	return err
}

// Get gets value of key from the merkle tree.
func (st *DBStorage) Get(_ context.Context, key []byte, dbTx interface{}) ([][]byte, error) {
	const getValueByKeySQL = "SELECT value FROM rht WHERE key = ?"
	var data [][]byte
	var left, right []byte
	err := st.getExecQuerier(dbTx).QueryRow(getValueByKeySQL, key).Scan(&left, &right)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, gerror.ErrStorageNotFound
	}
	data = append(data, left, right)
	return data, err
}

// Set inserts a key-value pair into the db.
// If record with such a key already exists its assumed that the value is correct,
// because it's a reverse hash table, and the key is a hash of the value
func (st *DBStorage) Set(_ context.Context, key []byte, value [][]byte, depositID uint64, dbTx interface{}) error {
	const setNodeSQL = "INSERT INTO rht (deposit_id, key, left_node, right_value) VALUES (?, ?, ?, ?)"
	_, err := st.getExecQuerier(dbTx).Exec(setNodeSQL, depositID, key, value[0], value[1])
	return err
}

// BulkSet is similar to Set, but it inserts multiple key-value pairs into the db.
func (st *DBStorage) BulkSet(_ context.Context, rows [][]interface{}, dbTx interface{}) error {
// The data structure within the type interface is adapted to the database while preserving the original function interface.
	var rs [][]interface{}
	for _, r := range rows {
		var newRow []interface{}
		for j, v := range r {
			if j != 1 {
				newRow = append(newRow, v)
			} else {
				newRow = append(newRow, v.([][]byte)[0])
				newRow = append(newRow, v.([][]byte)[1])
			}
		}
		rs = append(rs, newRow)
	}
	_, err := st.copyFrom("rht", []string{"key", "left_node", "right_node", "deposit_id"}, rs, dbTx)
	return err
}

// AddRollupExitLeaves inserts multiple entries into the db.
func (st *DBStorage) AddRollupExitLeaves(ctx context.Context, rows [][]interface{}, dbTx interface{}) error {
	_, err := st.copyFrom("rollup_exit", []string{"leaf", "rollup_id", "root", "block_id"}, rows, dbTx)
	return err
}

func (st *DBStorage) copyFrom(tableName string, columnNames []string, rows [][]interface{}, dbTx interface{}) (int64, error) {
	if len(rows) == 0 {
		return 0, nil // No rows to insert
	}

	// Build the SQL statement
	placeholders := make([]string, len(rows))
	for i := range rows {
		rowPlaceholders := make([]string, len(columnNames))
		for j := range columnNames {
			rowPlaceholders[j] = "?"
		}
		placeholders[i] = fmt.Sprintf("(%s)", strings.Join(rowPlaceholders, ", "))
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s", tableName, strings.Join(columnNames, ", "), strings.Join(placeholders, ", "))

	// Flatten the data for the Exec call
	args := make([]interface{}, 0, len(rows)*len(columnNames))
	for _, row := range rows {
		args = append(args, row...)
	}

	// Execute the insert statement
	result, err := st.getExecQuerier(dbTx).Exec(query, args...)
	if err != nil {
		return 0, err
	}

	// Get the number of rows affected
	copiedRows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return copiedRows, nil
}

// GetRollupExitLeavesByRoot gets the leaves of the rollupExitTree given a root
func (st *DBStorage) GetRollupExitLeavesByRoot(_ context.Context, root common.Hash, dbTx interface{}) ([]etherman.RollupExitLeaf, error) {
	const getLeavesSQL = "SELECT id, leaf, rollup_id, root, block_id FROM rollup_exit WHERE root = ? ORDER BY rollup_id ASC"
	rows, err := st.getExecQuerier(dbTx).Query(getLeavesSQL, root)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, gerror.ErrStorageNotFound
	} else if err != nil {
		return nil, err
	}
	var leaves []etherman.RollupExitLeaf

	for rows.Next() {
		var leaf etherman.RollupExitLeaf
		err = rows.Scan(&leaf.ID, &leaf.Leaf, &leaf.RollupId, &leaf.Root, &leaf.BlockID)
		if err != nil {
			return nil, err
		}
		leaves = append(leaves, leaf)
	}
	return leaves, nil
}

// IsRollupExitRoot checks if db contains the root
func (st *DBStorage) IsRollupExitRoot(_ context.Context, root common.Hash, dbTx interface{}) (bool, error) {
	const getLeavesSQL = "SELECT count(*) FROM rollup_exit WHERE root = ?"
	var count int
	err := st.getExecQuerier(dbTx).QueryRow(getLeavesSQL, root).Scan(&count)
	if errors.Is(err, sql.ErrNoRows) {
		return false, gerror.ErrStorageNotFound
	} else if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// GetLatestRollupExitLeaves gets the latest leaves of the rollupExitTree
func (st *DBStorage) GetLatestRollupExitLeaves(_ context.Context, dbTx interface{}) ([]etherman.RollupExitLeaf, error) {
	const getLeavesSQL = `SELECT distinct re.id, re.leaf, re.rollup_id, re.root, re.block_id
		FROM rollup_exit re
		INNER JOIN
			(SELECT distinct rollup_id, MAX(id) AS maxid
			FROM rollup_exit
			GROUP BY rollup_id) groupedre
		ON re.id = groupedre.maxid
		ORDER BY re.rollup_id asc;
	`
	rows, err := st.getExecQuerier(dbTx).Query(getLeavesSQL)
	if err != nil {
		return nil, err
	}
	var leaves []etherman.RollupExitLeaf

	for rows.Next() {
		var leaf etherman.RollupExitLeaf
		err = rows.Scan(&leaf.ID, &leaf.Leaf, &leaf.RollupId, &leaf.Root, &leaf.BlockID)
		if err != nil {
			return nil, err
		}
		leaves = append(leaves, leaf)
	}
	return leaves, nil
}

// GetLastDepositCount gets the last deposit count from the merkle tree.
func (st *DBStorage) GetLastDepositCount(_ context.Context, networkID uint32, dbTx interface{}) (uint32, error) {
	var depositCnt int64
	const getLastDepositCountSQL = "SELECT coalesce(MAX(deposit_cnt), -1) FROM deposit WHERE id = (SELECT coalesce(MAX(deposit_id), -1) FROM root WHERE network = ?)"
	err := st.getExecQuerier(dbTx).QueryRow(getLastDepositCountSQL, networkID).Scan(&depositCnt)
	if err != nil {
		return 0, nil
	}
	if depositCnt < 0 {
		return 0, gerror.ErrStorageNotFound
	}
	return uint32(depositCnt), nil // nolint:gosec
}

// GetClaimCount gets the claim count for the destination address.
func (st *DBStorage) GetClaimCount(_ context.Context, destAddr string, dbTx interface{}) (uint64, error) {
	const getClaimCountSQL = "SELECT COUNT(*) FROM claim WHERE dest_addr = ?"
	var claimCount uint64
	err := st.getExecQuerier(dbTx).QueryRow(getClaimCountSQL, common.FromHex(destAddr)).Scan(&claimCount)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, gerror.ErrStorageNotFound
	}
	return claimCount, err
}

// GetClaims gets the claim list which be smaller than claim_index.
func (st *DBStorage) GetClaims(_ context.Context, destAddr string, limit, offset uint32, dbTx interface{}) ([]*etherman.Claim, error) {
	const getClaimsSQL = "SELECT claim_index, orig_net, orig_addr, amount, dest_addr, block_id, network_id, tx_hash, rollup_index, mainnet_flag FROM claim WHERE dest_addr = ? ORDER BY block_id DESC LIMIT ? OFFSET ?"
	rows, err := st.getExecQuerier(dbTx).Query(getClaimsSQL, common.FromHex(destAddr), limit, offset)
	if err != nil {
		return nil, err
	}
	var claims []*etherman.Claim

	for rows.Next() {
		var (
			claim  etherman.Claim
			amount string
		)
		err = rows.Scan(&claim.Index, &claim.OriginalNetwork, &claim.OriginalAddress, &amount, &claim.DestinationAddress, &claim.BlockID, &claim.NetworkID, &claim.TxHash, &claim.RollupIndex, &claim.MainnetFlag)
		if err != nil {
			return nil, err
		}
		claim.Amount, _ = new(big.Int).SetString(amount, 10) //nolint:mnd
		claims = append(claims, &claim)
	}
	return claims, nil
}

// GetDeposits gets the deposit list which be smaller than depositCount.
func (st *DBStorage) GetDeposits(_ context.Context, destAddr string, limit, offset uint32, dbTx interface{}) ([]*etherman.Deposit, error) {
	const getDepositsSQL = "SELECT d.id, leaf_type, orig_net, orig_addr, amount, dest_net, dest_addr, deposit_cnt, block_id, b.block_num, d.network_id, tx_hash, metadata, ready_for_claim FROM deposit as d INNER JOIN block as b ON d.network_id = b.network_id AND d.block_id = b.id WHERE dest_addr = ? ORDER BY d.block_id DESC, d.deposit_cnt DESC LIMIT ? OFFSET ?"
	rows, err := st.getExecQuerier(dbTx).Query(getDepositsSQL, common.FromHex(destAddr), limit, offset)
	if err != nil {
		return nil, err
	}

	return parseDeposits(rows, true)
}

// GetDepositCount gets the deposit count for the destination address.
func (st *DBStorage) GetDepositCount(_ context.Context, destAddr string, dbTx interface{}) (uint64, error) {
	const getDepositCountSQL = "SELECT COUNT(*) FROM deposit WHERE dest_addr = ?"
	var depositCount uint64
	err := st.getExecQuerier(dbTx).QueryRow(getDepositCountSQL, common.FromHex(destAddr)).Scan(&depositCount)
	return depositCount, err
}

// UpdateL1DepositsStatus updates the ready_for_claim status of L1 deposits.
func (st *DBStorage) UpdateL1DepositsStatus(_ context.Context, exitRoot []byte, destinationNetwork uint32, dbTx interface{}) ([]*etherman.Deposit, error) {
	// Retrieve the updated rows
	selectQuery := `SELECT id, leaf_type, orig_net, orig_addr, amount, dest_net, dest_addr, deposit_cnt, block_id, network_id, tx_hash, metadata, ready_for_claim
		FROM deposit
		WHERE deposit_cnt <= (
			SELECT deposit.deposit_cnt FROM root INNER JOIN deposit ON deposit.id = root.deposit_id WHERE root.root = ? AND root.network = 0
		)
		AND network_id = 0 AND ready_for_claim = false AND dest_net = ?;
	`
	rows, err := st.getExecQuerier(dbTx).Query(selectQuery, exitRoot, destinationNetwork)
	if err != nil {
		return nil, err
	}
	const updateDepositsStatusSQL = `UPDATE deposit SET ready_for_claim = true 
		WHERE deposit_cnt <=
			(SELECT deposit.deposit_cnt FROM root INNER JOIN deposit ON deposit.id = root.deposit_id WHERE root.root = ? AND root.network = 0) 
			AND network_id = 0 AND ready_for_claim = false AND dest_net = ?;`
	_, err = st.getExecQuerier(dbTx).Exec(updateDepositsStatusSQL, exitRoot, destinationNetwork)
	if err != nil {
		return nil, err
	}
	return parseDeposits(rows, false)
}

// UpdateL2DepositsStatus updates the ready_for_claim status of L2 deposits.
func (st *DBStorage) UpdateL2DepositsStatus(_ context.Context, exitRoot []byte, rollupID, networkID uint32, dbTx interface{}) error {
	const updateL2DepositsStatusSQL = `UPDATE deposit SET ready_for_claim = true
		WHERE deposit_cnt <=
		(SELECT deposit.deposit_cnt FROM root INNER JOIN deposit ON deposit.id = root.deposit_id WHERE root.root = (select leaf from rollup_exit where root = ? and rollup_id = ?) AND root.network = ?)
			AND network_id = ? AND ready_for_claim = false;`
	_, err := st.getExecQuerier(dbTx).Exec(updateL2DepositsStatusSQL, exitRoot, rollupID, networkID)
	return err
}

// GetDepositsFromOtherL2ToClaim returns L2 deposits whose destination is an specific L2
func (st *DBStorage) GetDepositsFromOtherL2ToClaim(_ context.Context, destinationNetwork uint32, dbTx interface{}) ([]*etherman.Deposit, error) {
	const getL2DepositsToClaimStatusSQL = `select deposit.id, deposit.leaf_type, deposit.orig_net, deposit.orig_addr, deposit.amount, deposit.dest_net, deposit.dest_addr, deposit.deposit_cnt, deposit.block_id, deposit.network_id, deposit.tx_hash, deposit.metadata, deposit.ready_for_claim FROM deposit where deposit.deposit_cnt not in (select claim_index FROM claim where claim.network_id = ?) and deposit.network_id !=0 and deposit.dest_net = ? and ready_for_claim =true order by deposit.id desc;`
	rows, err := st.getExecQuerier(dbTx).Query(getL2DepositsToClaimStatusSQL, destinationNetwork)
	if err != nil {
		return nil, err
	}
	return parseDeposits(rows, false)
}

// GetLatestTrustedGERByDeposit return the latest trusted ger for an specific deposit
func (st *DBStorage) GetLatestTrustedGERByDeposit(_ context.Context, depositCnt, networkID, destinationNetwork uint32, dbTx interface{}) (common.Hash, error) {
	const getLatestTrustedGERByDeposit = `SELECT exit_root.global_exit_root FROM deposit inner join root on root.deposit_id = deposit.id inner join rollup_exit on rollup_exit.leaf = root.root inner join exit_root on exit_root.exit_roots[2]= rollup_exit.root where exit_root.allowed = true and deposit_cnt = ? and deposit.network_id = ? and dest_net = ? and rollup_exit.rollup_id = ? and exit_root.network_id = deposit.dest_net order by exit_root.id desc limit 1`
	var ger common.Hash
	err := st.getExecQuerier(dbTx).QueryRow(getLatestTrustedGERByDeposit, depositCnt, networkID, destinationNetwork).Scan(&ger)
	if errors.Is(err, sql.ErrNoRows) {
		return common.Hash{}, gerror.ErrStorageNotFound
	}
	return ger, err
}

// AddClaimTx adds a claim monitored transaction to the storage.
func (st *DBStorage) AddClaimTx(_ context.Context, mTx ctmtypes.MonitoredTx, dbTx interface{}) error {
	const addMonitoredTxSQL = `INSERT INTO monitored_txs 
		(deposit_id, from_addr, to_addr, nonce, value, data, gas, status, history, created_at, updated_at, group_id, global_exit_root)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := st.getExecQuerier(dbTx).Exec(addMonitoredTxSQL, mTx.DepositID, mTx.From, mTx.To, mTx.Nonce, mTx.Value.String(),
		mTx.Data, mTx.Gas, mTx.Status, mTx.HistoryHashesString(), time.Now().UTC(), time.Now().UTC(), mTx.GroupID, mTx.GlobalExitRoot)
	return err
}

// UpdateClaimTx updates a claim monitored transaction in the storage.
func (st *DBStorage) UpdateClaimTx(_ context.Context, mTx ctmtypes.MonitoredTx, dbTx interface{}) error {
	const updateMonitoredTxSQL = `UPDATE monitored_txs 
		SET from_addr = ?
		, to_addr = ?
		, nonce = ?
		, value = ?
		, data = ?
		, gas = ?
		, status = ?
		, history = ?
		, updated_at = ?
		, group_id = ?
		WHERE deposit_id = ?`
	_, err := st.getExecQuerier(dbTx).Exec(updateMonitoredTxSQL, mTx.DepositID, mTx.From, mTx.To, mTx.Nonce, mTx.Value.String(),
		mTx.Data, mTx.Gas, mTx.Status, mTx.HistoryHashesString(), time.Now().UTC(), mTx.GroupID)
	return err
}

// GetClaimTxsByStatus gets the monitored transactions by status.
func (st *DBStorage) GetClaimTxsByStatus(_ context.Context, statuses []ctmtypes.MonitoredTxStatus, rollupID uint32, dbTx interface{}) ([]ctmtypes.MonitoredTx, error) {
	placeholders := strings.Repeat("?,", len(statuses)-1) + "?"
	getMonitoredTxsSQL := fmt.Sprintf("SELECT deposit_id, from_addr, to_addr, nonce, value, data, gas, status, history, created_at, updated_at, group_id, global_exit_root FROM monitored_txs INNER JOIN deposit ON deposit.id = monitored_txs.deposit_id WHERE status IN(%s) AND deposit.dest_net = ? ORDER BY created_at ASC", placeholders)
	StatusesInterfaces := make([]interface{}, len(statuses))
	for i, status := range statuses {
		StatusesInterfaces[i] = status
	}
	rows, err := st.getExecQuerier(dbTx).Query(getMonitoredTxsSQL, append(StatusesInterfaces, rollupID)...)
	if errors.Is(err, sql.ErrNoRows) {
		return []ctmtypes.MonitoredTx{}, nil
	} else if err != nil {
		return nil, err
	}

	var mTxs []ctmtypes.MonitoredTx
	for rows.Next() {
		var (
			value, historyStr string
		)
		mTx := ctmtypes.MonitoredTx{}
		err = rows.Scan(&mTx.DepositID, &mTx.From, &mTx.To, &mTx.Nonce, &value, &mTx.Data, &mTx.Gas, &mTx.Status, &historyStr, &mTx.CreatedAt, &mTx.UpdatedAt, &mTx.GroupID, &mTx.GlobalExitRoot)
		if err != nil {
			return mTxs, err
		}
		mTx.Value, _ = new(big.Int).SetString(value, 10) //nolint:mnd
		mTx.History = historyStrToHashArray(historyStr)
		mTxs = append(mTxs, mTx)
	}

	return mTxs, nil
}

func historyStrToHashArray(historyStr string) map[common.Hash]bool {
	hStr := strings.Split(historyStr, ",")
	history := make(map[common.Hash]bool)
	for _, h := range hStr {
		history[common.HexToHash(h)] = true
	}
	return history
}

// GetPendingDepositsToClaim gets the deposit list which is not claimed in the destination network.
func (st *DBStorage) GetPendingDepositsToClaim(_ context.Context, destAddress common.Address, destNetwork, leafType, limit, offset uint32, dbTx interface{}) ([]*etherman.Deposit, uint64, error) {
	desAddrSQL := ""
	if destAddress != (common.Address{}) {
		str := strings.TrimPrefix(destAddress.String(), "0x")
		desAddrSQL = "AND dest_addr = X'" + str + "'"
	}
	getNumberPendingDepositsToClaimSQL := "SELECT count(*) FROM deposit WHERE dest_net = ? AND ready_for_claim = true AND leaf_type = ? " + desAddrSQL + " AND deposit_cnt NOT IN (SELECT claim_index FROM claim WHERE claim.network_id = ?)"
	var totalCount uint64
	err := st.getExecQuerier(dbTx).QueryRow(getNumberPendingDepositsToClaimSQL, destNetwork, leafType, destNetwork).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}
	getPendingDepositsToClaimSQL := "SELECT d.id, leaf_type, orig_net, orig_addr, amount, dest_net, dest_addr, deposit_cnt, block_id, b.block_num, d.network_id, tx_hash, metadata, ready_for_claim FROM deposit AS d INNER JOIN block AS b ON d.block_id = b.id WHERE dest_net = ? AND ready_for_claim = true AND leaf_type = ? " + desAddrSQL + " AND d.deposit_cnt NOT IN (SELECT claim_index FROM claim WHERE claim.network_id = ?) ORDER BY d.deposit_cnt ASC LIMIT ? OFFSET ?"
	rows, err := st.getExecQuerier(dbTx).Query(getPendingDepositsToClaimSQL, destNetwork, leafType, destNetwork, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	deposits, err := parseDeposits(rows, true)
	if err != nil {
		return nil, 0, err
	}
	return deposits, totalCount, nil
}

func (st *DBStorage) AddRemoveL2GER(ctx context.Context, globalExitRoot etherman.GlobalExitRoot, dbTx interface{}) error {
	const insertRemoveGERSQL = `INSERT INTO remove_exit_root 
		(block_id, global_exit_root, network_id)
		VALUES (?, ?, ?)`
	_, err := st.getExecQuerier(dbTx).Exec(insertRemoveGERSQL, globalExitRoot.BlockID, globalExitRoot.GlobalExitRoot, globalExitRoot.NetworkID)
	if err != nil {
		return err
	}

	// Modify the allowed column in the exit_root table for this globalExitRoot
	const updateGERStatusSQL = "UPDATE exit_root SET allowed = false WHERE global_exit_root = ? AND network_id = ?;"
	var gerID uint64
	_, err = st.getExecQuerier(dbTx).Exec(updateGERStatusSQL, globalExitRoot.GlobalExitRoot, globalExitRoot.NetworkID)
	if err != nil {
		return err
	}

	selectQuery := "SELECT id FROM exit_root WHERE global_exit_root = ? AND network_id = ? AND allowed = false;"
	err = st.getExecQuerier(dbTx).QueryRow(selectQuery, globalExitRoot.GlobalExitRoot, globalExitRoot.NetworkID).Scan(&gerID)
	if err != nil {
		log.Fatalf("Failed to execute select query: %v", err)
	}

	// Look for the deposits where the ready for claim flag needs to be modified
	getPreviousGERSQL := "SELECT global_exit_root FROM exit_root WHERE network_id = ? AND id < ? ORDER BY id DESC LIMIT 1"
	var prevGer common.Hash
	err = st.getExecQuerier(dbTx).QueryRow(getPreviousGERSQL, globalExitRoot.NetworkID, gerID).Scan(&prevGer)
	if err != nil {
		return err
	}
	prevDepositID, _, err := st.GetDepositCountByGER(ctx, prevGer, globalExitRoot.NetworkID, false, dbTx)
	if errors.Is(err, gerror.ErrStorageNotFound) {
		log.Warnf("No previous deposit found using this prevGER: %s in L1, NetworkID: %d, currentGER: %s", prevGer.String(), globalExitRoot.NetworkID, globalExitRoot.GlobalExitRoot.String())
		prevDepositID = 0
	} else if err != nil {
		return err
	}
	CurrentDepositID, _, err := st.GetDepositCountByGER(ctx, globalExitRoot.GlobalExitRoot, globalExitRoot.NetworkID, false, dbTx)
	if errors.Is(err, gerror.ErrStorageNotFound) {
		log.Warnf("No deposit found for this GER: %s in L1, NetworkID: %d", globalExitRoot.GlobalExitRoot.String(), globalExitRoot.NetworkID)
		CurrentDepositID = math.MaxUint64
	} else if err != nil {
		return err
	}
	// Disable the ready for claim flag for this deposits
	const updateDepositsFlagSQL = `UPDATE deposit SET ready_for_claim = false WHERE id > ? AND id <= ? AND network_id = 0 AND dest_net = ?`
	_, err = st.getExecQuerier(dbTx).Exec(updateDepositsFlagSQL, prevDepositID, CurrentDepositID, globalExitRoot.NetworkID)
	return err
}

// GetDepositCountByGER gets the deposit count by the GER.
func (st *DBStorage) GetDepositCountByGER(_ context.Context, ger common.Hash, network uint32, rollupsTree bool, dbTx interface{}) (uint64, uint32, error) {
	var (
		depositID               uint64
		depositCnt, rootNetwork uint32
	)
	arrayIndex := 1
	if rollupsTree {
		arrayIndex = 2
		rootNetwork = network
	}

	const getDepositCountByGERSQL = "SELECT deposit.id, deposit.deposit_cnt FROM deposit INNER JOIN root ON deposit.id = root.deposit_id inner join exit_root on root.root = exit_root.exit_roots[?] WHERE exit_root.global_exit_root = ? and exit_root.network_id = ? AND root.network = ?"
	err := st.getExecQuerier(dbTx).QueryRow(getDepositCountByGERSQL, arrayIndex, ger, network, rootNetwork).Scan(&depositID, &depositCnt)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, 0, gerror.ErrStorageNotFound
	}
	return depositID, depositCnt, err
}

// UpdateDepositsStatusForTesting updates the ready_for_claim status of all deposits for testing.
func (st *DBStorage) UpdateDepositsStatusForTesting(_ context.Context, dbTx interface{}) error {
	const updateDepositsStatusSQL = "UPDATE deposit SET ready_for_claim = true;"
	_, err := st.getExecQuerier(dbTx).Exec(updateDepositsStatusSQL)
	return err
}

// UpdateBlocksForTesting updates the hash of blocks.
func (st *DBStorage) UpdateBlocksForTesting(_ context.Context, networkID uint32, blockNum uint64, dbTx interface{}) error {
	const updateBlocksSQL = "UPDATE block SET block_hash = SUBSTRING(block_hash FROM 1 FOR LENGTH(block_hash)-1) || '\x61' WHERE network_id = ? AND block_num >= ?"
	_, err := st.getExecQuerier(dbTx).Exec(updateBlocksSQL, networkID, blockNum)
	return err
}

// // QueryRowTesting is used for testing purposes.
// func (st *DBStorage) QueryRowTesting(_ context.Context, query string, args []interface{}, dbTx interface{}) (interface{}, error) {
// 	e := st.getExecQuerier(dbTx)
// 	var result interface{}
// 	err := e.QueryRow(query, args...).Scan(&result)
// 	return result, err
// }

func parseDeposits(rows *sql.Rows, needBlockNum bool) ([]*etherman.Deposit, error) {
	var deposits []*etherman.Deposit
	for rows.Next() {
		var (
			deposit etherman.Deposit
			amount  string
			err     error
		)
		if needBlockNum {
			err = rows.Scan(&deposit.Id, &deposit.LeafType, &deposit.OriginalNetwork, &deposit.OriginalAddress, &amount, &deposit.DestinationNetwork, &deposit.DestinationAddress, &deposit.DepositCount, &deposit.BlockID, &deposit.BlockNumber, &deposit.NetworkID, &deposit.TxHash, &deposit.Metadata, &deposit.ReadyForClaim)
		} else {
			err = rows.Scan(&deposit.Id, &deposit.LeafType, &deposit.OriginalNetwork, &deposit.OriginalAddress, &amount, &deposit.DestinationNetwork, &deposit.DestinationAddress, &deposit.DepositCount, &deposit.BlockID, &deposit.NetworkID, &deposit.TxHash, &deposit.Metadata, &deposit.ReadyForClaim)
		}
		if err != nil {
			return nil, err
		}
		deposit.Amount, _ = new(big.Int).SetString(amount, 10) //nolint:mnd
		deposits = append(deposits, &deposit)
	}
	return deposits, nil
}
