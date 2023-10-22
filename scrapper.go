package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/sidkhuntia/go-rss/internal/database"
)

func startScrapping(db *database.Queries, councurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Starting scrapping with %d workers every %v duration", councurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		log.Println("Fetching RSS feeds")

		feedsTofetch, err := db.GetNextFeedsToFetch(context.Background(), int32(councurrency))
		if err != nil {
			log.Printf("Error while fetching feeds: %v", err)
			continue
		}

		wg := sync.WaitGroup{}

		for _, feed := range feedsTofetch {
			wg.Add(1)
			go scrapeFeed(db, feed, &wg)
		}
		wg.Wait()

	}
}

func scrapeFeed(db *database.Queries, feed database.Feed, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("Fetching feed %s", feed.Url)

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)

	if err != nil {
		log.Printf("Error while marking feed as fetched: %v", err)
		return
	}

	rssFeed, err := fetchRSS(feed.Url)

	if err != nil {
		log.Printf("Error while fetching feed %s: %v", feed.Url, err)
		return
	}
	postCounter := 0
	for _, item := range rssFeed.Channel.Item {

		description := sql.NullString{}

		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Error while parsing date %s: %v", item.PubDate, err)
			continue
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Updatedat:   time.Now(),
			Createdat:   time.Now(),
			Title:       item.Title,
			Description: description,
			Url:         item.Link,
			Feedid:      feed.ID,
			Publishedat: pubAt,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("Error while creating post: %v", err)
			continue
		}
		postCounter++
	}
	log.Printf("Fetched feed from %s, found %v posts found", feed.Name, postCounter)
}
