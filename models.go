package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/sidkhuntia/go-rss/internal/database"
)

// types to desrcibe our return types in the API
type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
	UserId    uuid.UUID `json:"user_id"`
}
type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	URL         string    `json:"url"`
	PublishedAt time.Time `json:"published_at"`
	FeedId      uuid.UUID `json:"feed_id"`
}

// helper functions to convert from database types to API types
func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		CreatedAt: dbFeed.Createdat,
		UpdatedAt: dbFeed.Updatedat,
		URL:       dbFeed.Url,
		UserId:    dbFeed.Userid,
	}
}
func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, len(dbFeeds))
	for i, dbFeed := range dbFeeds {
		feeds[i] = databaseFeedToFeed(dbFeed)
	}
	return feeds
}
func databaseFeedFollowToFeedFollow(dbFeed database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.Createdat,
		UpdatedAt: dbFeed.Updatedat,
		UserId:    dbFeed.Userid,
		FeedId:    dbFeed.Feedid,
	}
}
func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := make([]FeedFollow, len(dbFeedFollows))
	for i, dbFeedFollow := range dbFeedFollows {
		feedFollows[i] = databaseFeedFollowToFeedFollow(dbFeedFollow)
	}
	return feedFollows
}
func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.Createdat,
		UpdatedAt: dbUser.Updatedat,
	}
}

func databasePostToPost(dbPost database.Post) Post {
	description := dbPost.Description.String
	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.Createdat,
		UpdatedAt:   dbPost.Updatedat,
		Title:       dbPost.Title,
		Description: &description,
		URL:         dbPost.Url,
		PublishedAt: dbPost.Publishedat,
		FeedId:      dbPost.Feedid,
	}
}

func databasePostsToPosts(dbPosts []database.Post) []Post {
	posts := make([]Post, len(dbPosts))
	for i, dbPost := range dbPosts {
		posts[i] = databasePostToPost(dbPost)
	}
	return posts
}
