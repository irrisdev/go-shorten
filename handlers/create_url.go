package handlers

import (
	"fmt"
	"github.com/irrisdev/go-shorten/models"
	"github.com/irrisdev/go-shorten/utils"
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"
	"time"
)

func CreateURL(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost && (r.Header.Get("HX-Request")) == "true" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

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
		Scheme: r.URL.Scheme,
		Host:   r.Host,
		Path:   entry.ShortURL,
	}

	completeURL := u.String()
	htmlStr := fmt.Sprintf(`<label for="urlShortened">Shortened URL</label>
    <div class="input-group" id="result" ><input type="text" value="%s" name="urlShortened" class="form-control" id="urlOut" disabled></div><a href="%s" target="_blank"><button class="btn btn-primary btn-block" style="margin-top: 0.5rem">Visit</button></a>
	<button class="btn btn-primary btn-block" style="margin-top: 0.5rem" onclick="copyToClipboard()">Copy</button>`, completeURL, completeURL)

	tmpl, _ := template.New("t").Parse(htmlStr)

	tmpl.Execute(w, nil)

}
