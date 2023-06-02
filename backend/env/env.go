package env

import (
	"database/sql"
	"log"
	"main/config"
	"time"

	"github.com/alexedwards/scs/v2"
	_ "github.com/lib/pq"
)

type Env struct {
	*sql.DB
	Session *scs.SessionManager
}

func createDB(config config.DBConfig) *sql.DB {
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
	return db
}

func createSession() *scs.SessionManager {
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	return session
}

func NewEnv(config config.DBConfig) *Env {
	db := createDB(config)
	session := createSession()
	return &Env{db, session}
}
