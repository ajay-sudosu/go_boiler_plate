package main

import (
	"abc/internal/server"
	"log"
)

// @title        Sample Code to start with go using echo
// @version      1.0
// @description  This is a sample server.
// @host         localhost:8080
// @BasePath     /api
func main() {
	if err := server.InitServer(); err != nil {
		log.Fatal(err)
	}
}
