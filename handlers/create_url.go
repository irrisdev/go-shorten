package handlers

import (
	"net/http"
)

func CreateURL(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		break
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}
