package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// repsondWithError is a function that sends an error response with a given message and status code to the client.
// If the status code is greater than 499, it logs the error message with a 500 status code.
// It takes in a http.ResponseWriter, an integer status code and a string message as parameters.
// It returns nothing.
func repsondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 500 error: " + msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{Error: msg})
}

// respondWithJSON encodes the given payload to JSON format and writes it to the http.ResponseWriter.
// It sets the HTTP status code and Content-Type header before writing the response.
// If there is an error while encoding the payload, it logs the error and returns a 500 status code.
// Parameters:
// - w: http.ResponseWriter - the response writer to write the JSON response to
// - code: int - the HTTP status code to set for the response
// - payload: interface{} - the data to be encoded to JSON and sent as response

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error encoding JSON %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
