package book_hub

import "be-library/app/model"

type Service interface {
	CreateBookHub(bookHub RequestBookHub) error
	GetAllBookHubs() []ResponseBookHub
}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}

func (s *service_) CreateBookHub(bookHub RequestBookHub) error {
	err := s.repository.CreateBookHub(&model.BookHub{Barcode: bookHub.Barcode, BookID: bookHub.BookId, ConditionID: bookHub.ConditionId, MemberID: bookHub.MemberId, ShelfId: bookHub.ShelfId, Price: bookHub.Price})
	return err
}

func (s *service_) GetAllBookHubs() []ResponseBookHub {
	bookHubs := s.repository.GetAllBookHubs()
	return bookHubs
}
