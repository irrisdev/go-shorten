package routes

import (
	"github.com/irrisdev/go-shorten/handlers"
	"net/http"
)

func InitRouter() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/shorten/", handlers.CreateURL)

	return mux
}
