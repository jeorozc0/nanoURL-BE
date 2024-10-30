// db/db.go
package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// Get DB URL from environment variable
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v\n", err)
	}

	// Create tables if they don't exist
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
			id VARCHAR(4) PRIMARY KEY,
			original_url TEXT NOT NULL,
			new_url TEXT NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatalf("Error creating table: %v\n", err)
	}
}
