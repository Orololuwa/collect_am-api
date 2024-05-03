CREATE TABLE kyc (
    id SERIAL PRIMARY KEY,
    certificate_of_registration character varying(255),
    proof_of_address character varying(255),
    bvn character varying(255),
    business_address text,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_timestamp_kyc()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language plpgsql;

CREATE TRIGGER update_timestamp_kyc_trigger
BEFORE UPDATE ON kyc
FOR EACH ROW
WHEN (NEW.updated_at IS DISTINCT FROM OLD.updated_at)
EXECUTE FUNCTION update_timestamp_kyc();