CREATE TABLE IF NOT EXISTS orders(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    status VARCHAR(10),
    customer_id UUID,
    CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES customers(id)
    
);


CREATE TABLE IF NOT EXISTS order_items(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    selling_price NUMERIC(15,2) NOT NULL,
    units INTEGER NOT NULL,
    product_id UUID NOT NULL,
    order_id UUID NOT NULL,
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products(id),
    CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders(id)
);