package config

import (
	"log"

	"github.com/thiiluh/kombibeer/pkg/beers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&beers.Beer{})
	DB = db
}
