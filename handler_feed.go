package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sidkhuntia/go-rss/internal/database"
)

// handlerCreateFeed handles the creation of a new feed by decoding the request body into a parameters struct,
// creating a new feed in the database using the decoded parameters and returning the created feed as a JSON response.
// If there is an error parsing the JSON or creating the feed, an appropriate error response is returned.
func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		repsondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Url:       params.URL,
		Userid:    user.ID,
		Updatedat: time.Now(),
		Createdat: time.Now(),
	})

	if err != nil {
		repsondWithError(w, 500, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

// handlerGetFeeds handles the GET request to retrieve all the feeds from the database.
// It returns a JSON response containing the feeds and an error response if there was an error retrieving the feeds.
func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		repsondWithError(w, 500, fmt.Sprintf("Error getting feeds: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}
