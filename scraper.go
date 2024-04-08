package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/deepak04112002/rssagg/internal/database"
	"github.com/google/uuid"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {

	log.Printf("Scraping on %v goroutine every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching feed:", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, &feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed *database.Feed) {
	defer wg.Done()

	lastFetchedAt, err := db.MarkFeedAsFetchedReturning(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched:", err)
		return
	}

	// Check if last_fetched_at is valid and log the time
	if lastFetchedAt.Valid {
		log.Println("Last fetched at:", lastFetchedAt.Time)
	} else {
		log.Println("Last fetched at is null")
	}
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error Fetching feed:", err)
		return
	}
	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Couldn't Parse date %v with err %v", item.PubDate, err)
			continue
		}
		 err = db.CreatePost(context.Background(),
			database.CreatePostParams{
				ID:          uuid.NewString(),
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				Title:       item.Title,
				Description: description,
				PublishedAt: pubAt,
				Url:         item.Link,
				FeedID:      feed.ID,
			})
		if err != nil {
			if strings.Contains(err.Error(),"Duplicate entry"){
				continue
			}
			log.Println("failed to create post:", err)
		}
	}
	log.Printf("Feed %s collected,%v posts found", feed.Name, len(rssFeed.Channel.Item))
}
