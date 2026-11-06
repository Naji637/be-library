package book

import (
	"be-library/app/model"
	"be-library/app/service/book_hub"
	util "be-library/app/util/barcode"
	"errors"
	"fmt"
)

type Service interface {
	CreateBook(book RequestBook) error
	GetAllBooks() []ResponseBook
	UpdateBook(book RequestBook, id string) error
	DeleteBook(id string) error
	CreateBarcode(requestBarcode RequestBarcode) error
}

type service_ struct {
	repository        Repository
	barcode           util.Barcode
	repositoryBookHub book_hub.Repository
}

func NewService(repository Repository, barcode util.Barcode, repositoryBookHub book_hub.Repository) *service_ {
	return &service_{repository, barcode, repositoryBookHub}
}

func (s *service_) CreateBook(book RequestBook) error {
	return s.repository.CreateBook(&model.Book{Title: book.Title, CategoryId: book.CategoryId, Author: book.Author})
}
func (s *service_) GetAllBooks() []ResponseBook {
	books := s.repository.GetAllBooks()
	return books
}
func (s *service_) UpdateBook(book RequestBook, id string) error {
	return s.repository.UpdateBook(&model.Book{Title: book.Title, CategoryId: book.CategoryId, Author: book.Author}, id)
}
func (s *service_) DeleteBook(id string) error {
	return s.repository.DeleteBook(id)
}
func (s *service_) CreateBarcode(requestBarcode RequestBarcode) error {
	barcode := s.barcode.CreateBarcode()
	fmt.Println("barcode", barcode)
	var err error
	if barcode != "" {
		if err = s.repositoryBookHub.CreateBookHub(&model.BookHub{
			Barcode:     barcode,
			BookID:      requestBarcode.BookId,
			ConditionID: requestBarcode.ConditionId,
			ShelfId:     requestBarcode.ShelfId,
			MemberID:    requestBarcode.MemberId,
			Price:       requestBarcode.Price,
		}); err == nil {
			book := s.repository.GetBookById(requestBarcode.BookId)
			fmt.Println("book", book)
			if err = s.repository.UpdateStock(requestBarcode.BookId, book.Stock+1); err == nil {
				return nil
			}
		}
		fmt.Println("err:", err)

	}
	return errors.New("error create barcode")
}
