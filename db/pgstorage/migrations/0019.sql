-- +migrate Up

CREATE TABLE IF NOT EXISTS sync.status(
    network_id       BIGINT,
    percentage       INT,
    remaining_blocks BIGINT,
    synced           BOOLEAN,
    PRIMARY KEY (network_id)
);

-- +migrate Down

DROP TABLE IF EXISTS sync.status;
