-- +migrate Up
ALTER TABLE sync.claim ADD COLUMN IF NOT EXISTS global_index VARCHAR(255) DEFAULT '';

-- +migrate Down
ALTER TABLE sync.claim DROP COLUMN IF EXISTS global_index;
