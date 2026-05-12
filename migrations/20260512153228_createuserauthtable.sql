-- +goose Up
CREATE TABLE userauths (
    userauthid SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    u_password TEXT NOT NULL,
    u_status TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    userdetails_id INT REFERENCES userdetails(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE userauths;
