// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"
	"time"
)

type Feed struct {
	ID            string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	Url           string
	UserID        string
	LastFetchedAt sql.NullTime
}

type FeedsFollow struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string
	FeedID    string
}

type Post struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Description sql.NullString
	PublishedAt time.Time
	Url         string
	FeedID      string
}

type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}
