package main

import (
	"fmt"
	"github.com/irrisdev/go-shorten/config"
	"github.com/irrisdev/go-shorten/routes"
	"github.com/irrisdev/go-shorten/templates"
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

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	//Initialise Routes
	router := routes.InitRouter()

	//Initialise Templates
	templates.InitTemplates()

	//Start Server
	log.Println("Started on port", port)
	fmt.Println("To close connection CTRL+C")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("Error starting server ", err)
	}

}
