package shelf

import "be-library/app/model"

type Service interface {
	CreateShelf(shelf RequestShelf) error
	GetAllShelves() []ResponseShelf

}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}

func (s *service_) CreateShelf(shelf RequestShelf) error {
	err := s.repository.CreateShelves(&model.Shelf{ShelfNumber: shelf.ShelfNumber})
	return err
}

func (s *service_) GetAllShelves() []ResponseShelf{
	shelves := s.repository.GetAllShelves()
	return shelves
}