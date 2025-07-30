-- name: GetItems :many
SELECT product_name, product_description, price, in_stock FROM items;

-- name: CreateItem :one
INSERT INTO items (id, product_name, product_description, price, in_stock, updated_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: ResetItems :exec
TRUNCATE items;