package db

import (
	"database/sql"
	"log"
	"main/config"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewDB(config config.DBConfig) *DB {
	connectionString := config.ConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Connection string %v failed error: %v", connectionString, err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("Ping failed error: %v", err)
	}
	return &DB{db}
}
