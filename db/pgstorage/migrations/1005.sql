-- +migrate Down
alter table sync.block alter column received_at drop default;

-- +migrate Up
alter table sync.block alter column received_at set default current_timestamp;
