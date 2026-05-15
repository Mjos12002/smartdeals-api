-- +goose Up
CREATE TABLE user_roles (
    id SERIAL PRIMARY KEY,
    user_auth_id INT REFERENCES user_auths(id) ON DELETE CASCADE,
    role_id INT REFERENCES roles(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE userroles;
