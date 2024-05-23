package templates

import (
	"fmt"
	"github.com/irrisdev/go-shorten/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var tmpl *template.Template

func InitTemplates() {
	layout := filepath.Join("templates", "layout.html")

	tmpl = template.Must(template.New("layout").ParseFiles(layout))

	tmpl = template.Must(tmpl.New("indexForm").Parse(models.IndexForm))

	tmpl = template.Must(tmpl.New("indexHX").Parse(models.IndexHX))

}

func ResultHandler(w http.ResponseWriter, completeURL string) {

	result := fmt.Sprintf(models.ResultForm, completeURL, completeURL)

	tmpl, _ := template.New("result").Parse(result)

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
	}

}

func IndexHandler(w http.ResponseWriter) {
	if err := tmpl.ExecuteTemplate(w, "layout", "indexForm"); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
	}
}

func IndexHX(w http.ResponseWriter) {

	if err := tmpl.Execute(w, "indexHX"); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
	}
}
