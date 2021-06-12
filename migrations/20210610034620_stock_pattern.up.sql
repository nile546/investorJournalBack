CREATE TABLE stock_patterns (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR,
    description TEXT,
    icon TEXT,
    user_id BIGINT REFERENCES users (id) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL

)
