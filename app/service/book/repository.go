package book

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateBook(book *model.Book) error
	GetAllBooks() []ResponseBook
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateBook(book *model.Book) error {
	err := r.db.Create(book).Error
	return err
}

func (r *repository) GetAllBooks() []ResponseBook {
	var books []ResponseBook
	r.db.Find(&books)
	return books
}	