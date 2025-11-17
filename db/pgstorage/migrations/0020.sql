-- +migrate Up

CREATE TABLE IF NOT EXISTS sync.backward_let(
    id                   BIGSERIAL,
    block_id             BIGINT REFERENCES sync.block (id) ON DELETE CASCADE,
    previous_deposit_cnt BIGINT,
    previous_root        BYTEA,
    new_deposit_cnt      BIGINT,
    new_root             BYTEA,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS sync.set_unset_claim(
    id               BIGSERIAL,
    block_id         BIGINT REFERENCES sync.block (id) ON DELETE CASCADE,
    mainnet_flag     BOOLEAN,
    rollup_index     BIGINT,
    index            BIGINT,
    global_index     VARCHAR(255),
    type             VARCHAR(255), -- "SET" or "UNSET"
    PRIMARY KEY (id)
);

-- +migrate Down

DROP TABLE IF EXISTS sync.backward_let;
DROP TABLE IF EXISTS sync.set_unset_claim;


