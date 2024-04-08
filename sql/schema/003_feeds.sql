-- +goose Up
CREATE TABLE feeds(
    id VARCHAR(36) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    UNIQUE KEY (url(255)),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE feeds;