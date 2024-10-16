-- +goose up

CREATE TABLE feeds(
  id UUID PRIMARY KEY,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  name VARCHAR(255),
  url VARCHAR(255) UNIQUE,
  user_id UUID,
  CONSTRAINT fk_user
  FOREIGN KEY(user_id)
  REFERENCES users(id)
  ON DELETE CASCADE
);

-- +goose down
DROP TABLE feeds;
