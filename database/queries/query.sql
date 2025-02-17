-- name: CreateUser :one
INSERT INTO users (
    id,
    email,
    hashed_password,
    created_at
) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;
