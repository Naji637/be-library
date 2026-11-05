package borrowing

import "be-library/app/model"

type Service interface {
	CreateBorrowing(borrowing RequestBorrowing) error
	GetAllBorrowing() []ResponseBorrowing
}

type service_ struct {
	repository Repository
}

func NewService(repository Repository) *service_ {
	return &service_{repository}
}
func (s *service_) CreateBorrowing(borrowing RequestBorrowing) error {
	err := s.repository.CreateBorrowing(&model.Borrowing{MemberID: borrowing.MemberId, BookHubID: borrowing.BookHubId, BorrowDate: borrowing.BorrowDate, ReturnDate: borrowing.ReturnDate, BorrowOfficerId: borrowing.BorrowOfficerId, ReturnOfficerId: borrowing.ReturnOfficerId, BorrowConditionId: borrowing.BorrowConditionId, ReturnConditionId: borrowing.ReturnConditionId, Status: borrowing.Status})
	return err
}
func (s *service_) GetAllBorrowing() []ResponseBorrowing {
	borrowings := s.repository.GetAllBorrowing()
	return borrowings
}
