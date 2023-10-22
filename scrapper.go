package main

import (
	"context"
	"log"
	"sync"
	"time"

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

	for _, item := range rssFeed.Channel.Item {
	
		log.Println(item.Title)
	}
	log.Printf("Fetched feed from %s, found %v posts found", feed.Url, len(rssFeed.Channel.Item))
}
