-- +goose Up
CREATE TABLE userauthentications (
    userauthenticationid SERIAL PRIMARY KEY,
    authentication_status TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    userauthid INT REFERENCES userauths(userauthid) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE userauthentications;
