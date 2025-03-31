-- +migrate Up

PRAGMA foreign_keys = ON;
PRAGMA journal_mode = WAL;
PRAGMA busy_timeout = 20000;

CREATE TABLE IF NOT EXISTS rht (
    key BLOB NOT NULL,
    left_node BLOB,
    right_node BLOB,
    deposit_id INTEGER NOT NULL,
    PRIMARY KEY (key, deposit_id),
    FOREIGN KEY (deposit_id) REFERENCES deposit(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS rollup_exit (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    leaf BLOB,
    rollup_id INTEGER,
    root BLOB,
    block_id INTEGER NOT NULL,
    FOREIGN KEY (block_id) REFERENCES block(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS root (
    root BLOB,
    network INTEGER NOT NULL,
    deposit_id INTEGER NOT NULL,
    PRIMARY KEY (deposit_id),
    FOREIGN KEY (deposit_id) REFERENCES deposit(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS block (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    block_num INTEGER,
    block_hash BLOB NOT NULL UNIQUE,
    network_id INTEGER
);

-- insert the block with block_id = 0 for the trusted exit root table
INSERT INTO block (id, block_hash) VALUES (0, X'0000000000000000000000000000000000000000000000000000000000000000');

CREATE TABLE IF NOT EXISTS claim (
    network_id INTEGER NOT NULL,
    claim_index INTEGER NOT NULL,
    orig_net INTEGER,
    orig_addr BLOB NOT NULL,
    amount TEXT,
    dest_addr BLOB NOT NULL,
    block_id INTEGER NOT NULL,
    tx_hash BLOB NOT NULL,
    rollup_index INTEGER DEFAULT 0 NOT NULL,
    mainnet_flag boolean DEFAULT false NOT NULL,
    PRIMARY KEY (claim_index, rollup_index, network_id, mainnet_flag),
    FOREIGN KEY (block_id) REFERENCES block(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS deposit (
    leaf_type integer,
    network_id INTEGER NOT NULL,
    orig_net INTEGER,
    orig_addr BLOB NOT NULL,
    amount TEXT,
    dest_net INTEGER NOT NULL,
    dest_addr BLOB NOT NULL,
    block_id INTEGER NOT NULL,
    deposit_cnt INTEGER NOT NULL,
    tx_hash BLOB NOT NULL,
    metadata BLOB NOT NULL,
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    ready_for_claim boolean DEFAULT false NOT NULL,
    FOREIGN KEY (block_id) REFERENCES block(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS exit_root (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    block_id INTEGER,
    global_exit_root BLOB,
    mainnet_exit_root BLOB,
    rollups_exit_root BLOB,
    network_id INTEGER DEFAULT 0 NOT NULL,
    allowed boolean DEFAULT true NOT NULL,
    FOREIGN KEY (block_id) REFERENCES block(id) ON DELETE CASCADE,
    UNIQUE (block_id, global_exit_root, network_id)
);

CREATE TABLE IF NOT EXISTS monitored_txs (
    deposit_id INTEGER NOT NULL,
    from_addr BLOB NOT NULL,
    to_addr BLOB,
    nonce INTEGER NOT NULL,
    value TEXT,
    data BLOB,
    gas INTEGER NOT NULL,
    status TEXT NOT NULL,
    history TEXT,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL,
    group_id INTEGER,
    global_exit_root BLOB DEFAULT X'0000000000000000000000000000000000000000' NOT NULL,
    PRIMARY KEY (deposit_id)
);

CREATE TABLE IF NOT EXISTS monitored_txs_group (
    group_id INTEGER NOT NULL,
    status TEXT NOT NULL,
    deposit_ids TEXT,
    num_retries integer NOT NULL,
    compressed_tx_data BLOB,
    claim_tx_history TEXT,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL,
    last_log TEXT,
    PRIMARY KEY (group_id)
);

CREATE TABLE IF NOT EXISTS remove_exit_root (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    block_id INTEGER,
    global_exit_root BLOB,
    network_id INTEGER,
    FOREIGN KEY (block_id) REFERENCES block(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS token_wrapped (
    network_id INTEGER NOT NULL,
    orig_net INTEGER NOT NULL,
    orig_token_addr BLOB NOT NULL,
    wrapped_token_addr BLOB NOT NULL,
    block_id INTEGER NOT NULL,
    name TEXT,
    symbol TEXT,
    decimals integer,
    PRIMARY KEY (network_id, orig_net, orig_token_addr),
    FOREIGN KEY (block_id) REFERENCES block(id) ON DELETE CASCADE
);

CREATE INDEX rht_key_idx ON rht (key);

CREATE INDEX root_idx ON root (root);

CREATE INDEX root_network_idx ON root (root, network);

CREATE INDEX block_idx ON exit_root (block_id);

CREATE INDEX claim_block_id ON claim (block_id);

CREATE INDEX deposit_block_id ON deposit (block_id);

CREATE INDEX token_wrapped_block_id ON token_wrapped (block_id);

-- +migrate Down
DROP TABLE IF EXISTS token_wrapped;
DROP TABLE IF EXISTS remove_exit_root;
DROP TABLE IF EXISTS exit_root;
DROP TABLE IF EXISTS claim;
DROP TABLE IF EXISTS deposit;
DROP TABLE IF EXISTS block;
DROP TABLE IF EXISTS root;
DROP TABLE IF EXISTS rollup_exit;
DROP TABLE IF EXISTS rht;
DROP TABLE IF EXISTS monitored_txs_group;
DROP TABLE IF EXISTS monitored_txs;
