package book

import (
	"be-library/app/model"
	// "be-library/app/service/book"

	"gorm.io/gorm"
)

type Repository interface {
	CreateBook(book *model.Book) error
	GetAllBooks() []ResponseBook
	UpdateBook(book *model.Book, id string) error
	DeleteBook(id string) error
	UpdateStock(bookId string, stock int) error
	GetBookById(id string) model.Book

	// GetBookByBarcode(barcode string) model.Book
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
	r.db.Find(&books)
	return books
}

func (r *repository) UpdateBook(book *model.Book, id string) error {
	return r.db.Model(&model.Book{}).Where("id=?", id).Updates(model.Book{Title: book.Title, CategoryId: book.CategoryId, Stock: book.Stock, Author: book.Author}).Error
}

func (r *repository) DeleteBook(id string) error {

	return r.db.Where("id=?", id).Delete(&model.Book{}).Error
}
func (r *repository) UpdateStock(bookId string, stock int) error {
	return r.db.Model(&model.Book{}).Where("id=?", bookId).Updates(map[string]interface{}{"stock": stock}).Error
}
func (r *repository) GetBookById(id string) model.Book {
	var book model.Book
	r.db.Where("id=?", id).Find(&book)
	return book
}
