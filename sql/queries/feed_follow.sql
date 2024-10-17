-- name: CreateFeedFollow :one
WITH inserted_feed_follow as (
INSERT INTO feed_follow (id, created_at, updated_at, user_id, feed_id)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *)
SELECT inserted_feed_follow.*, 
feeds.name AS feed_name,
users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
INNER JOIN users ON inserted_feed_follow.user_id = users.id;

-- name: GetFeedFollowForUser :many
SELECT feed_follow.id,
feed_follow.created_at,
feed_follow.updated_at,
feed_follow.user_id,
feed_follow.feed_id,
feeds.name AS feed_name,
users.name AS user_name FROM feed_follow
INNER JOIN feeds ON feed_follow.feed_id = feeds.id
INNER JOIN users ON feed_follow.user_id = users.id
WHERE users.name = $1;

-- name: DeleteFeedFollow :one
DELETE FROM feed_follow
USING feeds, users
WHERE feed_follow.feed_id = feeds.id
AND feed_follow.user_id = users.id
AND users.name = $1
AND feeds.url = $2
RETURNING feed_follow.*;
