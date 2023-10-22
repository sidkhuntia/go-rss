package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/sidkhuntia/go-rss/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

// main function is the entry point of the application.
// It loads environment variables, initializes database connection, sets up API routes and starts the server.
func main() {
	// Load environment variables from .env file
	godotenv.Load(".env")

	// Get the port number from environment variable
	portString := os.Getenv("PORT")

	// If port number is not set, default to 8080
	if portString == "" {
		portString = "8080"
		log.Println("No PORT environment variable detected, defaulting to " + portString)
	}

	// Get the database URL from environment variable
	dbURL := os.Getenv("DB_URL")

	// If database URL is not set, log a message
	if dbURL == "" {
		log.Println("No DB_URL environment variable detected")
	}

	// Open a connection to the database
	dbConcc, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	// Initialize API configuration with database connection
	apiCfg := apiConfig{
		DB: database.New(dbConcc),
	}

	// Create a new router
	router := chi.NewRouter()

	// Use CORS middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "OPTIONS", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Create a new sub-router for version 1 of the API
	v1router := chi.NewRouter()

	// Set up routes for health check and error handling
	v1router.Get("/healthz", handlerReadiness)
	v1router.Get("/error", handlerError)

	// Set up routes for user management
	v1router.Post("/users", apiCfg.handlerCreateUser)
	v1router.Get("/users/me", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	// Set up routes for feed management
	v1router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1router.Get("/feeds", apiCfg.handlerGetFeeds)

	// Set up routes for feed follow management
	v1router.Post("/feeds/follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1router.Get("/feeds/follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1router.Delete("/feeds/follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	// Mount the sub-router for version 1 of the API
	router.Mount("/v1", v1router)

	// Start the server
	fmt.Println("Listening on port " + portString)
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
