CREATE TABLE stock_strategies (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR,
    description TEXT,
    user_id BIGINT REFERENCES users (id) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
)
