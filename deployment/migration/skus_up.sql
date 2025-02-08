CREATE TABLE IF NOT EXISTS skus (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    hub_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT
);