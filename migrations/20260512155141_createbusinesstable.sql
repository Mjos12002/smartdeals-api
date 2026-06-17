-- +goose Up
CREATE TABLE businesses (
    id SERIAL PRIMARY KEY,
    name TEXT,
    description TEXT,
    logo_url TEXT,
    addresses_id INT REFERENCES addresses(id) ON DELETE CASCADE,
    user_profiles_id INT REFERENCES user_profiles(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE businesses;
