package main

import (
	"fmt"
	"github.com/irrisdev/go-shorten/config"
	"github.com/irrisdev/go-shorten/models"
	"github.com/irrisdev/go-shorten/routes"
	"github.com/irrisdev/go-shorten/utils"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {

	//Load env variables
	config.LoadConfig()

	if err := utils.InitDB(); err != nil {
		log.Fatal(err)
	}

	s := models.Server{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}

	//Initialise Routes
	router := routes.InitRouter()

	srv := &http.Server{
		Addr:         net.JoinHostPort(s.Host, s.Port),
		Handler:      router,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	//Start Server
	log.Println("Started on port", s.Port)
	fmt.Println("To close connection CTRL+C")

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server ", err)
	}

}
