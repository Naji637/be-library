package book

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateBook(book *model.Book) error
	GetAllBooks() []ResponseBook
	UpdateBook(book *model.Book, id string) error
	DeleteBook(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateBook(book *model.Book) error {
	return r.db.Create(book).Error
}

func (r *repository) GetAllBooks() []ResponseBook {
	var books []ResponseBook
	r.db.Debug().Find(&books)
	return books
}

func (r *repository) UpdateBook(book *model.Book, id string) error {
	return r.db.Model(&model.Book{}).Where("id=?", id).Updates(model.Book{Title: book.Title, CategoryId: book.CategoryId, Stock: book.Stock, Author: book.Author}).Error
}

func (r *repository) DeleteBook(id string) error {
	
	return r.db.Debug().Where("id=?", id).Delete(&model.Book{}).Error
}
