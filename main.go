package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Hello world")

	godotenv.Load(".env")

	portstring := os.Getenv("PORT")

	if portstring == "" {
		log.Fatal("PORT is empty")
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

	router.Mount("/v1", arouter)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portstring,
	}

	log.Printf("server running on %v", portstring)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)

	}

}
