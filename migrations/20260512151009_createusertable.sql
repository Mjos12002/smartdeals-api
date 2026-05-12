-- +goose Up
CREATE TABLE userdetails (
    id SERIAL PRIMARY KEY,
    email TEXT,
    phone TEXT,
    u_address TEXT,
    fname TEXT,
    lname TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE userdetails;
