-- +goose Up
ALTER TABLE users
ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT encode(sha256(random()::text::bytea), 'hex');

ALTER TABLE users
ADD COLUMN phone_number VARCHAR(15) NOT NULL DEFAULT '+234';

-- +goose Down
ALTER TABLE users DROP COLUMN phone_number;
ALTER TABLE users DROP COLUMN api_key;
