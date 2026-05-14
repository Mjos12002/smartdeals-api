-- +goose Up
ALTER TABLE businesses ADD COLUMN phone VARCHAR(15) NOT NULL;

-- +goose Down
ALTER TABLE businesses DROP COLUMN phone;
