-- name: CreateFeedFollow :exec
INSERT INTO feeds_follows (id, created_at, updated_at, user_id, feed_id)
VALUES (?, ?, ?, ?, ?);
-- name: GetFeedFollowByID :one
SELECT * FROM feeds_follows WHERE id = ?;
-- name: GetFeedFollows :many
SELECT * FROM feeds_follows;
-- name: DeleteFeedFollows :exec
DELETE FROM feeds_follows WHERE id = ? AND user_id = ?;