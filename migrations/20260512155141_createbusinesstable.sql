-- +goose Up
CREATE TABLE businesses (
    id SERIAL PRIMARY KEY,
    name TEXT,
    contact TEXT,
    description TEXT,
    logo_url TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NOW(),
    address_id INT REFERENCES addresses(id) ON DELETE CASCADE,
    user_details_id INT REFERENCES user_details(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE businesses;
