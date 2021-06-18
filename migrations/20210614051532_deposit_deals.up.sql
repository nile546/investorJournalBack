CREATE TABLE deposit_deals(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    bank_instrument_id BIGINT REFERENCES banks_instruments (id),
    currency SMALLINT,
    enter_datetime TIMESTAMP NOT NULL,
    percent FLOAT,
    exit_datetime TIMESTAMP,    
    start_deposit  BIGINT,
	end_deposit    BIGINT,
    result BIGINT,
    user_id BIGINT REFERENCES users (id)
)

