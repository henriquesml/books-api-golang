package main

import (
	"github.com/henriquesml/go-api/database"
	"github.com/henriquesml/go-api/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
