package main

import (
	"net/http"
)

// handlerReadiness is a function that handles the readiness check for the server.
// It responds with an empty JSON object and a status code of 200.
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
