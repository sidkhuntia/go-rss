package main

import (
	"fmt"
	"net/http"

	"github.com/sidkhuntia/go-rss/internal/auth"
	"github.com/sidkhuntia/go-rss/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

// middlewareAuth is a function that takes an authHandler and returns an http.HandlerFunc.
// It checks the API key in the request header and retrieves the user from the database.
// If the API key is invalid or the user cannot be retrieved, it responds with an error.
// Otherwise, it calls the authHandler with the response writer, request, and user.

// we can returing a anonymous function that takes an http.ResponseWriter and *http.Request beacause chi Router expects a http.HandlerFunc

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			repsondWithError(w, 401, fmt.Sprintf("Error getting API key: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUser(r.Context(), apiKey)

		if err != nil {
			repsondWithError(w, 401, fmt.Sprintf("Error getting user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
