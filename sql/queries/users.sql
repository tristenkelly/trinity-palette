-- name: CreateUser :one
INSERT INTO users (id, username, email, hashed_password, created_at, updated_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetUser :exec
SELECT * FROM users
WHERE id = $1;