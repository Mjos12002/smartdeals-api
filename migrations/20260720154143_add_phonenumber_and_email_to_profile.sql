-- +goose Up
ALTER TABLE user_profiles ADD phone_number TEXT;
ALTER TABLE user_profiles ADD email TEXT;

-- +goose Down

