package main

import (
	"abc/internal/server"
	"log"
)

// @title        E-Commerce API
// @version      1.0
// @description  This is a sample server for an e-commerce system.
// @host         localhost:8080
// @BasePath     /api
func main() {
	if err := server.InitServer(); err != nil {
		log.Fatal(err)
	}
}
