package handlers

import (
	"io"
	"net/http"
)

func CreateURL(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	io.WriteString(w, r.Method)

}
