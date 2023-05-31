package env

import (
	"database/sql"
	"log"
	"main/config"
	"time"

	_ "github.com/lib/pq"
)

type Env struct {
	*sql.DB
}

func NewEnv(config config.DBConfig) *Env {
	connectionString := config.ConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Connection string %v failed error: %v", connectionString, err)
	}

	// Maximum Idle Connections
	db.SetMaxIdleConns(5)
	// Maximum Open Connections
	db.SetMaxOpenConns(10)
	// Idle Connection Timeout
	db.SetConnMaxIdleTime(1 * time.Second)
	// Connection Lifetime
	db.SetConnMaxLifetime(30 * time.Second)

	err = db.Ping()
	if err != nil {
		log.Fatalf("Ping failed error: %v", err)
	}
	return &Env{db}
}
