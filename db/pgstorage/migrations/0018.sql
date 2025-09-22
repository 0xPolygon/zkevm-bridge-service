-- +migrate Up
ALTER TABLE sync.deposit ADD COLUMN IF NOT EXISTS ignore BOOLEAN DEFAULT FALSE;

-- +migrate Down
ALTER TABLE sync.deposit DROP COLUMN IF EXISTS ignore;
