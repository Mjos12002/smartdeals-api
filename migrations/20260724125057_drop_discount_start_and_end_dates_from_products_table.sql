-- +goose Up
ALTER TABLE products DROP discount_start_date;
ALTER TABLE products DROP discount_end_date;

-- +goose Down
