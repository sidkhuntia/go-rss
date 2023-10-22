package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts the API key from the Authorization header in the provided HTTP headers.
// It returns the API key as a string and an error if the header is missing or invalid.
// Accepted format:
// Authorization: Bearer <api key>

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no Authorization header provided")
	}

	keys := strings.Split(val, " ")

	if len(keys) != 2 {
		return "", errors.New("invalid Authorization header provided")
	}

	if keys[0] != "Bearer" {
		return "", errors.New("invalid Authorization header prefix provided")
	}

	return keys[1], nil

}
