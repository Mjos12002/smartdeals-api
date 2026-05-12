-- +goose Up
CREATE TABLE businesses (
    id SERIAL PRIMARY KEY,
    b_name TEXT,
    b_contact TEXT,
    b_description TEXT,
    logo_url TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    address_id INT REFERENCES addresses(id) ON DELETE CASCADE,
    userdetails_id INT REFERENCES userdetails(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE businesses;
