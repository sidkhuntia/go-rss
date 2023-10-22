package main

import (
	"net/http"
)

// handlerError handles errors and responds with an error message and status code 400.
func handlerError(w http.ResponseWriter, r *http.Request) {
	repsondWithError(w, 400, "Something went wrong")
}
