package main

import (
	"log"
	"os"

	"github.com/thiiluh/kombibeer/api/routes"
	"github.com/thiiluh/kombibeer/internal/config"
)

func main() {
	host := os.Getenv("host")

	if host == "" {
		host = "localhost"
	}
	config.Connect(host)
	if err := routes.InitRoutes(); err != nil {
		log.Fatal(err)
	}
}
