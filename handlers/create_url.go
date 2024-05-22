package handlers

import (
	"fmt"
	"net/http"
)

func CreateURL(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		fmt.Fprint(w, "Creating URL")
		break
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}
