package main

import (
	"log"
	"main/config"
	"main/env"
	"main/migrations"
	"main/router"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	log.Printf("Loading configs\n")
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	time.Sleep(15 * time.Second)
	config := config.NewDBConfig(os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	env := env.NewEnv(config)
	err = migrations.RunMigrations(env.DB, "migrations/sql")
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}
	r := router.NewRouterTree(env)
	log.Printf("Starting server on port 3000\n")
	http.ListenAndServe(":3000", env.Session.LoadAndSave(r))
}
