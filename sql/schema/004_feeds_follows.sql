-- +goose Up
CREATE TABLE feeds_follows(
    id VARCHAR(36) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id VARCHAR(36) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id VARCHAR(36) NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    UNIQUE KEY (user_id,feed_id)
);
-- +goose Down
DROP TABLE feeds_follows;

