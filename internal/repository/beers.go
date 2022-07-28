package repository

import (
	"errors"

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

func (b *beerRepository) Create(beer beers.Beer) (beers.Beer, error) {
	result := config.DB.Create(&beer)

	return beer, result.Error
}

func (b *beerRepository) FindOne(id int) (beers.Beer, error) {
	beer := beers.Beer{}
	result := config.DB.First(&beer, id)

	return beer, result.Error
}

func (b *beerRepository) UpdateAllFields(beer beers.Beer) (beers.Beer, error) {
	result := config.DB.Updates(&beer)
	if result.RowsAffected == 0 {
		return beer, errors.New("resource not found")
	}

	return beer, result.Error
}

func (b *beerRepository) Update(id int, beer beers.Beer) (beers.Beer, error) {
	bb := beers.Beer{}
	result := config.DB.Model(&bb).Where("Id = ?", id).UpdateColumns(&beer)
	if result.RowsAffected == 0 {
		return beer, errors.New("resource not found")
	}

	return beer, result.Error
}

func (b *beerRepository) Remove(id int) error {
	bb := beers.Beer{}
	result := config.DB.Delete(bb, id)
	if result.RowsAffected == 0 {
		return errors.New("resource not found")
	}

	return result.Error
}
