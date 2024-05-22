package handlers

import (
	"html/template"
	"io"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	io.WriteString(w, r.Method)

	tmpl := template.Must(template.ParseFiles("index.html"))

	tmpl.Execute(w, nil)
}
