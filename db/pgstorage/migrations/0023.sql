-- +migrate Up
-- PostgreSQL Index Optimization Migration
-- Purpose: Add missing indexes to improve query performance for zkEVM bridge service
-- Focus areas: Deposits, Claims, Blocks, Exit Roots, Merkle Tree operations

-- ============================================================================
-- SECTION 1: sync.deposit table indexes
-- ============================================================================

-- Index for GetNumberDeposits query: WHERE d.network_id = $1 AND b.block_num <= $2
-- Composite index to support both filtering and joining operations
CREATE INDEX IF NOT EXISTS idx_deposit_network_id_block_id
ON sync.deposit(network_id, block_id);

-- Index for deposit lookup by network and counter: WHERE network_id = $1 AND deposit_cnt = $2
-- Already has a composite filter pattern, adds coverage for frequent lookups
CREATE INDEX IF NOT EXISTS idx_deposit_network_id_deposit_cnt
ON sync.deposit(network_id, deposit_cnt);

-- Index for destination address queries: WHERE dest_addr = $1
-- Used in GetDeposits, GetDepositCount, GetPendingDepositsToClaim
CREATE INDEX IF NOT EXISTS idx_deposit_dest_addr
ON sync.deposit(dest_addr);

-- Index for pending deposits query: WHERE dest_net = $1 AND ready_for_claim = true AND ignore = false
-- Supports GetPendingDepositsToClaim with multiple conditions
CREATE INDEX IF NOT EXISTS idx_deposit_dest_net_ready_ignore
ON sync.deposit(dest_net, ready_for_claim, ignore);

-- Index for deposits ready to claim: WHERE ready_for_claim = true AND dest_net = $1
-- Supports UpdateL1DepositsStatus and UpdateL2DepositsStatus queries
CREATE INDEX IF NOT EXISTS idx_deposit_ready_dest_net
ON sync.deposit(ready_for_claim, dest_net) WHERE ready_for_claim = true;

-- Index for deposit metadata lookups: WHERE network_id = $1 AND orig_addr = $2 AND dest_net = $3
-- Used in GetTokenMetadata
CREATE INDEX IF NOT EXISTS idx_deposit_metadata_lookup
ON sync.deposit(network_id, orig_addr, dest_net) WHERE metadata IS NOT NULL;

-- Index for GetDepositsFromOtherL2ToClaim complex query
-- WHERE network_id != 0 AND ignore = false AND dest_net = $1 AND ready_for_claim = true
CREATE INDEX IF NOT EXISTS idx_deposit_l2_claim_status
ON sync.deposit(dest_net, ready_for_claim, ignore, network_id)
WHERE network_id != 0 AND ignore = false AND ready_for_claim = true;

-- Index for range queries on deposit IDs used in AddRemoveL2GER
-- WHERE id > $1 AND id <= $2 AND network_id = 0 AND dest_net = $3
CREATE INDEX IF NOT EXISTS idx_deposit_id_network_dest
ON sync.deposit(id, network_id, dest_net);

-- Index for original address lookups (used in GetPendingDepositsToClaim with leafType filter)
CREATE INDEX IF NOT EXISTS idx_deposit_orig_addr
ON sync.deposit(orig_addr);

-- ============================================================================
-- SECTION 2: sync.claim table indexes
-- ============================================================================

-- Index for mainnet flag filtering: WHERE mainnet_flag = true AND network_id = $1
CREATE INDEX IF NOT EXISTS idx_claim_mainnet_network
ON sync.claim(mainnet_flag, network_id) WHERE mainnet_flag = true;

-- Index for rollup claim filtering: WHERE NOT mainnet_flag AND rollup_index + 1 = $1 AND network_id = $2
CREATE INDEX IF NOT EXISTS idx_claim_rollup_network
ON sync.claim(rollup_index, network_id) WHERE NOT mainnet_flag;

-- Index for destination address queries: WHERE dest_addr = $1
-- Used in GetClaimCount and GetClaims
CREATE INDEX IF NOT EXISTS idx_claim_dest_addr
ON sync.claim(dest_addr);

-- Index for GetDepositsFromOtherL2ToClaim subquery: WHERE network_id = $1 AND index = ?
-- Supports NOT IN subquery optimization
CREATE INDEX IF NOT EXISTS idx_claim_network_id_index
ON sync.claim(network_id, index);

-- Index for global_index lookups used in DeleteClaimByGlobalIndex
CREATE INDEX IF NOT EXISTS idx_claim_global_index_network
ON sync.claim(global_index, network_id);

-- ============================================================================
-- SECTION 3: sync.block table indexes
-- ============================================================================

