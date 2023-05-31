package main

import (
	"log"
	"main/config"
	"main/env"
	"main/handler"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config := config.NewDBConfig(os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	env := env.NewEnv(config)
	r := chi.NewRouter()

	r.Route("/sources", func(r chi.Router) {
		r.Get("/", handler.GetAllSources(env))
		r.Get("/{id}", handler.GetSources(env))
		r.Post("/", handler.CreateSources(env))
		r.Delete("/{id}", handler.DeleteSources(env))
		r.Put("/{id}", handler.UpdateSources(env))
	})

	r.Route("/users", func(r chi.Router) {

	})

	r.Route("/genres", func(r chi.Router) {

	})

	log.Printf("Starting server on port 3000\n")
	http.ListenAndServe(":3000", r)
}
