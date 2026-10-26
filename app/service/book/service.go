package book

import "be-library/app/model"

type Service interface {
	CreateBook(book RequestBook) error
	GetAllBooks() []ResponseBook
}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}
func (s *service_) CreateBook(book RequestBook) error {
	err := s.repository.CreateBook(&model.Book{Title: book.Title, CategoryId: book.CategoryId, Stock: book.Stock, Author: book.Author})
	return err
}
func (s *service_) GetAllBooks() []ResponseBook {
	books := s.repository.GetAllBooks()
	return books
}
