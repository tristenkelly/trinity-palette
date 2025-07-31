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

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByName :one
SELECT * FROM users
WHERE username = $1;

-- name: GetPassHash :one
SELECT * from users
WHERE email = $1;

-- name: ChangePass :exec
UPDATE users
SET hashed_password = $2
WHERE email = $1;

-- name: ChangeEmail :exec
UPDATE users
SET email = $2
WHERE username = $1;

-- name: isAdmin :one
SELECT is_admin
FROM users
WHERE id = $1;