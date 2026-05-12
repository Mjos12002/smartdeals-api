-- +goose Up
CREATE TABLE addresses (
    id SERIAL PRIMARY KEY,
    street TEXT,
    popular_name TEXT,
    province TEXT,
    district TEXT,
    sector TEXT,
    long_lat TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NOW()

);

-- +goose Down
DROP TABLE addresses;
