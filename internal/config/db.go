package config

import (
	"log"

	"github.com/thiiluh/kombibeer/pkg/beers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("beer.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&beers.Beer{})
	DB = db
}
