-- name: CreateUser :one
INSERT INTO users (id, username, email, hashed_password, created_at, updated_at, is_admin)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;

-- name: GetUser :exec
SELECT * FROM users
WHERE id = $1;

-- name: GetPassHash :one
SELECT * from users
WHERE email = $1;