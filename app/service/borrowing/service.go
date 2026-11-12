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
	ReturnBorrowing(borrowingReturn []RequestReturnBorrowing) error
	// GetAllBorrowing() []ResponseBorrowing
	GetBorrowingsByMembershipNumber(membershipNumber string) (ResponseGetMemberBorrowingBooks, error)
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
	memberClassification, _ := s.memberRepo.JoinMembersToMemberClassifications(borrowing.MemberId)
	if len(borrowing.Books) <= memberClassification.MaxBooks {
		borrowedBooks, _ := s.GetBorrowingsByMembershipNumber(memberClassification.MembershipNumber)
		for _, borrowingBooks := range borrowing.Books {
			bookHub, err := s.bookHubRepo.GetBookHubById(borrowingBooks.BookHubId)
			for _, responseBooks := range borrowedBooks.Borrowing {
				if responseBooks.BookHub.BookId == bookHub.BookID {
					return errors.New("buku yang dipinjam tidak boleh sama")
				}
			}
			if err != nil {
				log.Println(err)
				return errors.New("ups")
			}
			err = s.repository.CreateBorrowing(&model.Borrowing{MemberID: borrowing.MemberId, BookHubID: borrowingBooks.BookHubId, BorrowDate: time.Now(), BorrowConditionId: bookHub.ConditionID, Status: "dipinjam"})
			if err != nil {
				log.Println(err)
				return errors.New("ups")
			}
		}

	} else {
		return errors.New("melebihi batas pinjam buku")
	}

	return nil
}

func (s *service_) GetBorrowingsByMembershipNumber(membershipNumber string) (ResponseGetMemberBorrowingBooks, error) {
	return s.repository.GetBorrowedBooksByMembershipNumber(membershipNumber)
}

func (s *service_) ReturnBorrowing(borrowingReturn []RequestReturnBorrowing) error {
	timeNow := time.Now()
	for _, returnBorrowing := range borrowingReturn {
		err := s.repository.ReturnBorrowing(returnBorrowing.ID, &model.Borrowing{ReturnDate: &timeNow, ReturnConditionId: &returnBorrowing.ReturnConditionId, Status: "dikembalikan"})
		if err != nil {
			return errors.New("error return borrowing")
		}
		err = s.bookHubRepo.UpdateBookHubCondition(returnBorrowing.BookHubId, returnBorrowing.ReturnConditionId)
		if err != nil {
			return errors.New("error update book hub condition")
		}
	}
	return nil
}

// func (s *service_) GetAllBorrowing() []ResponseBorrowing {
// 	borrowings := s.repository.GetAllBorrowing()
// 	return borrowings
// }
