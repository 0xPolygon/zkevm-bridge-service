-- +migrate Up

CREATE TABLE IF NOT EXISTS sync.forward_let(
    id                   BIGSERIAL,
    block_id             BIGINT REFERENCES sync.block (id) ON DELETE CASCADE,
    previous_deposit_cnt BIGINT,
    previous_root        BYTEA,
    new_deposit_cnt      BIGINT,
    new_root             BYTEA,
    new_raw_leaves       BYTEA,
    PRIMARY KEY (id)
);

-- +migrate Down

DROP TABLE IF EXISTS sync.forward_let;
