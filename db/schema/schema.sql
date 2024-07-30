-- schema/schema.sql
CREATE TABLE ip_logs (
    id SERIAL PRIMARY KEY,
    ip_address VARCHAR(45) NOT NULL,
    country VARCHAR(100),
    region VARCHAR(100),
    city VARCHAR(100),
    latitude DECIMAL(9, 6),
    longitude DECIMAL(9, 6),
    user_agent TEXT,
    referrer TEXT,
    request_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
