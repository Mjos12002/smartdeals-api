-- +goose Up
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT,
    description TEXT,
    price DECIMAL(10, 2),
    discount DECIMAL(5, 2),
    discounted_price DECIMAL(10, 2),
    discount_start_date TIMESTAMP,
    discount_end_date TIMESTAMP,
    status TEXT,
    logo TEXT,
    product_categories_id INT REFERENCES product_categories(id) ON DELETE CASCADE,
    businesses_id INT REFERENCES businesses(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE products;
