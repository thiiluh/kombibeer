package usecase

import (
	"github.com/thiiluh/kombibeer/pkg/beers"
	interface_repository "github.com/thiiluh/kombibeer/pkg/beers/interface"
)

type beerUseCase struct {
	repo interface_repository.BeerRepository
}

func NewBeerUseCase(repo interface_repository.BeerRepository) *beerUseCase {
	return &beerUseCase{repo}
}

func (b *beerUseCase) FindAll() (beers []beers.Beer, err error) {
	return b.repo.FindAll()
}
