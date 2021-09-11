package main

import (
	"log"

	"github.com/BrunoUemura/golang-rest-api/src/database"
	"github.com/BrunoUemura/golang-rest-api/src/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.StartDB()
	server := server.NewServer()
	server.Run()
}