package main

import (
	"fmt"
	"github.com/irrisdev/go-shorten/config"
	"github.com/irrisdev/go-shorten/models"
	"github.com/irrisdev/go-shorten/routes"
	"github.com/irrisdev/go-shorten/utils"
	"log"
	"net/http"
	"os"
)

func main() {

	//Load env variables
	config.LoadConfig()

	if err := utils.InitDB(); err != nil {
		log.Fatal(err)
	}

	var host, port string
	if host = os.Getenv("HOST"); host == "" {
		host = "0.0.0.0"
	}
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	s := models.Server{
		Host: host,
		Port: port,
	}

	//Initialise Routes
	router := routes.InitRouter()

	//Start Server
	log.Println("Started on port", s.Port)
	fmt.Println("To close connection CTRL+C")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("Error starting server ", err)
	}

}
