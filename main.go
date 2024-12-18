package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/rss/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// To load the env variables in the env file
	godotenv.Load(".env")
	PORT := os.Getenv("PORT")
	DB_URL := os.Getenv("DB_URL")
	fmt.Print(DB_URL)

	// lets establish a connection to our database use builtin go sql library
	connection := establishDatabaseConnection(DB_URL)

	apiConnection := apiConfig{
		DB: connection,
	}

	// initiating an instance of chi router
	router := chi.NewRouter()

	// Basic CORS for handling crs related errors
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	router.Use(middleware.Logger)

	v1Router := chi.NewRouter()
	v1Router.Get("/health", handleServerReadiness)
	v1Router.Get("/err", handleErrorREquest)
	v1Router.Post("/register", apiConnection.handleUser)
	v1Router.Get("/auth", apiConnection.middlewareAuth(apiConnection.getUserByApiKey))
	v1Router.Post("/create-feed", apiConnection.middlewareAuth(apiConnection.handleCreateFeed))
	router.Mount("/api/v1", v1Router)

	// create a server
	server := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	// start your server to start listen for incoming request
	fmt.Printf(`Server listening on PORT:%v`, PORT)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("An error occur while starting your server")
	}

}
