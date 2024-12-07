-- name: CreateUser :one
INSERT INTO users (
    id, created_at, updated_at, name, email, api_key, phone_number
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: FindUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: FindUserByApiKey :one
SELECT * FROM users
WHERE api_key = $1;