-- name: CreateUser :exec
INSERT INTO users (id, created_at, updated_at, name,api_key)
VALUES (?, ?, ?, ?,CONCAT(SHA2(UUID(), 256)));

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = ?;


