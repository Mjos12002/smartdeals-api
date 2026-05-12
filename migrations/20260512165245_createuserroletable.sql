-- +goose Up
CREATE TABLE userroles (
    userole_id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    userauth_id INT REFERENCES userauths(userauthid) ON DELETE CASCADE,
    roles_id INT REFERENCES roles(roles_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE userroles;
