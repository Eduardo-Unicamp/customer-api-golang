CREATE TABLE IF NOT EXISTS products(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    price NUMERIC(15,2) NOT NULL,
    stock INTEGER NOT NULL

);