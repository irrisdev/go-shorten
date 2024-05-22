package main

import (
	"fmt"
	"github.com/irrisdev/go-shorten/config"
	"github.com/irrisdev/go-shorten/models"
	"github.com/irrisdev/go-shorten/routes"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {

	//Load env variables
	config.LoadConfig()

	s := models.Server{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}

	//Initialise Routes
	router := routes.InitRouter()

	//Start Server
	log.Println("Started on port", s.Port)
	fmt.Println("To close connection CTRL+C")

	err := http.ListenAndServe(net.JoinHostPort(s.Host, s.Port), router)
	if err != nil {
		log.Fatal("Error starting server ", err)
	}

}
