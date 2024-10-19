-- +goose up
CREATE TABLE posts(
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  title VARCHAR(255) NOT NULL,
  url VARCHAR(255) NOT NULL UNIQUE,
  description TEXT,
  published_at TIMESTAMP NOT NULL,
  feed_id UUID NOT NULL,
  CONSTRAINT fk_feed
  FOREIGN KEY(feed_id)
  REFERENCES feeds(id)
  ON DELETE CASCADE
);

-- +goose down
DROP TABLE posts;
