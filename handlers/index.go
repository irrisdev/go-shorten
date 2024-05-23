package handlers

import (
	"github.com/irrisdev/go-shorten/utils"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	} else if r.URL.Path != "/" {
		if value, exist := utils.CheckExist(r.URL.Path); exist {
			http.Redirect(w, r, value, http.StatusPermanentRedirect)
		} else {
			http.NotFound(w, r)
		}
		return
	}

	someExamples := utils.QueryAll(10)

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	if err := tmpl.Execute(w, someExamples); err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}

}
