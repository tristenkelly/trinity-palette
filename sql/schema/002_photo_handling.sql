-- +goose Up
ALTER TABLE items
ADD COLUMN product_image TEXT NOT NULL;


-- +goose Down
ALTER TABLE items
DROP COLUMN product_image;