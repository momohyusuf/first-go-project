-- +goose Up
CREATE TABLE users (
id UUID PRIMARY KEY,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL,
email VARCHAR(64) NOT NULL UNIQUE, 
name TEXT NOT NULL
);
-- Create an index on the email column
CREATE INDEX idx_users_email ON users (email);
-- +goose Down
DROP INDEX IF EXISTS idx_users_email;
DROP TABLE users;