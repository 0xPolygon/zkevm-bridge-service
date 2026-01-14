-- +migrate Up

CREATE TABLE IF NOT EXISTS sync.deposit_backup
(
    leaf_type   INTEGER,
    network_id  BIGINT,
    orig_net    BIGINT,
    orig_addr   BYTEA NOT NULL,
    amount      VARCHAR,
    dest_net    BIGINT NOT NULL,
    dest_addr   BYTEA NOT NULL,
    block_id    BIGINT REFERENCES sync.block (id) ON DELETE CASCADE,
    deposit_cnt BIGINT,
    tx_hash     BYTEA NOT NULL,
    metadata    BYTEA NOT NULL,
    id          BIGSERIAL,
    ready_for_claim BOOLEAN DEFAULT false NOT NULL,
    backward_let_id BIGINT NOT NULL,
    PRIMARY KEY (id)
);

-- Add index on backward_let_id for faster orphan lookups and JOIN operations
CREATE INDEX IF NOT EXISTS idx_deposit_backup_backward_let_id
ON sync.deposit_backup(backward_let_id);

-- +migrate Down

DROP INDEX IF EXISTS sync.idx_deposit_backup_backward_let_id;
DROP TABLE IF EXISTS sync.deposit_backup;