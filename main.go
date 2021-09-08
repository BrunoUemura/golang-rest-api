package main

import (
	"github.com/BrunoUemura/golang-rest-api/src/database"
	"github.com/BrunoUemura/golang-rest-api/src/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}