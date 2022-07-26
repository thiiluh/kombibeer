package main

import (
	"log"

	"github.com/thiiluh/kombibeer/api/routes"
	"github.com/thiiluh/kombibeer/internal/config"
)

func main() {
	config.Connect()
	if err := routes.InitRoutes(); err != nil {
		log.Fatal(err)
	}
}
