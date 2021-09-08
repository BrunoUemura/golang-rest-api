package main

import (
	"github.com/BrunoUemura/golang-rest-api/database"
	"github.com/BrunoUemura/golang-rest-api/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()
	server.Run()
}