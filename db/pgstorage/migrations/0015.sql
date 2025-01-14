-- +migrate Up

CREATE TABLE IF NOT EXISTS sync.remove_exit_root(
    id                      BIGSERIAL,
    block_id                BIGINT REFERENCES sync.block (id) ON DELETE CASCADE,
    global_exit_root        BYTEA,
    network_id              BIGINT,
    PRIMARY KEY (id)
);

ALTER TABLE sync.exit_root ADD COLUMN IF NOT EXISTS allowed BOOLEAN NOT NULL DEFAULT true;



-- +migrate Down

DROP TABLE IF EXISTS sync.remove_exit_root;

ALTER TABLE sync.exit_root DROP COLUMN IF EXISTS allowed;
