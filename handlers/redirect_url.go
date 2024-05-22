package handlers

import (
	"io"
	"net/http"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	io.WriteString(w, r.Method)

}
