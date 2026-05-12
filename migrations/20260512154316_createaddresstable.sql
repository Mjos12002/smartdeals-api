-- +goose Up
CREATE TABLE addresses (
    id SERIAL PRIMARY KEY,
    street TEXT,
    popularname TEXT,
    province TEXT,
    district TEXT,
    sector TEXT,
    longlat TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE addresses;
