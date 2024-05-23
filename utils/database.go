package utils

import (
	"database/sql"
	"errors"
	"github.com/irrisdev/go-shorten/models"
	"log"
	"os"
	"strings"
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

func CheckExist(path string) (string, bool) {

	s := strings.Split(path, "/")

	if len(s) > 3 {
		return "", false
	}

	code := s[len(s)-1]

	db, err := openDB()
	if err != nil {
		log.Println(err)
		return "", false
	}
	defer db.Close()

	var url string

	if err := db.QueryRow("SELECT OriginalURL FROM URLS WHERE ShortURL = ?", code).Scan(&url); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(err)
			return "", false
		}
		log.Println(err)
		return "", false
	}

	go click(code)

	return url, true
}

func click(code string) {

	db, err := openDB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	result, err := db.Exec("UPDATE URLS SET Clicks = Clicks + 1 WHERE ShortURL = ?", code)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)

}
