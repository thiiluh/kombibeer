package config

import (
	"fmt"
	"log"

	"github.com/thiiluh/kombibeer/pkg/beers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(host string, port int) {

	dsn := fmt.Sprintf("host=%s user=postgres password=postgres dbname=postgres port=%d", host, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&beers.Beer{})
	DB = db
}
