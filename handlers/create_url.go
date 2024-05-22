package handlers

import (
	"github.com/irrisdev/go-shorten/models"
	"github.com/irrisdev/go-shorten/utils"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func CreateURL(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost && (r.Header.Get("HX-Request")) == "true" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	log.Println("HTMX Request Received")
	log.Println(r.Header.Get("HX-Request"))

	inputUrl := r.PostFormValue("url")

	if !strings.HasPrefix(inputUrl, "http://") && !strings.HasPrefix(inputUrl, "https://") {
		inputUrl = "https://" + inputUrl
	}
	_, err := url.ParseRequestURI(inputUrl)
	if err != nil {
		log.Println("Invalid URL", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	currentTime := time.Now()

	entry := models.URL{
		OriginalURL: inputUrl,
		ShortURL:    utils.GenerateCode(inputUrl),
		Clicks:      0,
		Creation:    currentTime.Unix(),
		Expiration:  currentTime.Add(time.Hour * 24 * 7).Unix(),
	}

	if err := utils.InsertEntry(&entry); err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	u := &url.URL{
		Scheme: "https",
		Host:   net.JoinHostPort(os.Getenv("HOST"), os.Getenv("PORT")),
		Path:   entry.ShortURL,
	}

	completeURL := u.String()
	log.Println(completeURL)
}
