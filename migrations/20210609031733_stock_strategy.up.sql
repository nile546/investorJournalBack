CREATE TABLE stock_strategy (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR,
    description TEXT,
    user_id BIGINT REFERENCES users (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
)
