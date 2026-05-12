-- +goose Up
CREATE TABLE user_auths (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    status TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NOW(),
    user_details_id INT REFERENCES user_details(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE user_auths;
