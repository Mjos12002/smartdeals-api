-- +goose Up
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    role_name TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE roles;
