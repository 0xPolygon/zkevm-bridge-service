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

-- +migrate Down

DROP TABLE IF EXISTS sync.deposit_backup;