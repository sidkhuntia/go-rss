package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/sidkhuntia/go-rss/internal/database"
)

// handlerCreateFeedFollow handles the creation of a new feed follow for a given user.
// It decodes the request body to extract the feed ID and creates a new feed follow in the database.
// If successful, it responds with the created feed follow in JSON format.
func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		repsondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	feed_follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		Updatedat: time.Now(),
		Createdat: time.Now(),
		Userid:    user.ID,
		Feedid:    params.FeedID,
	})

	if err != nil {
		repsondWithError(w, 500, fmt.Sprintf("Error following feed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feed_follow))
}

// handlerGetFeedFollows handles GET requests to retrieve the list of feeds followed by a user.
func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	// Get the list of feeds followed by the user from the database.
	feed_follows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		// If there is an error while getting the list of feeds followed by the user, respond with an error message.
		repsondWithError(w, 500, fmt.Sprintf("Error getting followed feeds: %v", err))
		return
	}

	// Convert the list of database.FeedFollows to a list of FeedFollows and respond with a JSON containing the list.
	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feed_follows))
}

// handlerDeleteFeedFollow handles the DELETE request to delete a feed follow from the database.
func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	
	// It parses the feedFollowID from the URL parameter and deletes the feed follow from the database.
	feedFollowIDString := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDString)
	
	// If there is an error parsing the feedFollowID or deleting the feed follow, it responds with an error message.
	if err != nil {
		repsondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		Userid: user.ID,
		ID:     feedFollowID,
	})
	
	if err != nil {
		repsondWithError(w, 500, fmt.Sprintf("Error deleting feed: %v", err))
		return
	}
	
	// Otherwise, it responds with a 200 status code and an empty JSON object.
	respondWithJSON(w, 200, struct{}{})
}
