package category

import "be-library/app/model"

type Service interface {
	CreateCategory(category RequestCategory) error
	GetAllCategories() []ResponseCategory
}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}

func (s *service_) CreateCategory(category RequestCategory) error {
	err := s.repository.CreateCategory(&model.Category{Category: category.Category})
	return err
}
func (s *service_) GetAllCategories() []ResponseCategory {
	categories := s.repository.GetAllCategories()
	return categories
}