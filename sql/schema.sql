CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name VARCHAR(32) UNIQUE NOT NULL
);

CREATE TABLE feeds(
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name VARCHAR(255) NOT NULL,
  url VARCHAR(255) UNIQUE NOT NULL,
  user_id UUID NOT NULL,
  CONSTRAINT fk_user
  FOREIGN KEY(user_id)
  REFERENCES users(id)
  ON DELETE CASCADE
);

CREATE TABLE feed_follow(
  id UUID PRIMARY KEY,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  user_id UUID NOT NULL,
  feed_id UUID NOT NULL,
  CONSTRAINT fk_user
  FOREIGN KEY(user_id)
  REFERENCES users(id)
  ON DELETE CASCADE,
  CONSTRAINT fk_feed
  FOREIGN KEY(feed_id)
  REFERENCES feeds(id)
  ON DELETE CASCADE,
  UNIQUE (user_id, feed_id)
);
