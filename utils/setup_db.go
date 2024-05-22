package utils

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"log"
	"os"
)

// dbPath stores the path to the SQLite database
var dbPath string

// InitDB initializes the database connection and creates tables if necessary
func InitDB() error {
	dbPath = os.Getenv("DATABASE_PATH")

	if dbPath == "" {
		dbPath = "database/url.db"
	}

	if !checkDB() {

		_, err := os.Create(dbPath)
		if err != nil {
			log.Println()
			return err
		}

	}

	if err := createTables(); err != nil {
		log.Fatal("Failed to create tables:", err)
		return err
	}

	return nil
}

// createTables creates the necessary tables in the database
func createTables() error {

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to open database:", err)
		return err
	}

	createTableSQL := ` CREATE TABLE IF NOT EXISTS URLS(
    OriginalURL TEXT NOT NULL,
    ShortURL TEXT PRIMARY KEY,
    Clicks NUMBER NOT NULL,
    Creation NUMBER NOT NULL,
    Expiration NUMBER );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Println("Failed to create tables:", err)
		return err
	}

	return err
}

// checkDB checks if the database file exists
func checkDB() bool {
	_, err := os.Stat(dbPath)
	return !os.IsNotExist(err)
}
