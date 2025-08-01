-- +goose Up

CREATE TABLE items(
    id UUID PRIMARY KEY,
    product_name TEXT NOT NULL,
    product_description TEXT NOT NULL,
    price INTEGER NOT NULL,
    in_stock BOOLEAN NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    image_url TEXT NOT NULL
);

-- +goose Down
DROP TABLE items;