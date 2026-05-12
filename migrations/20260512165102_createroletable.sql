-- +goose Up
CREATE TABLE roles (
    roles_id SERIAL PRIMARY KEY,
    r_name TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE roles;
