-- name: CreatePost :one
INSERT INTO posts (title, body, created_at, updated_at, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetPosts :many
SELECT posts.title, posts.body, posts.created_at, posts.updated_at, users.username
FROM posts
INNER JOIN users ON posts.user_id = users.id
ORDER BY posts.created_at ASC
LIMIT 10;