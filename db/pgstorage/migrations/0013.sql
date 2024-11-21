-- +migrate Up
ALTER TABLE sync.claim ALTER COLUMN network_id TYPE BIGINT;
ALTER TABLE sync.deposit ALTER COLUMN network_id TYPE BIGINT;
ALTER TABLE sync.token_wrapped ALTER COLUMN network_id TYPE BIGINT;
ALTER TABLE sync.block ALTER COLUMN network_id TYPE BIGINT;
ALTER TABLE sync.exit_root ALTER COLUMN network_id TYPE BIGINT;
ALTER TABLE mt.root ALTER COLUMN network TYPE BIGINT;

ALTER TABLE sync.block ALTER COLUMN id TYPE BIGINT;
CREATE SEQUENCE IF NOT EXISTS sync.block_id_seq;
ALTER TABLE sync.block ALTER COLUMN id SET NOT NULL;
ALTER TABLE sync.block ALTER COLUMN id SET DEFAULT nextval('sync.block_id_seq');
ALTER SEQUENCE sync.block_id_seq OWNED BY sync.block.id;

ALTER TABLE sync.exit_root ALTER COLUMN id TYPE BIGINT;
CREATE SEQUENCE IF NOT EXISTS sync.exit_root_id_seq;
ALTER TABLE sync.exit_root ALTER COLUMN id SET NOT NULL;
ALTER TABLE sync.exit_root ALTER COLUMN id SET DEFAULT nextval('sync.exit_root_id_seq');
ALTER SEQUENCE sync.exit_root_id_seq OWNED BY sync.exit_root.id;

ALTER TABLE sync.deposit ALTER COLUMN id TYPE BIGINT;
CREATE SEQUENCE IF NOT EXISTS sync.deposit_id_seq;
ALTER TABLE sync.deposit ALTER COLUMN id SET NOT NULL;
ALTER TABLE sync.deposit ALTER COLUMN id SET DEFAULT nextval('sync.deposit_id_seq');
ALTER SEQUENCE sync.deposit_id_seq OWNED BY sync.deposit.id;

-- +migrate Down
ALTER TABLE sync.claim ALTER COLUMN network_id TYPE INTEGER;
ALTER TABLE sync.deposit ALTER COLUMN network_id TYPE INTEGER;
ALTER TABLE sync.token_wrapped ALTER COLUMN network_id TYPE INTEGER;
ALTER TABLE mt.root ALTER COLUMN network TYPE INTEGER;
ALTER TABLE sync.block ALTER COLUMN network_id TYPE INTEGER;
ALTER TABLE sync.exit_root ALTER COLUMN network_id TYPE INTEGER;

-- No need to revert the SERIAL to BIGSERIAL type changed