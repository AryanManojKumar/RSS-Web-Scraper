package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"webscrapper/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	fmt.Println("Hello world")

	godotenv.Load(".env")

	portstring := os.Getenv("PORT")

	if portstring == "" {
		log.Fatal("PORT is empty")
	}

	db := os.Getenv("DB_URL")

	if db == "" {
		log.Fatal("db url not found in environment")
	}

	connection, err := sql.Open("postgres", db)
	if err != nil {
		log.Fatal("Couldnt connect to the database", err)
	}

	apiCfg := apiConfig{
		DB: database.New(connection),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"}, // allow all HTTPS & HTTP origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"}, // allow all headers
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // cache preflight request for 5 minutes
	}))

	arouter := chi.NewRouter()

	arouter.Get("/healthz", readiness)
	arouter.Get("/error", errorhandler)
	arouter.Post("/users", apiCfg.handleuserdb)
	arouter.Get("/users", apiCfg.middlewearauth(apiCfg.handlerGetUserFromApi))
	arouter.Post("/feed", apiCfg.middlewearauth(apiCfg.handleuserfeed))
	arouter.Get("/feed", apiCfg.handlerGetfeeds)

	router.Mount("/v1", arouter)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portstring,
	}

	log.Printf("server running on %v", portstring)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)

	}

}
