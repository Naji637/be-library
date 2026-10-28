package book

import "be-library/app/model"

type Service interface {
	CreateBook(book RequestBook) error
	GetAllBooks() []ResponseBook
	UpdateBook(book RequestBook, id string) error
	DeleteBook(id string) error
}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}
func (s *service_) CreateBook(book RequestBook) error {
	return s.repository.CreateBook(&model.Book{Title: book.Title, CategoryId: book.CategoryId, Stock: book.Stock, Author: book.Author})
}
func (s *service_) GetAllBooks() []ResponseBook {
	books := s.repository.GetAllBooks()
	return books
}
func (s *service_) UpdateBook(book RequestBook, id string) error {
	return s.repository.UpdateBook(&model.Book{Title: book.Title, CategoryId: book.CategoryId, Stock: book.Stock, Author: book.Author}, id)
}
func (s *service_) DeleteBook(id string) error {
	return s.repository.DeleteBook(id)
}
