-- +goose Up
CREATE TABLE posts(
    id VARCHAR(36) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    url TEXT NOT NULL,
    feed_id VARCHAR(36) NOT NULL,
    UNIQUE (url(255)),
    FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE posts;