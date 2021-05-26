ALTER TABLE users
ADD CONSTRAINT UC_users_login UNIQUE (login);
ALTER TABLE users
ADD CONSTRAINT UC_users_email UNIQUE (email);
