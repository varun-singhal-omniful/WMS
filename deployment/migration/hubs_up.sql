CREATE TABLE IF NOT EXISTS hubs (
    id SERIAL PRIMARY KEY,
    tenant_id INT NOT NULL,
    manager_name VARCHAR(255) NOT NULL,
    manager_contact VARCHAR(20) NOT NULL,
    manager_email VARCHAR(255) NOT NULL
);