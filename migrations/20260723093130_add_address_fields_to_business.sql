-- +goose Up
ALTER TABLE businesses DROP addresses_id;
ALTER TABLE businesses ADD popular_name VARCHAR(20);
ALTER TABLE businesses ADD email VARCHAR(20);
ALTER TABLE businesses ADD phone_number VARCHAR(20);
ALTER TABLE businesses ADD twitter VARCHAR(20);
ALTER TABLE businesses ADD facebook VARCHAR(20);
ALTER TABLE businesses ADD instagram VARCHAR(20);
ALTER TABLE businesses ADD province VARCHAR(20);
ALTER TABLE businesses ADD district VARCHAR(20);


-- +goose Down