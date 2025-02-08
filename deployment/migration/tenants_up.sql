CREATE TABLE IF NOT EXISTS tenants (
    id SERIAL PRIMARY KEY,
    tenant_name VARCHAR(255) NOT NULL,
    registered_address TEXT NOT NULL,
    tenant_contact VARCHAR(20) NOT NULL,
    tenant_email VARCHAR(255) UNIQUE NOT NULL
);