-- Index for GetLastBlock and GetPreviousBlock queries
-- WHERE network_id = $1 ORDER BY block_num DESC
CREATE INDEX IF NOT EXISTS idx_block_network_id_block_num_desc
ON sync.block(network_id, block_num DESC);

-- Index for block hash lookups: WHERE block_hash = $1
-- Already covered by UNIQUE constraint block_hash_unique from migration 0004
-- UNIQUE constraints automatically create an index in PostgreSQL

-- Index for block number range queries used in Reset
-- WHERE block_num > $1 AND network_id = $2
CREATE INDEX IF NOT EXISTS idx_block_num_network_id
ON sync.block(block_num, network_id);

-- ============================================================================
-- SECTION 4: sync.exit_root table indexes
-- ============================================================================

-- Index for GetLatestL1SyncedExitRoot query
-- WHERE allowed = true AND block_id > 0 AND network_id = 0 ORDER BY id DESC
CREATE INDEX IF NOT EXISTS idx_exit_root_allowed_block_id_network
ON sync.exit_root(allowed, block_id, network_id, id DESC)
WHERE allowed = true AND block_id > 0;

-- Index for GetLatestTrustedExitRoot query
-- WHERE network_id = $1 AND allowed = true ORDER BY id DESC
CREATE INDEX IF NOT EXISTS idx_exit_root_network_allowed_id_desc
ON sync.exit_root(network_id, allowed, id DESC) WHERE allowed = true;

-- Index for GetL1ExitRootByGER query
-- WHERE allowed = true AND block_id > 0 AND global_exit_root = $1 AND network_id = 0
CREATE INDEX IF NOT EXISTS idx_exit_root_ger_lookup
ON sync.exit_root(global_exit_root, network_id, allowed, block_id)
WHERE allowed = true AND block_id > 0;

-- Index for GetL2ExitRootsByGER query
-- WHERE allowed = true AND block_id > 0 AND global_exit_root = $1 AND network_id != 0
CREATE INDEX IF NOT EXISTS idx_exit_root_l2_ger
ON sync.exit_root(global_exit_root, network_id, allowed, block_id)
WHERE allowed = true AND block_id > 0 AND network_id != 0;

-- Index for exit_roots array queries (used in GetDepositCountByGER)
-- Using GIN index for array containment operations
CREATE INDEX IF NOT EXISTS idx_exit_root_exit_roots_array
ON sync.exit_root USING GIN(exit_roots);

-- Index for UpdateL2GER: WHERE id = $1
-- Already has primary key, but adding for clarity in WHERE clauses

-- Index for GetPreviousGERSQL in AddRemoveL2GER
-- WHERE network_id = $1 AND id < $2 ORDER BY id DESC
CREATE INDEX IF NOT EXISTS idx_exit_root_network_id_desc
ON sync.exit_root(network_id, id DESC);

-- Index for UpdateGERStatusSQL in AddRemoveL2GER
-- WHERE global_exit_root = $1 AND network_id = $2
CREATE INDEX IF NOT EXISTS idx_exit_root_ger_network
ON sync.exit_root(global_exit_root, network_id);

-- ============================================================================
-- SECTION 5: sync.token_wrapped table indexes
-- ============================================================================

-- Index for GetTokenWrapped query: WHERE orig_net = $1 AND orig_token_addr = $2
-- Note: Primary key is (network_id, orig_net, orig_token_addr) per migration 0006
-- This index complements the primary key for queries that filter by orig_net and
-- orig_token_addr without network_id, which the PK cannot efficiently support
CREATE INDEX IF NOT EXISTS idx_token_wrapped_orig_net_addr
ON sync.token_wrapped(orig_net, orig_token_addr);

-- ============================================================================
-- SECTION 6: mt.root table indexes (Merkle Tree)
-- ============================================================================

-- Index for GetDepositCountByRoot: WHERE root = $1 AND network = $2
-- Migration 0002 already created root_network_idx: ON mt.root(root, network)
-- Verified this index exists

-- Index for GetRoot: WHERE deposit_cnt = $1 AND network = $2
-- The query logic maps deposit_cnt to deposit_id via sync.deposit JOIN
-- This index supports lookups by deposit_id and network
CREATE INDEX IF NOT EXISTS idx_root_deposit_id_network
ON mt.root(deposit_id, network);

-- Index for CheckIfRootExists: WHERE root = $1 AND network = $2
-- Already covered by root_network_idx from migration 0002

-- Index for GetLastDepositCount subquery: WHERE network = $1
CREATE INDEX IF NOT EXISTS idx_root_network
ON mt.root(network);

