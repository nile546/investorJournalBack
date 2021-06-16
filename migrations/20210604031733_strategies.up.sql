CREATE TABLE strategies (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR,
    description TEXT,
    user_id BIGINT REFERENCES users (id) ON DELETE CASCADE,
    instrument_type SMALLINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
)
