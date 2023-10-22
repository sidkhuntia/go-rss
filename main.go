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

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		portString = "8080"
		log.Println("No PORT environment variable detected, defaulting to " + portString)
	}

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Println("No DB_URL environment variable detected")
	}

	dbConcc, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	apiCfg := apiConfig{
		DB: database.New(dbConcc),
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
	v1router.Post("/users", apiCfg.handlerCreateUser)
	v1router.Get("/users/me", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))

	router.Mount("/v1", v1router)

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
