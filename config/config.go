package config

import (
	"github.com/joho/godotenv"
	"log"
)

var Server	*server
var MySQL	*mySQL

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	MySQL = setupMySQL()
	Server = setupServer()
}