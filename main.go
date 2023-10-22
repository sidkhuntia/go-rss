package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
	// Loading environment variables from .env file
	godotenv.Load(".env")

	// Getting the port number from environment variable
	portString := os.Getenv("PORT")

	// If port number is not set, defaulting to 8080
	if portString == "" {
		portString = "8080"
		log.Println("No PORT environment variable detected, defaulting to " + portString)
	}

	// taking the database URL from environment variable
	dbURL := os.Getenv("DB_URL")

	// If database URL is not set, logging a message
	if dbURL == "" {
		log.Println("No DB_URL environment variable detected")
	}

	// Opening a connection to the database
	dbConcc, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	// Initializing API configuration with database connection
	db := database.New(dbConcc)
	apiCfg := apiConfig{
		DB: db,
	}

	// Creating a new router
	router := chi.NewRouter()

	// starting scrapping
	go startScrapping(db, 10, time.Minute) // 10 workers every minute

	// Using CORS middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "OPTIONS", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Creating a new sub-router for version 1 of the API
	v1router := chi.NewRouter()

	// Setting up routes for health check and error handling
	v1router.Get("/healthz", handlerReadiness)
	v1router.Get("/error", handlerError)

	// Setting up routes for user management
	v1router.Post("/users", apiCfg.handlerCreateUser)
	v1router.Get("/users/me", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1router.Get("/users/me/posts", apiCfg.middlewareAuth(apiCfg.handlerPostsForUser))

	// Setting up routes for feed management
	v1router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1router.Get("/feeds", apiCfg.handlerGetFeeds)

	// Setting up routes for feed follow management
	v1router.Post("/feeds/follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1router.Get("/feeds/follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1router.Delete("/feeds/follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	// Mounting the sub-router for version 1 of the API
	router.Mount("/v1", v1router)

	// Starting the server
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
