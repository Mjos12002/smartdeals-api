-- +goose Up
ALTER TABLE product_category RENAME TO product_categories;

-- +goose Down

