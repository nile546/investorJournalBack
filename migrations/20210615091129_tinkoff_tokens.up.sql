CREATE TABLE tinkoff_tokens (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  token VARCHAR NOT NULL,
  user_id BIGINT REFERENCES users (id)
);
