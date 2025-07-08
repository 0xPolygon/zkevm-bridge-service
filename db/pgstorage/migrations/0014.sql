-- +migrate Up

ALTER TABLE sync.claim DROP CONSTRAINT claim_pkey;
ALTER TABLE sync.claim ADD PRIMARY KEY (index, rollup_index, network_id, mainnet_flag);

-- +migrate Down

ALTER TABLE sync.claim DROP CONSTRAINT claim_pkey;
ALTER TABLE sync.claim ADD PRIMARY KEY (index, rollup_index, network_id);