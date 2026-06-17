-- +goose Up
CREATE TABLE user_profiles (
    id SERIAL PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    user_auths_id INTEGER REFERENCES user_auths(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE user_profiles;
