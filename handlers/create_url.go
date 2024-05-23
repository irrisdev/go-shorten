package handlers

import (
	"github.com/irrisdev/go-shorten/models"
	"github.com/irrisdev/go-shorten/templates"
	"github.com/irrisdev/go-shorten/utils"
	"log"
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

	inputUrl := r.PostFormValue("url")

	//Appends SCHEMA to url if doesn't include
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
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	u := &url.URL{
		Scheme: os.Getenv("SCHEME"),
		Host:   r.Host,
		Path:   entry.ShortURL,
	}

	completeURL := u.String()

	templates.ResultHandler(w, completeURL)

}
