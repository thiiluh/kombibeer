package interface_repository

import "github.com/thiiluh/kombibeer/pkg/beers"

type BeerRepository interface {
	FindAll() (beers []beers.Beer, err error)
}
