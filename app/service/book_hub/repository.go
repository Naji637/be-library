package book_hub

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateBookHub(bookHub *model.BookHub) error
	GetAllBookHubs() []ResponseBookHub
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateBookHub(bookHub *model.BookHub) error {
	err := r.db.Create(bookHub).Error
	return err
}

func (r *repository) GetAllBookHubs() []ResponseBookHub {
	var bookHub []ResponseBookHub
	r.db.Find(&bookHub)
	return bookHub
}
