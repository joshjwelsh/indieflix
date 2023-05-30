package main

import (
	"fmt"
	"log"
	"main/config"
	"main/db"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config := config.NewDBConfig(os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	db := db.NewDB(config)
	db.Ping()
	fmt.Printf("Successfully connected! %v", config.ConnectionString())
}
