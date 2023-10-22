package main

import (
	"fmt"
	"net/http"

	"github.com/sidkhuntia/go-rss/internal/auth"
	"github.com/sidkhuntia/go-rss/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

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
