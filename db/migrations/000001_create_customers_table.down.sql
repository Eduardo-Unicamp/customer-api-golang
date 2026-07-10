CREATE TABLE IF NOT EXISTS customer (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
name VARCHAR(100) NOT NULL,
email VARCHAR(100) NOT NULL,
phone VARCHAR(20),
created_at TIMESTAMPTZ DEFAULT NOW(),
updated_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE OR REPLACE FUNCTION trigger_update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
	NEW.updated_at = NOW();
	RETURN NEW;
END;
$$LANGUAGE plpgsql;

CREATE TRIGGER update_timestamp
BEFORE UPDATE ON customer
FOR EACH row
EXECUTE FUNCTION trigger_update_timestamp();