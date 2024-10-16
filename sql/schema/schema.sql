CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name VARCHAR(32) UNIQUE NOT NULL
);

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
