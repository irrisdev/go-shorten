package utils

import (
	"database/sql"
	"github.com/irrisdev/go-shorten/models"
	"log"
)

func QueryAll() *[]models.URL {

	db, err := openDB()
	if err != nil {
		log.Println(err)
		return nil
	}
	defer db.Close()

	query := `
    SELECT id, OriginalURL, ShortURL, Creation, Expiration
    FROM URLS
    LIMIT 10;`

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	var urls []models.URL
	for rows.Next() {
		var url models.URL
		err := rows.Scan(&url.OriginalURL, &url.ShortURL, &url.Creation, &url.Expiration)
		if err != nil {
			log.Println(err)
			return nil
		}
		urls = append(urls, url)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil
	}

	return &urls
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to open database:", err)
		return nil, err
	}

	return db, nil
}
