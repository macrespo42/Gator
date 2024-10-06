-- +goose up

CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name VARCHAR(32) UNIQUE NOT NULL
);

-- +goose down
DROP TABLE users;
