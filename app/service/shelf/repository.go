package shelf

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateShelves(shelf *model.Shelf) error
	GetAllShelves() []ResponseShelf
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateShelves(shelf *model.Shelf) error {
	err := r.db.Create(shelf).Error
	return err
}

func (r *repository) GetAllShelves() []ResponseShelf {
	var shelves []ResponseShelf
	r.db.Find(&shelves)
	return shelves
}
