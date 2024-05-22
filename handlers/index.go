package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "Index Page")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
