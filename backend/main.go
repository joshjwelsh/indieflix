package main

import (
	"fmt"
	"log"
	"main/config"
	"main/env"
	"main/handler"
	"main/model"
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

	insert_result, err := model.InsertSources(env.DB, "test", "test")
	if err != nil {
		log.Fatalf("Error inserting sources: %v", err)
	}
	fmt.Printf("Insert Result: %v", insert_result)
	http.Handle("/sources", handler.GetAllSources(env))
	http.Handle("/sources/", handler.GetSource(env))
	http.ListenAndServe(":3000", nil)
}
