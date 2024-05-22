package utils

import (
	"database/sql"
	"github.com/irrisdev/go-shorten/models"
	"log"
	"os"
)

func QueryAll(n int) *[]models.URL {

	db, err := openDB()
	if err != nil {
		log.Println(err)
		return nil
	}
	defer db.Close()

	rows, err := db.Query("SELECT OriginalURL, ShortURL FROM URLS LIMIT ?", n)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	var urls []models.URL

	for rows.Next() {
		var url models.URL
		if err := rows.Scan(&url.OriginalURL, &url.ShortURL); err != nil {
			return &urls
		}
		urls = append(urls, url)
	}
	if err = rows.Err(); err != nil {
		return &urls
	}

	return &urls
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DATABASE_PATH"))
	if err != nil {
		log.Fatal("Failed to open database:", err)
		return nil, err
	}

	return db, nil
}

func InsertEntry(url *models.URL) error {

	db, err := openDB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO URLS VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(url.OriginalURL, url.ShortURL, url.Clicks, url.Creation, url.Expiration)
	if err != nil {
		log.Println(err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
