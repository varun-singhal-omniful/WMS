CREATE TABLE IF NOT EXISTS inventories (
    id SERIAL PRIMARY KEY,
    hub_id INT NOT NULL,
    sku_id INT NOT NULL,
    quantity INT NOT NULL
);