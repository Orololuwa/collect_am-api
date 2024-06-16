ALTER TABLE kyc
    ADD COLUMN business_id INTEGER,
    ADD CONSTRAINT fk_business_id 
        FOREIGN KEY (business_id) 
        REFERENCES businesses(id) 
        ON DELETE  CASCADE;

ALTER TABLE businesses
    ADD COLUMN user_id INTEGER UNIQUE,
    ADD CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE;