package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sidkhuntia/go-rss/internal/database"
)

// handlerCreateUser creates a new user with the given name and returns the created user as JSON response.
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// Define a struct to hold the request parameters.
	type parameters struct {
		Name string `json:"name"`
	}

	// Decode the request body into the parameters struct.
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	// If there is an error decoding the request body, respond with an error message.
	if err != nil {
		repsondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Create a new user with the given name.
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Updatedat: time.Now(),
		Createdat: time.Now(),
	})

	// If there is an error creating the user, respond with an error message.
	if err != nil {
		repsondWithError(w, 500, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	// Respond with the created user as JSON.
	respondWithJSON(w, 201, databaseUserToUser(user))
}

// handlerGetUser handles GET requests to retrieve a user's information.
// It responds with a JSON object containing the user's information.
func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJSON(w, 200, databaseUserToUser(user))

}
