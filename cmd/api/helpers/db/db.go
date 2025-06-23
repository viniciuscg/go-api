package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not load .env file (might be using real env vars)")
	}

	dbSource := os.Getenv("DB_SOURCE")
	if dbSource == "" {
		return nil, ErrMissingDBSource
	}

	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("✅ Connected to the database")
	return db, nil
}

var ErrMissingDBSource = &MissingDBSourceError{}

type MissingDBSourceError struct{}

func (e *MissingDBSourceError) Error() string {
	return "DB_SOURCE environment variable is missing"
}
