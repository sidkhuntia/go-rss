package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sidkhuntia/go-rss/internal/auth"
	"github.com/sidkhuntia/go-rss/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		repsondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Updatedat: time.Now(),
		Createdat: time.Now(),
	})

	if err != nil {
		repsondWithError(w, 500, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)

	if err != nil {
		repsondWithError(w, 403, fmt.Sprintf("Error parsing API key: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUser(r.Context(), apiKey)

	if err != nil {
		repsondWithError(w, 400, fmt.Sprintf("Error getting user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))

}
