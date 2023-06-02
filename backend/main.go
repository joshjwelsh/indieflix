package main

import (
	"log"
	"main/api"
	"main/config"
	"main/env"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config := config.NewDBConfig(os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	env := env.NewEnv(config)
	r := api.NewRouterTree(env)
	log.Printf("Starting server on port 3000\n")
	http.ListenAndServe(":3000", env.Session.LoadAndSave(r))
}
