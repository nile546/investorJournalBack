CREATE TABLE crypto_deals(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    crypto_instrument_id BIGINT REFERENCES crypto_instruments (id),
    crypto_strategy_id BIGINT REFERENCES crypto_strategies (id),
    crypto_pattern_id BIGINT REFERENCES crypto_patterns (id),
    position SMALLINT, 
    time_frame SMALLINT,
    enter_datetime TIMESTAMP NOT NULL,
    enter_point BIGINT,
    stop_loss BIGINT,
    quantity INT,
    exit_datetime TIMESTAMP,
    exit_point BIGINT,
    risk_ratio FLOAT,
    user_id BIGINT REFERENCES users (id)
)

