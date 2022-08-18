package main

import (
	"log"
	"os"

	"github.com/thiiluh/kombibeer/api/handlers"
	"github.com/thiiluh/kombibeer/internal/config"
)

func main() {
	host := os.Getenv("host")

	if host == "" {
		host = "localhost"
	}
	config.Connect(host, 5432)
	if err := handlers.InitRoutes(); err != nil {
		log.Fatal(err)
	}
}
