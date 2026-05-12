-- +goose Up
CREATE TABLE user_authentications (
    id SERIAL PRIMARY KEY,
    status TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NOW(),
    user_authid INT REFERENCES user_auths(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE userauthentications;
