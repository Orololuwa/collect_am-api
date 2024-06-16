ALTER TABLE kyc
    DROP CONSTRAINT fk_business_id,
    DROP COLUMN business_id;

ALTER TABLE businesses
    DROP CONSTRAINT fk_user_id,
    DROP COLUMN user_id;