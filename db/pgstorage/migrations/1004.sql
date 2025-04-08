-- +migrate Down

DROP TABLE IF EXISTS sync.bridge_balance;

-- +migrate Up

CREATE TABLE IF NOT EXISTS sync.bridge_balance
(
    id SERIAL PRIMARY KEY,
    original_token_addr BYTEA NOT NULL,
    network_id INTEGER NOT NULL,
    balance VARCHAR NOT NULL DEFAULT '0',
    create_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    modify_time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT bridge_balance_uidx UNIQUE (original_token_addr, network_id)
);
