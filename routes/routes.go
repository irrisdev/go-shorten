package routes

import (
	"go-shorten/handlers"
	"net/http"
)

func InitRouter() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/shorten", handlers.CreateURL)
	mux.HandleFunc("/redirect/{id}", handlers.RedirectURL)

	return mux
}