-- ============================================================================
-- SECTION 7: mt.rht table indexes (Reverse Hash Table)
-- ============================================================================

-- Index for Get operation: WHERE key = $1
-- Migration 0003 already created rht_key_idx on (key)
-- The table has a deposit_id column as foreign key to sync.deposit(id)
-- Verified this index exists

-- ============================================================================
-- SECTION 8: mt.rollup_exit table indexes
-- ============================================================================

-- Index for GetRollupExitLeavesByRoot: WHERE root = $1 ORDER BY rollup_id ASC
CREATE INDEX IF NOT EXISTS idx_rollup_exit_root_rollup_id
ON mt.rollup_exit(root, rollup_id);

-- Index for IsRollupExitRoot: WHERE root = $1
CREATE INDEX IF NOT EXISTS idx_rollup_exit_root
ON mt.rollup_exit(root);

-- Index for GetLatestRollupExitLeaves: GROUP BY rollup_id, MAX(id)
-- Index to support grouping and aggregation
CREATE INDEX IF NOT EXISTS idx_rollup_exit_rollup_id_id
ON mt.rollup_exit(rollup_id, id DESC);

-- Index for UpdateL2DepositsStatus subquery: WHERE root = $1 AND rollup_id = $2
-- Already covered by idx_rollup_exit_root_rollup_id above

-- ============================================================================
-- SECTION 9: sync.monitored_txs table indexes
-- ============================================================================

-- Index for GetClaimTxsByStatus query with JOIN
-- WHERE status = ANY($1) AND sync.deposit.dest_net = $2 ORDER BY created_at ASC
CREATE INDEX IF NOT EXISTS idx_monitored_txs_deposit_id
ON sync.monitored_txs(deposit_id);

-- Index for status filtering: WHERE status = ANY($1)
CREATE INDEX IF NOT EXISTS idx_monitored_txs_status_created
ON sync.monitored_txs(status, created_at);

-- Index for AddClaimTx: ON CONFLICT (deposit_id) DO NOTHING
-- Already has deposit_id as primary key

-- Index for UpdateClaimTx: WHERE deposit_id = $1
-- Already covered by primary key

-- Index for group_id lookups
CREATE INDEX IF NOT EXISTS idx_monitored_txs_group_id
ON sync.monitored_txs(group_id) WHERE group_id IS NOT NULL;

-- ============================================================================
-- SECTION 10: sync.remove_exit_root table indexes
-- ============================================================================

-- Index for potential lookups by global_exit_root and network_id
CREATE INDEX IF NOT EXISTS idx_remove_exit_root_ger_network
ON sync.remove_exit_root(global_exit_root, network_id);

-- Index for block_id references
CREATE INDEX IF NOT EXISTS idx_remove_exit_root_block_id
ON sync.remove_exit_root(block_id);

-- ============================================================================
-- SECTION 11: sync.backward_let table indexes
-- ============================================================================

-- Index for potential lookups by roots and deposit counts
CREATE INDEX IF NOT EXISTS idx_backward_let_new_root
ON sync.backward_let(new_root);

CREATE INDEX IF NOT EXISTS idx_backward_let_previous_root
ON sync.backward_let(previous_root);

-- Index for block_id for reorg operations
CREATE INDEX IF NOT EXISTS idx_backward_let_block_id
ON sync.backward_let(block_id);

-- ============================================================================
-- SECTION 12: sync.forward_let table indexes
-- ============================================================================

-- Index for forward LET lookups by roots
CREATE INDEX IF NOT EXISTS idx_forward_let_new_root
ON sync.forward_let(new_root);

CREATE INDEX IF NOT EXISTS idx_forward_let_previous_root
ON sync.forward_let(previous_root);

-- Index for block_id for reorg operations
CREATE INDEX IF NOT EXISTS idx_forward_let_block_id
ON sync.forward_let(block_id);

-- ============================================================================
-- SECTION 13: sync.set_unset_claim table indexes
-- ============================================================================

-- Index for potential lookups by index, rollup_index, network operations
CREATE INDEX IF NOT EXISTS idx_set_unset_claim_index_rollup
ON sync.set_unset_claim(index, rollup_index, mainnet_flag);

-- Index for global_index lookups
CREATE INDEX IF NOT EXISTS idx_set_unset_claim_global_index
ON sync.set_unset_claim(global_index);

-- Index for type filtering (SET vs UNSET)
CREATE INDEX IF NOT EXISTS idx_set_unset_claim_type
ON sync.set_unset_claim(type);

-- Index for block_id for reorg operations
CREATE INDEX IF NOT EXISTS idx_set_unset_claim_block_id
ON sync.set_unset_claim(block_id);

