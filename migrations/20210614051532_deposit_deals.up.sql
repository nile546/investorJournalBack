CREATE TABLE deposit_deals(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    bank_instrument_id BIGINT REFERENCES banks_instruments (id),
    strategy_id BIGINT REFERENCES strategies (id),
    enter_datetime TIMESTAMP NOT NULL,
    percent FLOAT,
    exit_datetime TIMESTAMP,    
    StartDeposit  BIGINT,
	EndDeposit    BIGINT,
    user_id BIGINT REFERENCES users (id)
)

