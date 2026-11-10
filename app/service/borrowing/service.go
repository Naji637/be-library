package borrowing

import (
	"be-library/app/model"
	"be-library/app/service/book_hub"
	"be-library/app/service/member"
	"errors"
	"log"
	"time"
)

type Service interface {
	CreateBorrowing(borrowing RequestBorrowing) error
	// GetAllBorrowing() []ResponseBorrowing
	// GetBorrowingsByMembershipNumber(membershipNumber string) *ResponseBorrowing
}

type service_ struct {
	repository  Repository
	memberRepo  member.Repository
	bookHubRepo book_hub.Repository
}

func NewService(repository Repository, memberRepo member.Repository, bookHubRepo book_hub.Repository) *service_ {
	return &service_{repository, memberRepo, bookHubRepo}
}
func (s *service_) CreateBorrowing(borrowing RequestBorrowing) error {
	for _, borrowingBooks := range borrowing.Books {
		condition, err := s.bookHubRepo.GetBookHubById(borrowingBooks.BookHubId)
		if err != nil {
			log.Println(err)
			return errors.New("ups")
		}
		err = s.repository.CreateBorrowing(&model.Borrowing{MemberID: borrowing.MemberId, BookHubID: borrowingBooks.BookHubId, BorrowDate: time.Now(), BorrowConditionId: condition.ConditionID, Status: "dipinjam"})
		if err != nil {
			log.Println(err)
			return errors.New("ups")
		}
	}
	return nil
}

func (s *service_) GetBorrowingsByMembershipNumber(membershipNumber string) (ResponseGetMemberBorrowingBooks, error) {
	return s.repository.GetBorrowedBooksByMemberId(membershipNumber)
}

// func (s *service_) GetAllBorrowing() []ResponseBorrowing {
// 	borrowings := s.repository.GetAllBorrowing()
// 	return borrowings
// }
