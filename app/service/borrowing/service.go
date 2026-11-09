package borrowing

import (
	"be-library/app/model"
	"be-library/app/service/member"
	"fmt"
)

type Service interface {
	CreateBorrowing(borrowing RequestBorrowing) error
	// GetAllBorrowing() []ResponseBorrowing
	GetBorrowingsByMembershipNumber(membershipNumber string) *ResponseBorrowing
}

type service_ struct {
	repository Repository
	memberRepo member.Repository
}

func NewService(repository Repository, memberRepo member.Repository) *service_ {
	return &service_{repository, memberRepo}
}
func (s *service_) CreateBorrowing(borrowing RequestBorrowing) error {
	err := s.repository.CreateBorrowing(&model.Borrowing{MemberID: borrowing.MemberId, BookHubID: borrowing.BookHubId, BorrowDate: borrowing.BorrowDate, ReturnDate: borrowing.ReturnDate, BorrowOfficerId: borrowing.BorrowOfficerId, ReturnOfficerId: borrowing.ReturnOfficerId, BorrowConditionId: borrowing.BorrowConditionId, ReturnConditionId: borrowing.ReturnConditionId, Status: borrowing.Status})
	return err
}

func (s *service_) GetBorrowingsByMembershipNumber(membershipNumber string) *ResponseBorrowing {
	data, err := s.memberRepo.JoinMembersToMemberClassifications(membershipNumber)
	fmt.Println(data, err)
	return &ResponseBorrowing{}
}

// func (s *service_) GetAllBorrowing() []ResponseBorrowing {
// 	borrowings := s.repository.GetAllBorrowing()
// 	return borrowings
// }
