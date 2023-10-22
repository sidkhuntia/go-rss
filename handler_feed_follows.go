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

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feed_follows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		repsondWithError(w, 500, fmt.Sprintf("Error getting followed feeds: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feed_follows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDString := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDString)

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

	respondWithJSON(w, 200, struct{}{})
}
