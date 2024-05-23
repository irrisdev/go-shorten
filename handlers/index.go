package handlers

import (
	"github.com/irrisdev/go-shorten/templates"
	"github.com/irrisdev/go-shorten/utils"
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

	//someExamples := utils.QueryAll(10)

	if r.Header.Get("HX-Request") == "true" {

		templates.IndexHX(w)

	} else {

		templates.IndexHandler(w)

	}

}
