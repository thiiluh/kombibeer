package interface_repository

import "github.com/thiiluh/kombibeer/pkg/beers"

type BeerRepository interface {
	FindAll() (beers []beers.Beer, err error)
	Create(beer beers.Beer) (beers.Beer, error)
	FindOne(id int) (beers.Beer, error)
	Update(id int, beer beers.Beer) (beers.Beer, error)
	UpdateAllFields(beer beers.Beer) (beers.Beer, error)
	Remove(id int) error
}
