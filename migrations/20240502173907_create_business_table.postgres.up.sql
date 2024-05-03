CREATE TABLE businesses (
    id SERIAL PRIMARY KEY,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    description character varying(255),
    sector character varying(255) NOT NULL,
    is_corporate_affairs BOOLEAN DEFAULT FALSE NOT NULL,
    logo character varying(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_timestamp_businesses()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language plpgsql;

CREATE TRIGGER update_timestamp_businesses_trigger
BEFORE UPDATE ON businesses
FOR EACH ROW
WHEN (NEW.updated_at IS DISTINCT FROM OLD.updated_at)
EXECUTE FUNCTION update_timestamp_businesses();