-- ============================================================================
-- SECTION 14: sync.deposit_backup table indexes
-- ============================================================================

-- Index for deposit_cnt ordering (used in ResetDeposits and GetAndDeleteOrphanDepositBackups)
CREATE INDEX IF NOT EXISTS idx_deposit_backup_deposit_cnt
ON sync.deposit_backup(deposit_cnt);

-- Index for network_id filtering
CREATE INDEX IF NOT EXISTS idx_deposit_backup_network_id
ON sync.deposit_backup(network_id);

-- Note: idx_deposit_backup_backward_let_id already exists from migration 0022

-- ============================================================================
-- SECTION 15: sync.status table indexes
-- ============================================================================

-- Index for GetSyncStatus: ORDER BY network_id ASC
-- Primary key already exists on network_id
-- Verified coverage

-- +migrate Down

-- Remove all indexes created in this migration
DROP INDEX IF EXISTS sync.idx_deposit_network_id_block_id;
DROP INDEX IF EXISTS sync.idx_deposit_network_id_deposit_cnt;
DROP INDEX IF EXISTS sync.idx_deposit_dest_addr;
DROP INDEX IF EXISTS sync.idx_deposit_dest_net_ready_ignore;
DROP INDEX IF EXISTS sync.idx_deposit_ready_dest_net;
DROP INDEX IF EXISTS sync.idx_deposit_metadata_lookup;
DROP INDEX IF EXISTS sync.idx_deposit_l2_claim_status;
DROP INDEX IF EXISTS sync.idx_deposit_id_network_dest;
DROP INDEX IF EXISTS sync.idx_deposit_orig_addr;

DROP INDEX IF EXISTS sync.idx_claim_mainnet_network;
DROP INDEX IF EXISTS sync.idx_claim_rollup_network;
DROP INDEX IF EXISTS sync.idx_claim_dest_addr;
DROP INDEX IF EXISTS sync.idx_claim_network_id_index;
DROP INDEX IF EXISTS sync.idx_claim_global_index_network;

DROP INDEX IF EXISTS sync.idx_block_network_id_block_num_desc;
DROP INDEX IF EXISTS sync.idx_block_num_network_id;

DROP INDEX IF EXISTS sync.idx_exit_root_allowed_block_id_network;
DROP INDEX IF EXISTS sync.idx_exit_root_network_allowed_id_desc;
DROP INDEX IF EXISTS sync.idx_exit_root_ger_lookup;
DROP INDEX IF EXISTS sync.idx_exit_root_l2_ger;
DROP INDEX IF EXISTS sync.idx_exit_root_exit_roots_array;
DROP INDEX IF EXISTS sync.idx_exit_root_network_id_desc;
DROP INDEX IF EXISTS sync.idx_exit_root_ger_network;

DROP INDEX IF EXISTS sync.idx_token_wrapped_orig_net_addr;

DROP INDEX IF EXISTS mt.idx_root_deposit_id_network;
DROP INDEX IF EXISTS mt.idx_root_network;

DROP INDEX IF EXISTS mt.idx_rollup_exit_root_rollup_id;
DROP INDEX IF EXISTS mt.idx_rollup_exit_root;
DROP INDEX IF EXISTS mt.idx_rollup_exit_rollup_id_id;

DROP INDEX IF EXISTS sync.idx_monitored_txs_deposit_id;
DROP INDEX IF EXISTS sync.idx_monitored_txs_status_created;
DROP INDEX IF EXISTS sync.idx_monitored_txs_group_id;

DROP INDEX IF EXISTS sync.idx_remove_exit_root_ger_network;
DROP INDEX IF EXISTS sync.idx_remove_exit_root_block_id;

DROP INDEX IF EXISTS sync.idx_backward_let_new_root;
DROP INDEX IF EXISTS sync.idx_backward_let_previous_root;
DROP INDEX IF EXISTS sync.idx_backward_let_block_id;

DROP INDEX IF EXISTS sync.idx_forward_let_new_root;
DROP INDEX IF EXISTS sync.idx_forward_let_previous_root;
DROP INDEX IF EXISTS sync.idx_forward_let_block_id;

DROP INDEX IF EXISTS sync.idx_set_unset_claim_index_rollup;
DROP INDEX IF EXISTS sync.idx_set_unset_claim_global_index;
DROP INDEX IF EXISTS sync.idx_set_unset_claim_type;
DROP INDEX IF EXISTS sync.idx_set_unset_claim_block_id;

DROP INDEX IF EXISTS sync.idx_deposit_backup_deposit_cnt;
DROP INDEX IF EXISTS sync.idx_deposit_backup_network_id;
