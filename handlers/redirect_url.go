package handlers

import (
	"fmt"
	"net/http"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "Redirecting")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
