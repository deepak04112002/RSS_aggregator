-- name: CreateFeed :exec
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetFeedByID :one
SELECT * FROM feeds WHERE id = ?;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds 
ORDER BY 
  CASE 
    WHEN last_fetched_at IS NULL THEN 0 
    ELSE 1 
  END,
  last_fetched_at ASC
LIMIT ?;

-- name: MarkFeedAsFetched :exec
UPDATE feeds
SET last_fetched_at = NOW(),
    updated_at = NOW()
WHERE id = ?;

-- name: MarkFeedAsFetchedReturning :one
SELECT last_fetched_at FROM feeds WHERE id = ?;

-- name: GetUpdatedFeed :one
SELECT *
FROM feeds
WHERE id = ?;
