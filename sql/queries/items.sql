-- name: GetItems :many
SELECT id, product_name, product_description, price, in_stock, image_url FROM items;

-- name: CreateItem :one
INSERT INTO items (id, product_name, product_description, price, in_stock, updated_at, image_url)
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

-- name: ResetItems :exec
TRUNCATE items;

-- name: DeleteItem :exec
DELETE FROM items 
WHERE id = $1;