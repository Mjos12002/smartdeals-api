-- +goose Up
CREATE TABLE user_authentications (
    id SERIAL PRIMARY KEY,
    status TEXT,
    user_auths_id INT REFERENCES user_auths(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE user_authentications;
