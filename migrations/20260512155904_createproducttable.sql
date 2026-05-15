-- +goose Up
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    product_name TEXT,
    product_description TEXT,
    price DECIMAL(10, 2),
    discount DECIMAL(5, 2),
    discounted_price DECIMAL(10, 2),
    discount_start_date TIMESTAMP,
    discount_end_date TIMESTAMP,
    product_status TEXT,
    product_category TEXT,
    image_url TEXT,
    business_id INT REFERENCES businesses(id) ON DELETE CASCADE
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE products;
