package config

import (
	"github.com/joho/godotenv"
	"log"
)

// The purpose of this function is to load the $env variables and make it easy to change them

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
}
