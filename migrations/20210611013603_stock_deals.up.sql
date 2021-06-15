CREATE TABLE stock_deals(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    stock_instrument_id BIGINT REFERENCES stocks_instruments (id),
    currency SMALLINT,
    strategy_id BIGINT REFERENCES strategies (id),
    pattern_id BIGINT REFERENCES patterns (id),
    position SMALLINT, 
    time_frame SMALLINT,
    enter_datetime TIMESTAMP NOT NULL,
    enter_point BIGINT,
    stop_loss BIGINT,
    quantity INT,
    exit_datetime TIMESTAMP,
    exit_point BIGINT,
    risk_ratio FLOAT,
    variability bool,
    user_id BIGINT REFERENCES users (id)
)

