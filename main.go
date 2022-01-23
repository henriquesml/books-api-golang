package main

import "github.com/henriquesml/go-api/server"

func main() {
	server := server.NewServer()
	server.Run()
}
