package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// fmt.Println("Hello World")
	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		portString = "8080"
		log.Println("No PORT environment variable detected, defaulting to " + portString)
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "OPTIONS", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()
	v1router.Get("/healthz", handlerReadiness)
	v1router.Get("/error", handlerError)

	router.Mount("/v1", v1router)

	fmt.Println("Listening on port " + portString)
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
