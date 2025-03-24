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
	"github.com/lib/pq"
)

// DBStorage implements the Storage interface
type DBStorage struct {
	DB *sql.DB
}

// getExecQuerier determines which execQuerier to use, dbTx or the main pgxpool
func (st *DBStorage) getExecQuerier(dbTx *sql.Tx) execQuerier {
	if dbTx != nil {
		return dbTx
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
func (st *DBStorage) Rollback(ctx context.Context, dbTx *sql.Tx) error {
	if dbTx != nil {
		return dbTx.Rollback()
	}

	return gerror.ErrNilDBTransaction
}

// Commit commits a db transaction.
func (st *DBStorage) Commit(ctx context.Context, dbTx *sql.Tx) error {
	if dbTx != nil {
		return dbTx.Commit()
	}
	return gerror.ErrNilDBTransaction
}

// BeginDBTransaction starts a transaction block.
func (st *DBStorage) BeginDBTransaction(ctx context.Context) (*sql.Tx, error) {
	return st.DB.Begin()
}

// GetLastBlock gets the last block.
func (st *DBStorage) GetLastBlock(ctx context.Context, networkID uint32, dbTx *sql.Tx) (*etherman.Block, error) {
	var block etherman.Block
	const getLastBlockSQL = "SELECT id, block_num, block_hash, network_id FROM sync.block where network_id = $1 ORDER BY block_num DESC LIMIT 1"

	e := st.getExecQuerier(dbTx)
	err := e.QueryRow(getLastBlockSQL, networkID).Scan(&block.ID, &block.BlockNumber, &block.BlockHash, &block.NetworkID)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, gerror.ErrStorageNotFound
	}

	return &block, err
}

// AddBlock adds a new block to the storage.
func (st *DBStorage) AddBlock(ctx context.Context, block *etherman.Block, dbTx *sql.Tx) (uint64, error) {
	var blockID uint64
	const addBlockSQL = `WITH block_id AS 
		(INSERT INTO sync.block (block_num, block_hash, network_id) 
		VALUES ($1, $2, $3) ON CONFLICT (block_hash) DO NOTHING RETURNING id)
		SELECT * from block_id
		UNION ALL
		SELECT id FROM sync.block WHERE block_hash = $2;`
	e := st.getExecQuerier(dbTx)
	err := e.QueryRow(addBlockSQL, block.BlockNumber, block.BlockHash, block.NetworkID).Scan(&blockID)

	if err == sql.ErrNoRows {
		err = nil
	}

	return blockID, err
}

// AddGlobalExitRoot adds a new ExitRoot to the db.
func (st *DBStorage) AddGlobalExitRoot(ctx context.Context, exitRoot *etherman.GlobalExitRoot, dbTx *sql.Tx) error {
	const addExitRootSQL = "INSERT INTO sync.exit_root (block_id, global_exit_root, exit_roots, network_id, allowed) VALUES ($1, $2, $3, $4, true)"
	exitRoots := [][]byte{}
	if len(exitRoot.ExitRoots) != 0 {
		exitRoots = [][]byte{exitRoot.ExitRoots[0][:], exitRoot.ExitRoots[1][:]}
	}
	e := st.getExecQuerier(dbTx)
	_, err := e.Exec(addExitRootSQL, exitRoot.BlockID, exitRoot.GlobalExitRoot, pq.Array(exitRoots), exitRoot.NetworkID)
	return err
}

// AddDeposit adds new deposit to the storage.
func (st *DBStorage) AddDeposit(ctx context.Context, deposit *etherman.Deposit, dbTx *sql.Tx) (uint64, error) {
	const addDepositSQL = "INSERT INTO sync.deposit (leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
	e := st.getExecQuerier(dbTx)
	var depositID uint64
	err := e.QueryRow(addDepositSQL, deposit.LeafType, deposit.NetworkID, deposit.OriginalNetwork, deposit.OriginalAddress, deposit.Amount.String(), deposit.DestinationNetwork, deposit.DestinationAddress, deposit.BlockID, deposit.DepositCount, deposit.TxHash, deposit.Metadata).Scan(&depositID)
	return depositID, err
}

// AddClaim adds new claim to the storage.
func (st *DBStorage) AddClaim(ctx context.Context, claim *etherman.Claim, dbTx *sql.Tx) error {
	const addClaimSQL = "INSERT INTO sync.claim (network_id, index, orig_net, orig_addr, amount, dest_addr, block_id, tx_hash, rollup_index, mainnet_flag) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	e := st.getExecQuerier(dbTx)
	_, err := e.Exec(addClaimSQL, claim.NetworkID, claim.Index, claim.OriginalNetwork, claim.OriginalAddress, claim.Amount.String(), claim.DestinationAddress, claim.BlockID, claim.TxHash, claim.RollupIndex, claim.MainnetFlag)
	return err
}

// GetTokenMetadata gets the metadata of the dedicated token.
func (st *DBStorage) GetTokenMetadata(ctx context.Context, networkID, destNet uint32, originalTokenAddr common.Address, dbTx *sql.Tx) ([]byte, error) {
	var metadata []byte
	const getMetadataSQL = "SELECT metadata from sync.deposit WHERE network_id = $1 AND orig_addr = $2 AND dest_net = $3 AND metadata IS NOT NULL LIMIT 1"
	e := st.getExecQuerier(dbTx)
	err := e.QueryRow(getMetadataSQL, networkID, originalTokenAddr, destNet).Scan(&metadata)
	return metadata, err
}

// AddTokenWrapped adds new wrapped token to the storage.
func (st *DBStorage) AddTokenWrapped(ctx context.Context, tokenWrapped *etherman.TokenWrapped, dbTx *sql.Tx) error {
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

	const addTokenWrappedSQL = "INSERT INTO sync.token_wrapped (network_id, orig_net, orig_token_addr, wrapped_token_addr, block_id, name, symbol, decimals) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	e := st.getExecQuerier(dbTx)
	_, err = e.Exec(addTokenWrappedSQL, tokenWrapped.NetworkID, tokenWrapped.OriginalNetwork, tokenWrapped.OriginalTokenAddress, tokenWrapped.WrappedTokenAddress, tokenWrapped.BlockID, tokenMetadata.Name, tokenMetadata.Symbol, tokenMetadata.Decimals)
	return err
}

// Reset resets the state to a block for the given DB tx.
func (st *DBStorage) Reset(ctx context.Context, blockNumber uint64, networkID uint32, dbTx *sql.Tx) error {
	const resetSQL = "DELETE FROM sync.block WHERE block_num > $1 AND network_id = $2"
	e := st.getExecQuerier(dbTx)
	_, err := e.Exec(resetSQL, blockNumber, networkID)
	return err
}

// GetPreviousBlock gets the offset previous L1 block respect to latest.
func (st *DBStorage) GetPreviousBlock(ctx context.Context, networkID uint32, offset uint64, dbTx *sql.Tx) (etherman.Block, error) {
	var block etherman.Block
	const getPreviousBlockSQL = "SELECT block_num, block_hash, network_id FROM sync.block WHERE network_id = $1 ORDER BY block_num DESC LIMIT 1 OFFSET $2"
	e := st.getExecQuerier(dbTx)
	err := e.QueryRow(getPreviousBlockSQL, networkID, offset).Scan(&block.BlockNumber, &block.BlockHash, &block.NetworkID)
	if errors.Is(err, sql.ErrNoRows) {
		return etherman.Block{}, gerror.ErrStorageNotFound
	}
	return block, err
}

// GetNumberDeposits gets the number of  deposits.
func (st *DBStorage) GetNumberDeposits(ctx context.Context, networkID uint32, blockNumber uint64, dbTx *sql.Tx) (uint32, error) {
	var nDeposits int64
	const getNumDepositsSQL = "SELECT coalesce(MAX(deposit_cnt), -1) FROM sync.deposit as d INNER JOIN sync.block as b ON d.network_id = b.network_id AND d.block_id = b.id WHERE d.network_id = $1 AND b.block_num <= $2"
	err := st.getExecQuerier(dbTx).QueryRow(getNumDepositsSQL, networkID, blockNumber).Scan(&nDeposits)
	return uint32(nDeposits + 1), err // nolint:gosec
}

// AddTrustedGlobalExitRoot adds new global exit root which comes from the trusted sequencer.
func (st *DBStorage) AddTrustedGlobalExitRoot(ctx context.Context, trustedExitRoot *etherman.GlobalExitRoot, dbTx *sql.Tx) (bool, error) {
	const addTrustedGerSQL = `
		INSERT INTO sync.exit_root (block_id, global_exit_root, exit_roots, network_id, allowed)
		VALUES (0, $1, $2, $3, true)
		ON CONFLICT ON CONSTRAINT UC DO NOTHING;`
	res, err := st.getExecQuerier(dbTx).Exec(addTrustedGerSQL, trustedExitRoot.GlobalExitRoot, pq.Array([][]byte{trustedExitRoot.ExitRoots[0][:], trustedExitRoot.ExitRoots[1][:]}), trustedExitRoot.NetworkID)
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}

// GetClaim gets a specific claim from the storage.
func (st *DBStorage) GetClaim(ctx context.Context, depositCount, originNetworkID, networkID uint32, dbTx *sql.Tx) (*etherman.Claim, error) {
	var (
		claim  etherman.Claim
		amount string
	)
	// origin rollup ID is calculated as follows:
	// // if mainnet_flag: 0
	// // else: rollup_index + 1
	// destination rollup ID == network_id: network that has received the claim, therefore, the destination rollupID of the claim

	const getClaimSQLOriginMainnet = `
	SELECT index, orig_net, orig_addr, amount, dest_addr, block_id, network_id, tx_hash, rollup_index 
	FROM sync.claim 
	WHERE index = $1 AND mainnet_flag AND network_id = $2;
	`

	const getClaimSQLOriginRollup = `
	SELECT index, orig_net, orig_addr, amount, dest_addr, block_id, network_id, tx_hash, rollup_index 
	FROM sync.claim 
	WHERE index = $1 AND NOT mainnet_flag AND rollup_index + 1 = $2 AND network_id = $3;
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
func (st *DBStorage) GetDeposit(ctx context.Context, depositCounterUser, networkID uint32, dbTx *sql.Tx) (*etherman.Deposit, error) {
	var (
		deposit etherman.Deposit
		amount  string
	)
	const getDepositSQL = "SELECT d.id, leaf_type, orig_net, orig_addr, amount, dest_net, dest_addr, deposit_cnt, block_id, b.block_num, d.network_id, tx_hash, metadata, ready_for_claim FROM sync.deposit as d INNER JOIN sync.block as b ON d.network_id = b.network_id AND d.block_id = b.id WHERE d.network_id = $1 AND deposit_cnt = $2"
	err := st.getExecQuerier(dbTx).QueryRow(getDepositSQL, networkID, depositCounterUser).Scan(&deposit.Id, &deposit.LeafType, &deposit.OriginalNetwork, &deposit.OriginalAddress, &amount, &deposit.DestinationNetwork, &deposit.DestinationAddress, &deposit.DepositCount, &deposit.BlockID, &deposit.BlockNumber, &deposit.NetworkID, &deposit.TxHash, &deposit.Metadata, &deposit.ReadyForClaim)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, gerror.ErrStorageNotFound
	}
	deposit.Amount, _ = new(big.Int).SetString(amount, 10) //nolint:mnd

	return &deposit, err
}

// GetLatestExitRoot gets the latest global exit root.
func (st *DBStorage) GetLatestExitRoot(ctx context.Context, networkID, destNetwork uint32, dbTx *sql.Tx) (*etherman.GlobalExitRoot, error) {
	if networkID == 0 {
		return st.GetLatestTrustedExitRoot(ctx, destNetwork, dbTx)
	}

	return st.GetLatestL1SyncedExitRoot(ctx, dbTx)
}

// GetLatestL1SyncedExitRoot gets the latest L1 synced global exit root.
func (st *DBStorage) GetLatestL1SyncedExitRoot(ctx context.Context, dbTx *sql.Tx) (*etherman.GlobalExitRoot, error) {
	var (
		ger       etherman.GlobalExitRoot
		exitRoots [][]byte
	)
	const getLatestL1SyncedExitRootSQL = "SELECT block_id, global_exit_root, exit_roots FROM sync.exit_root WHERE allowed = true AND block_id > 0 AND network_id = 0 ORDER BY id DESC LIMIT 1"
	err := st.getExecQuerier(dbTx).QueryRow(getLatestL1SyncedExitRootSQL).Scan(&ger.BlockID, &ger.GlobalExitRoot, pq.Array(&exitRoots))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &ger, gerror.ErrStorageNotFound
		}
		return nil, err
	}
	ger.ExitRoots = []common.Hash{common.BytesToHash(exitRoots[0]), common.BytesToHash(exitRoots[1])}
	return &ger, nil
}

// GetL1ExitRootByGER gets the latest L1 synced global exit root.
func (st *DBStorage) GetL1ExitRootByGER(ctx context.Context, ger common.Hash, dbTx *sql.Tx) (*etherman.GlobalExitRoot, error) {
	var (
		gerData   etherman.GlobalExitRoot
		exitRoots [][]byte
	)
	const getSyncedExitRootSQL = "SELECT block_id, global_exit_root, exit_roots FROM sync.exit_root WHERE allowed = true AND block_id > 0 AND global_exit_root = $1 AND network_id = 0 ORDER BY id DESC LIMIT 1"
	err := st.getExecQuerier(dbTx).QueryRow(getSyncedExitRootSQL, ger).Scan(&gerData.BlockID, &gerData.GlobalExitRoot, pq.Array(&exitRoots))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &gerData, gerror.ErrStorageNotFound
		}
		return nil, err
	}
	gerData.ExitRoots = []common.Hash{common.BytesToHash(exitRoots[0]), common.BytesToHash(exitRoots[1])}
	return &gerData, nil
}

// GetL2ExitRootsByGER gets the global exit roots in all the L2 networks.
func (st *DBStorage) GetL2ExitRootsByGER(ctx context.Context, ger common.Hash, dbTx *sql.Tx) ([]etherman.GlobalExitRoot, error) {
	const getL2ExitRootsByGERSQL = "SELECT block_id, global_exit_root, network_id, id FROM sync.exit_root WHERE allowed = true AND block_id > 0 AND global_exit_root = $1 AND network_id != 0 AND cardinality(exit_roots) = 0 ORDER BY id DESC"
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
func (st *DBStorage) UpdateL2GER(ctx context.Context, ger etherman.GlobalExitRoot, dbTx *sql.Tx) error {
	const updateL2GERSQL = `UPDATE sync.exit_root
		SET exit_roots = $2
		WHERE id = $1`
	_, err := st.getExecQuerier(dbTx).Exec(updateL2GERSQL, ger.ID, pq.Array([][]byte{ger.ExitRoots[0][:], ger.ExitRoots[1][:]}))
	return err
}

// GetLatestTrustedExitRoot gets the latest trusted global exit root.
func (st *DBStorage) GetLatestTrustedExitRoot(ctx context.Context, networkID uint32, dbTx *sql.Tx) (*etherman.GlobalExitRoot, error) {
	var (
		ger       etherman.GlobalExitRoot
		exitRoots [][]byte
	)
	const getLatestTrustedExitRootSQL = "SELECT global_exit_root, exit_roots, block_id, network_id, id FROM sync.exit_root WHERE network_id = $1 AND allowed = true ORDER BY id DESC LIMIT 1"
	err := st.getExecQuerier(dbTx).QueryRow(getLatestTrustedExitRootSQL, networkID).Scan(&ger.GlobalExitRoot, pq.Array(&exitRoots), &ger.BlockID, &ger.NetworkID, &ger.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, gerror.ErrStorageNotFound
		}
		return nil, err
	}
	if len(exitRoots) == 2 { //nolint:mnd
		ger.ExitRoots = []common.Hash{common.BytesToHash(exitRoots[0]), common.BytesToHash(exitRoots[1])}
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
func (st *DBStorage) GetTokenWrapped(ctx context.Context, originalNetwork uint32, originalTokenAddress common.Address, dbTx *sql.Tx) (*etherman.TokenWrapped, error) {
	const getWrappedTokenSQL = "SELECT network_id, orig_net, orig_token_addr, wrapped_token_addr, block_id, name, symbol, decimals FROM sync.token_wrapped WHERE orig_net = $1 AND orig_token_addr = $2"

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
			updateWrappedTokenSQL := "UPDATE sync.token_wrapped SET name = $3, symbol = $4, decimals = $5  WHERE orig_net = $1 AND orig_token_addr = $2" //nolint: gosec
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
func (st *DBStorage) GetDepositCountByRoot(ctx context.Context, root []byte, network uint32, dbTx *sql.Tx) (uint32, error) {
	var depositCount uint32
	const getDepositCountByRootSQL = "SELECT sync.deposit.deposit_cnt FROM mt.root INNER JOIN sync.deposit ON sync.deposit.id = mt.root.deposit_id WHERE mt.root.root = $1 AND mt.root.network = $2"
	err := st.getExecQuerier(dbTx).QueryRow(getDepositCountByRootSQL, root, network).Scan(&depositCount)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, gerror.ErrStorageNotFound
	}
	return depositCount, nil
}

// CheckIfRootExists checks that the root exists on the db.
func (st *DBStorage) CheckIfRootExists(ctx context.Context, root []byte, network uint32, dbTx *sql.Tx) (bool, error) {
	var count uint
	const getDepositCountByRootSQL = "SELECT count(*) FROM mt.root WHERE root = $1 AND network = $2"
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
func (st *DBStorage) GetRoot(ctx context.Context, depositCnt, network uint32, dbTx *sql.Tx) ([]byte, error) {
	var root []byte
	const getRootByDepositCntSQL = "SELECT root FROM mt.root inner join sync.deposit on mt.root.deposit_id = sync.deposit.id WHERE sync.deposit.deposit_cnt = $1 AND network = $2"
	err := st.getExecQuerier(dbTx).QueryRow(getRootByDepositCntSQL, depositCnt, network).Scan(&root)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, gerror.ErrStorageNotFound
	}
	return root, err
}

// SetRoot store the root with deposit count to the storage.
func (st *DBStorage) SetRoot(ctx context.Context, root []byte, depositID uint64, network uint32, dbTx *sql.Tx) error {
	const setRootSQL = "INSERT INTO mt.root (root, deposit_id, network) VALUES ($1, $2, $3);"
	_, err := st.getExecQuerier(dbTx).Exec(setRootSQL, root, depositID, network)
	return err
}

// Get gets value of key from the merkle tree.
func (st *DBStorage) Get(ctx context.Context, key []byte, dbTx *sql.Tx) ([][]byte, error) {
	const getValueByKeySQL = "SELECT value FROM mt.rht WHERE key = $1"
	var data [][]byte
	err := st.getExecQuerier(dbTx).QueryRow(getValueByKeySQL, key).Scan(pq.Array(&data))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, gerror.ErrStorageNotFound
	}
	return data, err
}

// Set inserts a key-value pair into the db.
// If record with such a key already exists its assumed that the value is correct,
// because it's a reverse hash table, and the key is a hash of the value
func (st *DBStorage) Set(ctx context.Context, key []byte, value [][]byte, depositID uint64, dbTx *sql.Tx) error {
	const setNodeSQL = "INSERT INTO mt.rht (deposit_id, key, value) VALUES ($1, $2, $3)"
	_, err := st.getExecQuerier(dbTx).Exec(setNodeSQL, depositID, key, pq.Array(value))
	return err
}

// BulkSet is similar to Set, but it inserts multiple key-value pairs into the db.
func (st *DBStorage) BulkSet(ctx context.Context, rows [][]interface{}, dbTx *sql.Tx) error {
	_, err := copyFrom("mt.rht", []string{"key", "value", "deposit_id"}, rows, dbTx)
	return err
}

// AddRollupExitLeaves iinserts multiple entries into the db.
func (st *DBStorage) AddRollupExitLeaves(ctx context.Context, rows [][]interface{}, dbTx *sql.Tx) error {
	_, err := copyFrom("mt.rollup_exit", []string{"leaf", "rollup_id", "root", "block_id"}, rows, dbTx)
	return err
}

func copyFrom(tableName string, columnNames []string, rows [][]interface{}, dbTx *sql.Tx) (int64, error) {
	if len(rows) == 0 {
		return 0, nil // No rows to insert
	}

	// Build the SQL statement
	placeholders := make([]string, len(rows))
	for i := range rows {
		placeholders[i] = fmt.Sprintf("($%d, $%d, $%d)", i*len(columnNames)+1, i*len(columnNames)+2, i*len(columnNames)+3)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s", tableName, strings.Join(columnNames, ", "), strings.Join(placeholders, ", "))

	// Flatten the data for the Exec call
	args := make([]interface{}, 0, len(rows)*len(columnNames))
	for _, row := range rows {
		args = append(args, row...)
	}

	// Execute the insert statement
	result, err := dbTx.Exec(query, args...)
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
func (st *DBStorage) GetRollupExitLeavesByRoot(ctx context.Context, root common.Hash, dbTx *sql.Tx) ([]etherman.RollupExitLeaf, error) {
	const getLeavesSQL = "SELECT id, leaf, rollup_id, root, block_id FROM mt.rollup_exit WHERE root = $1 ORDER BY rollup_id ASC"
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
func (st *DBStorage) IsRollupExitRoot(ctx context.Context, root common.Hash, dbTx *sql.Tx) (bool, error) {
	const getLeavesSQL = "SELECT count(*) FROM mt.rollup_exit WHERE root = $1"
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
func (st *DBStorage) GetLatestRollupExitLeaves(ctx context.Context, dbTx *sql.Tx) ([]etherman.RollupExitLeaf, error) {
	const getLeavesSQL = `SELECT distinct re.id, re.leaf, re.rollup_id, re.root, re.block_id
		FROM mt.rollup_exit re
		INNER JOIN
			(SELECT distinct rollup_id, MAX(id) AS maxid
			FROM mt.rollup_exit
			GROUP BY rollup_id) groupedre
		ON re.id = groupedre.maxid
		ORDER BY rollup_id asc;
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
func (st *DBStorage) GetLastDepositCount(ctx context.Context, networkID uint32, dbTx *sql.Tx) (uint32, error) {
	var depositCnt int64
	const getLastDepositCountSQL = "SELECT coalesce(MAX(deposit_cnt), -1) FROM sync.deposit WHERE id = (SELECT coalesce(MAX(deposit_id), -1) FROM mt.root WHERE network = $1)"
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
func (st *DBStorage) GetClaimCount(ctx context.Context, destAddr string, dbTx *sql.Tx) (uint64, error) {
	const getClaimCountSQL = "SELECT COUNT(*) FROM sync.claim WHERE dest_addr = $1"
	var claimCount uint64
	err := st.getExecQuerier(dbTx).QueryRow(getClaimCountSQL, common.FromHex(destAddr)).Scan(&claimCount)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, gerror.ErrStorageNotFound
	}
	return claimCount, err
}

// GetClaims gets the claim list which be smaller than index.
func (st *DBStorage) GetClaims(ctx context.Context, destAddr string, limit, offset uint32, dbTx *sql.Tx) ([]*etherman.Claim, error) {
	const getClaimsSQL = "SELECT index, orig_net, orig_addr, amount, dest_addr, block_id, network_id, tx_hash, rollup_index, mainnet_flag FROM sync.claim WHERE dest_addr = $1 ORDER BY block_id DESC LIMIT $2 OFFSET $3"
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
func (st *DBStorage) GetDeposits(ctx context.Context, destAddr string, limit, offset uint32, dbTx *sql.Tx) ([]*etherman.Deposit, error) {
	const getDepositsSQL = "SELECT d.id, leaf_type, orig_net, orig_addr, amount, dest_net, dest_addr, deposit_cnt, block_id, b.block_num, d.network_id, tx_hash, metadata, ready_for_claim FROM sync.deposit as d INNER JOIN sync.block as b ON d.network_id = b.network_id AND d.block_id = b.id WHERE dest_addr = $1 ORDER BY d.block_id DESC, d.deposit_cnt DESC LIMIT $2 OFFSET $3"
	rows, err := st.getExecQuerier(dbTx).Query(getDepositsSQL, common.FromHex(destAddr), limit, offset)
	if err != nil {
		return nil, err
	}

	return parseDeposits(rows, true)
}

// GetDepositCount gets the deposit count for the destination address.
func (st *DBStorage) GetDepositCount(ctx context.Context, destAddr string, dbTx *sql.Tx) (uint64, error) {
	const getDepositCountSQL = "SELECT COUNT(*) FROM sync.deposit WHERE dest_addr = $1"
	var depositCount uint64
	err := st.getExecQuerier(dbTx).QueryRow(getDepositCountSQL, common.FromHex(destAddr)).Scan(&depositCount)
	return depositCount, err
}

// UpdateL1DepositsStatus updates the ready_for_claim status of L1 deposits.
func (st *DBStorage) UpdateL1DepositsStatus(ctx context.Context, exitRoot []byte, destinationNetwork uint32, dbTx *sql.Tx) ([]*etherman.Deposit, error) {
	const updateDepositsStatusSQL = `UPDATE sync.deposit SET ready_for_claim = true 
		WHERE deposit_cnt <=
			(SELECT sync.deposit.deposit_cnt FROM mt.root INNER JOIN sync.deposit ON sync.deposit.id = mt.root.deposit_id WHERE mt.root.root = $1 AND mt.root.network = 0) 
			AND network_id = 0 AND ready_for_claim = false AND dest_net = $2
			RETURNING id, leaf_type, orig_net, orig_addr, amount, dest_net, dest_addr, deposit_cnt, block_id, network_id, tx_hash, metadata, ready_for_claim;`
	rows, err := st.getExecQuerier(dbTx).Query(updateDepositsStatusSQL, exitRoot, destinationNetwork)
	if err != nil {
		return nil, err
	}

	return parseDeposits(rows, false)
}

// UpdateL2DepositsStatus updates the ready_for_claim status of L2 deposits.
func (st *DBStorage) UpdateL2DepositsStatus(ctx context.Context, exitRoot []byte, rollupID, networkID uint32, dbTx *sql.Tx) error {
	const updateL2DepositsStatusSQL = `UPDATE sync.deposit SET ready_for_claim = true
		WHERE deposit_cnt <=
		(SELECT sync.deposit.deposit_cnt FROM mt.root INNER JOIN sync.deposit ON sync.deposit.id = mt.root.deposit_id WHERE mt.root.root = (select leaf from mt.rollup_exit where root = $1 and rollup_id = $2) AND mt.root.network = $3)
			AND network_id = $3 AND ready_for_claim = false;`
	_, err := st.getExecQuerier(dbTx).Exec(updateL2DepositsStatusSQL, exitRoot, rollupID, networkID)
	return err
}

// GetDepositsFromOtherL2ToClaim returns L2 deposits whose destination is an specific L2
func (st *DBStorage) GetDepositsFromOtherL2ToClaim(ctx context.Context, destinationNetwork uint32, dbTx *sql.Tx) ([]*etherman.Deposit, error) {
	const getL2DepositsToClaimStatusSQL = `select sync.deposit.id, sync.deposit.leaf_type, sync.deposit.orig_net, sync.deposit.orig_addr, sync.deposit.amount, sync.deposit.dest_net, sync.deposit.dest_addr, sync.deposit.deposit_cnt, sync.deposit.block_id, sync.deposit.network_id, sync.deposit.tx_hash, sync.deposit.metadata, sync.deposit.ready_for_claim FROM sync.deposit where sync.deposit.deposit_cnt not in (select index FROM sync.claim where sync.claim.network_id = $1) and sync.deposit.network_id !=0 and sync.deposit.dest_net = $1 and ready_for_claim =true order by sync.deposit.id desc;`
	rows, err := st.getExecQuerier(dbTx).Query(getL2DepositsToClaimStatusSQL, destinationNetwork)
	if err != nil {
		return nil, err
	}
	return parseDeposits(rows, false)
}

// GetLatestTrustedGERByDeposit return the latest trusted ger for an specific deposit
func (st *DBStorage) GetLatestTrustedGERByDeposit(ctx context.Context, depositCnt, networkID, destinationNetwork uint32, dbTx *sql.Tx) (common.Hash, error) {
	const getLatestTrustedGERByDeposit = `SELECT sync.exit_root.global_exit_root FROM sync.deposit inner join mt.root on mt.root.deposit_id = sync.deposit.id inner join mt.rollup_exit on mt.rollup_exit.leaf = mt.root.root inner join sync.exit_root on sync.exit_root.exit_roots[2]= mt.rollup_exit.root where sync.exit_root.allowed = true and deposit_cnt = $1 and sync.deposit.network_id = $2 and dest_net = $3 and mt.rollup_exit.rollup_id = $2 and sync.exit_root.network_id = sync.deposit.dest_net order by sync.exit_root.id desc limit 1`
	var ger common.Hash
	err := st.getExecQuerier(dbTx).QueryRow(getLatestTrustedGERByDeposit, depositCnt, networkID, destinationNetwork).Scan(&ger)
	if errors.Is(err, sql.ErrNoRows) {
		return common.Hash{}, gerror.ErrStorageNotFound
	}
	return ger, err
}

// AddClaimTx adds a claim monitored transaction to the storage.
func (st *DBStorage) AddClaimTx(ctx context.Context, mTx ctmtypes.MonitoredTx, dbTx *sql.Tx) error {
	const addMonitoredTxSQL = `INSERT INTO sync.monitored_txs 
		(deposit_id, from_addr, to_addr, nonce, value, data, gas, status, history, created_at, updated_at, group_id, global_exit_root)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	_, err := st.getExecQuerier(dbTx).Exec(addMonitoredTxSQL, mTx.DepositID, mTx.From, mTx.To, mTx.Nonce, mTx.Value.String(),
		mTx.Data, mTx.Gas, mTx.Status, pq.Array(mTx.HistoryHashSlice()), time.Now().UTC(), time.Now().UTC(), mTx.GroupID, mTx.GlobalExitRoot)
	return err
}

// UpdateClaimTx updates a claim monitored transaction in the storage.
func (st *DBStorage) UpdateClaimTx(ctx context.Context, mTx ctmtypes.MonitoredTx, dbTx *sql.Tx) error {
	const updateMonitoredTxSQL = `UPDATE sync.monitored_txs 
		SET from_addr = $2
		, to_addr = $3
		, nonce = $4
		, value = $5
		, data = $6
		, gas = $7
		, status = $8
		, history = $9
		, updated_at = $10
		, group_id = $11
		WHERE deposit_id = $1`
	_, err := st.getExecQuerier(dbTx).Exec(updateMonitoredTxSQL, mTx.DepositID, mTx.From, mTx.To, mTx.Nonce, mTx.Value.String(),
		mTx.Data, mTx.Gas, mTx.Status, pq.Array(mTx.HistoryHashSlice()), time.Now().UTC(), mTx.GroupID)
	return err
}

// GetClaimTxsByStatus gets the monitored transactions by status.
func (st *DBStorage) GetClaimTxsByStatus(ctx context.Context, statuses []ctmtypes.MonitoredTxStatus, rollupID uint32, dbTx *sql.Tx) ([]ctmtypes.MonitoredTx, error) {
	const getMonitoredTxsSQL = "SELECT deposit_id, from_addr, to_addr, nonce, value, data, gas, status, history, created_at, updated_at, group_id, global_exit_root FROM sync.monitored_txs INNER JOIN sync.deposit ON sync.deposit.id = sync.monitored_txs.deposit_id WHERE status = ANY($1) AND sync.deposit.dest_net = $2 ORDER BY created_at ASC"
	rows, err := st.getExecQuerier(dbTx).Query(getMonitoredTxsSQL, pq.Array(statuses), rollupID)
	if errors.Is(err, sql.ErrNoRows) {
		return []ctmtypes.MonitoredTx{}, nil
	} else if err != nil {
		return nil, err
	}

	var mTxs []ctmtypes.MonitoredTx
	for rows.Next() {
		var (
			value   string
			history [][]byte
		)
		mTx := ctmtypes.MonitoredTx{}
		err = rows.Scan(&mTx.DepositID, &mTx.From, &mTx.To, &mTx.Nonce, &value, &mTx.Data, &mTx.Gas, &mTx.Status, pq.Array(&history), &mTx.CreatedAt, &mTx.UpdatedAt, &mTx.GroupID, &mTx.GlobalExitRoot)
		if err != nil {
			return mTxs, err
		}
		mTx.Value, _ = new(big.Int).SetString(value, 10) //nolint:mnd
		mTx.History = make(map[common.Hash]bool)
		for _, h := range history {
			mTx.History[common.BytesToHash(h)] = true
		}
		mTxs = append(mTxs, mTx)
	}

	return mTxs, nil
}

// GetPendingDepositsToClaim gets the deposit list which is not claimed in the destination network.
func (st *DBStorage) GetPendingDepositsToClaim(ctx context.Context, destAddress common.Address, destNetwork, leafType, limit, offset uint32, dbTx *sql.Tx) ([]*etherman.Deposit, uint64, error) {
	desAddrSQL := ""
	if destAddress != (common.Address{}) {
		str := strings.TrimPrefix(destAddress.String(), "0x")
		desAddrSQL = "AND dest_addr = decode('" + str + "','hex')"
	}
	getNumberPendingDepositsToClaimSQL := "SELECT count(*) FROM sync.deposit WHERE dest_net = $1 AND ready_for_claim = true AND leaf_type = $2 " + desAddrSQL + " AND deposit_cnt NOT IN (SELECT index FROM sync.claim WHERE sync.claim.network_id = $1)"
	var totalCount uint64
	err := st.getExecQuerier(dbTx).QueryRow(getNumberPendingDepositsToClaimSQL, destNetwork, leafType).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}
	getPendingDepositsToClaimSQL := "SELECT d.id, leaf_type, orig_net, orig_addr, amount, dest_net, dest_addr, deposit_cnt, block_id, b.block_num, d.network_id, tx_hash, metadata, ready_for_claim FROM sync.deposit AS d INNER JOIN sync.block AS b ON d.block_id = b.id WHERE dest_net = $1 AND ready_for_claim = true AND leaf_type = $2 " + desAddrSQL + " AND d.deposit_cnt NOT IN (SELECT index FROM sync.claim WHERE sync.claim.network_id = $1) ORDER BY d.deposit_cnt ASC LIMIT $3 OFFSET $4"
	rows, err := st.getExecQuerier(dbTx).Query(getPendingDepositsToClaimSQL, destNetwork, leafType, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	deposits, err := parseDeposits(rows, true)
	if err != nil {
		return nil, 0, err
	}
	return deposits, totalCount, nil
}

func (st *DBStorage) AddRemoveL2GER(ctx context.Context, globalExitRoot etherman.GlobalExitRoot, dbTx *sql.Tx) error {
	const insertRemoveGERSQL = `INSERT INTO sync.remove_exit_root 
		(block_id, global_exit_root, network_id)
		VALUES ($1, $2, $3)`
	_, err := st.getExecQuerier(dbTx).Exec(insertRemoveGERSQL, globalExitRoot.BlockID, globalExitRoot.GlobalExitRoot, globalExitRoot.NetworkID)
	if err != nil {
		return err
	}

	// Modify the allowed column in the exit_root table for this globalExitRoot
	const updateGERStatusSQL = "UPDATE sync.exit_root SET allowed = false WHERE global_exit_root = $1 AND network_id = $2 RETURNING id;"
	var gerID uint64
	err = st.getExecQuerier(dbTx).QueryRow(updateGERStatusSQL, globalExitRoot.GlobalExitRoot, globalExitRoot.NetworkID).Scan(&gerID)
	if err != nil {
		return err
	}

	// Look for the deposits where the ready for claim flag needs to be modified
	getPreviousGERSQL := "SELECT global_exit_root FROM sync.exit_root WHERE network_id = $1 AND id < $2 ORDER BY id DESC LIMIT 1"
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
	const updateDepositsFlagSQL = `UPDATE sync.deposit SET ready_for_claim = false WHERE id > $1 AND id <= $2 AND network_id = 0 AND dest_net = $3`
	_, err = st.getExecQuerier(dbTx).Exec(updateDepositsFlagSQL, prevDepositID, CurrentDepositID, globalExitRoot.NetworkID)
	return err
}

// GetDepositCountByGER gets the deposit count by the GER.
func (st *DBStorage) GetDepositCountByGER(ctx context.Context, ger common.Hash, network uint32, rollupsTree bool, dbTx *sql.Tx) (uint64, uint32, error) {
	var (
		depositID               uint64
		depositCnt, rootNetwork uint32
	)
	arrayIndex := 1
	if rollupsTree {
		arrayIndex = 2
		rootNetwork = network
	}

	const getDepositCountByGERSQL = "SELECT sync.deposit.id, sync.deposit.deposit_cnt FROM sync.deposit INNER JOIN mt.root ON sync.deposit.id = mt.root.deposit_id inner join sync.exit_root on mt.root.root = sync.exit_root.exit_roots[$1] WHERE sync.exit_root.global_exit_root = $2 and sync.exit_root.network_id = $3 AND mt.root.network = $4"
	err := st.getExecQuerier(dbTx).QueryRow(getDepositCountByGERSQL, arrayIndex, ger, network, rootNetwork).Scan(&depositID, &depositCnt)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, 0, gerror.ErrStorageNotFound
	}
	return depositID, depositCnt, err
}

// UpdateDepositsStatusForTesting updates the ready_for_claim status of all deposits for testing.
func (st *DBStorage) UpdateDepositsStatusForTesting(ctx context.Context, dbTx *sql.Tx) error {
	const updateDepositsStatusSQL = "UPDATE sync.deposit SET ready_for_claim = true;"
	_, err := st.getExecQuerier(dbTx).Exec(updateDepositsStatusSQL)
	return err
}

// UpdateBlocksForTesting updates the hash of blocks.
func (st *DBStorage) UpdateBlocksForTesting(ctx context.Context, networkID uint32, blockNum uint64, dbTx *sql.Tx) error {
	const updateBlocksSQL = "UPDATE sync.block SET block_hash = SUBSTRING(block_hash FROM 1 FOR LENGTH(block_hash)-1) || '\x61' WHERE network_id = $1 AND block_num >= $2"
	_, err := st.getExecQuerier(dbTx).Exec(updateBlocksSQL, networkID, blockNum)
	return err
}

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
