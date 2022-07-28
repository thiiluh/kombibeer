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

func (b *beerUseCase) Create(beer beers.Beer) (beers.Beer, error) {
	return b.repo.Create(beer)
}

func (b *beerUseCase) FindOne(id int) (beers.Beer, error) {
	return b.repo.FindOne(id)
}

func (b *beerUseCase) Update(id int, beer beers.Beer) (beers.Beer, error) {
	return b.repo.Update(id, beer)
}

func (b *beerUseCase) UpdateAllFields(id int, beer beers.Beer) (beers.Beer, error) {
	beer.Id = id
	return b.repo.UpdateAllFields(beer)
}

func (b *beerUseCase) Remove(id int) error {
	return b.repo.Remove(id)
}
