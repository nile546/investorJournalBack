CREATE TABLE stockdeal_parts (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  quantity INT,
  type SMALLINT,
  price BIGINT,
  datetime TIMESTAMP NOT NULL,
  stock_deal_id BIGINT REFERENCES stock_deals (id)
);
