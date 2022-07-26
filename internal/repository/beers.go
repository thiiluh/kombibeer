package repository

import (
	"github.com/thiiluh/kombibeer/internal/config"
	"github.com/thiiluh/kombibeer/pkg/beers"
)

type beerRepository struct{}

func NewBeerRepository() (repo *beerRepository) {
	return &beerRepository{}
}

func (b *beerRepository) FindAll() (beers []beers.Beer, err error) {
	result := config.DB.Find(&beers)
	err = result.Error

	return beers, err
}
