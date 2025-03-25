package sqlitestorage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	ctmtypes "github.com/0xPolygonHermez/zkevm-bridge-service/claimtxman/types"
	_ "github.com/mattn/go-sqlite3"
)

// GetLatestGroupID
func (st *DBStorage) GetLatestMonitoredTxGroupID(ctx context.Context, dbTx interface{}) (uint64, error) {
	const selectQuery = `SELECT group_id FROM monitored_txs_group ORDER BY group_id DESC LIMIT 1`
	var groupID uint64
	err := st.getExecQuerier(dbTx).QueryRow(selectQuery).Scan(&groupID)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return groupID, nil
}

// AddMonitoredTxsGroup
func (st *DBStorage) AddMonitoredTxsGroup(ctx context.Context, mTxGroup *ctmtypes.MonitoredTxGroupDBEntry, dbTx interface{}) error {
	const insertQuery = `INSERT INTO monitored_txs_group 
		(group_id,status, deposit_ids, num_retries, compressed_tx_data,claim_tx_history, created_at, updated_at) 
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)
		`
	if mTxGroup == nil {
		return errors.New("nil monitored tx group")
	}
	claimTxHistoryStr, err := mTxGroup.ClaimTxHistory.ToJson()
	if err != nil {
		return fmt.Errorf("fails to convert claimTxHistory to json. Err: %w", err)
	}
	_, err = st.getExecQuerier(dbTx).Exec(
		insertQuery,
		mTxGroup.GroupID,
		mTxGroup.Status,
		convertArrayUint64ToString(mTxGroup.DepositIDs),
		mTxGroup.NumRetries,
		mTxGroup.CompressedTxData,
		claimTxHistoryStr,
		mTxGroup.CreatedAt,
		mTxGroup.UpdatedAt)

	if err != nil {
		return err
	}

	return err
}

func convertArrayUint64ToString(array []uint64) string{
	var numbers string
	for i, number := range array {
		if i != 0 {
			numbers += ","
		}
		numbers += strconv.FormatUint(number, 10)
	}
	return numbers
}

func convertStringToArrayUint64(numbers string) ([]uint64, error) {
	var array []uint64
	numbString := strings.Split(numbers, ",")
	for _, s := range numbString {
		number, err := strconv.ParseUint(s, 10, 64) // Convert string to uint64
		if err != nil {
			return []uint64{}, err
		}
		array = append(array, number)
	}
	return array, nil
}


func (st *DBStorage) UpdateMonitoredTxsGroup(ctx context.Context, mTxGroup *ctmtypes.MonitoredTxGroupDBEntry, dbTx interface{}) error {
	const updateQuery = `UPDATE monitored_txs_group 
		SET num_retries = ?, compressed_tx_data = ?, claim_tx_history = ?, updated_at = ?, status = ?, last_log = ?
		WHERE group_id = ?
		`
	if mTxGroup == nil {
		return errors.New("nil monitored tx group")
	}
	claimTxHistoryStr, err := mTxGroup.ClaimTxHistory.ToJson()
	if err != nil {
		return fmt.Errorf("fails to convert claimTxHistory to json. Err: %w", err)
	}
	_, err = st.getExecQuerier(dbTx).Exec(
		updateQuery,
		mTxGroup.GroupID,
		mTxGroup.NumRetries,
		mTxGroup.CompressedTxData,
		claimTxHistoryStr,
		mTxGroup.UpdatedAt,
		mTxGroup.Status,
		mTxGroup.LastLog)

	if err != nil {
		return err
	}

	return err
}

func (st *DBStorage) GetMonitoredTxsGroups(ctx context.Context, groupIds []uint64, dbTx interface{}) (map[uint64]ctmtypes.MonitoredTxGroupDBEntry, error) {
	selectQuery := fmt.Sprintf("SELECT group_id, status, num_retries, compressed_tx_data,claim_tx_history, created_at, updated_at FROM monitored_txs_group WHERE group_id =IN(%s) ORDER BY created_at ASC", convertArrayUint64ToString(groupIds))
	groups := make(map[uint64]ctmtypes.MonitoredTxGroupDBEntry)
	rows, err := st.getExecQuerier(dbTx).Query(selectQuery)
	if errors.Is(err, sql.ErrNoRows) {
		return groups, nil
	} else if err != nil {
		return groups, err
	}

	for rows.Next() {
		var group ctmtypes.MonitoredTxGroupDBEntry
		var claimTxHistoryStr string
		err = rows.Scan(&group.GroupID, &group.Status, &group.NumRetries, &group.CompressedTxData, &claimTxHistoryStr, &group.CreatedAt, &group.UpdatedAt)
		if err != nil {
			return nil, err
		}
		group.ClaimTxHistory, err = ctmtypes.NewTxHistoryV2FromJson(claimTxHistoryStr)
		if err != nil {
			return nil, fmt.Errorf("fails to convert claimTxHistory from json. Err: %w", err)
		}
		groups[group.GroupID] = group
	}

	return groups, nil
}
