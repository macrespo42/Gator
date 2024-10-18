-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
  )
RETURNING *;

-- name: GetFeeds :many
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds;

-- name: GetFeedByUrl :one
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :one
UPDATE feeds
SET updated_at=NOW(), last_fetched_at=NOW()
WHERE feeds.id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT id, created_at, updated_at, name, url, user_id, last_fetched_at
FROM feeds
ORDER BY last_fetched_at DESC
LIMIT 1;
