-- +goose Up
ALTER TABLE businesses ADD street VARCHAR(20);

-- +goose Down

