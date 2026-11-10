package book_hub

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateBookHub(bookHub *model.BookHub) error
	GetAllBookHubs() []ResponseBookHub
	GetBookHubById(id string) (model.BookHub, error)
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

func (r *repository) GetBookHubById(id string) (model.BookHub, error) {
	var bookHub model.BookHub
	err := r.db.Where("id=?", id).Find(&bookHub).Error
	return bookHub, err
}

// func (r *repository) GetBookHubByBarcode(barcode string) model.BookHub {
// 	var bookHub ResponseBookHub
// }
