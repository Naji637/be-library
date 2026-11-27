package borrowing

import (
	"be-library/app/model"
	"be-library/app/service/book"
	"be-library/app/service/book_hub"
	"be-library/app/service/member"
	"errors"
	"fmt"
	"log"
	"time"
)

type Service interface {
	CreateBorrowing(borrowing RequestBorrowing) error
	ReturnBorrowing(borrowingReturn RequestReturnBorrowing) ([]string, error)
	// GetAllBorrowing() []ResponseBorrowing
	GetBorrowingsByMembershipNumber(membershipNumber string) (ResponseGetMemberBorrowingBooks, error)
}

type service_ struct {
	repository  Repository
	memberRepo  member.Repository
	bookHubRepo book_hub.Repository
	bookRepo    book.Repository
}

func NewService(repository Repository, memberRepo member.Repository, bookHubRepo book_hub.Repository, bookRepo book.Repository) *service_ {
	return &service_{repository, memberRepo, bookHubRepo, bookRepo}
}
func (s *service_) CreateBorrowing(borrowing RequestBorrowing) error {
	bookHubIds, err := uniqueBookHubId(borrowing.Books)
	if err != nil {
		return err
	}
	memberClassification, err := s.memberRepo.JoinMembersToMemberClassifications(borrowing.MemberId)

	if err != nil {
		return ErrorMemberNotFound
	}
	// fmt.Println("member class", memberClassification)
	if memberClassification.MembershipNumber == "" {
		return ErrorMemberNotFound
	}
	return s.repository.Transaction(func(repo Repository) error {
		activeBookIds, err := repo.GetActiveBorrowedBookIdsByMemberId(borrowing.MemberId)

		if err != nil {
			return ErrorBookAlreadyBorrowed
		}
		if len(activeBookIds)+len(bookHubIds) > memberClassification.MaxBooks {
			return fmt.Errorf("%w maksimal pinjam %v sudah pinjam %v ", ErrorMaxBooksExceeded, memberClassification.MaxBooks, len(activeBookIds))
		}
		heldBookIds := make(map[string]struct{}, len(activeBookIds))
		for _, bookId := range activeBookIds {
			heldBookIds[bookId] = struct{}{}
		}
		bookHubs, err := repo.GetBookHubForBorrowing(bookHubIds)
		// log.Println(bookHubs)
		// log.Println(len(bookHubs))
		if err != nil || len(bookHubs) == 0 {
			// log.Println("tes")
			return ErrorBookHubNotFound
		}
		bookHubById := make(map[string]BookHubForBorrowing, len(bookHubs))
		for _, bookHub := range bookHubs {
			bookHubById[bookHub.ID] = bookHub
		}

		timeNow := time.Now().UTC()
		requestBookIds := make(map[string]struct{}, len(bookHubIds))
		borrowings := make([]*model.Borrowing, 0, len(bookHubIds))
		for _, bookHubId := range bookHubIds {
			bookHub, found := bookHubById[bookHubId]
			if !found {
				return ErrorBookHubNotFound
			}
			if bookHub.IsBorrowed {
				return ErrorBookAlreadyBorrowed
			}
			if _, held := heldBookIds[bookHub.BookId]; held {
				return ErrorBookAlreadyHeldByMember
			}
			if _, duplicate := requestBookIds[bookHub.BookId]; duplicate {
				return ErrorDuplicateBookHubId
			}
			requestBookIds[bookHub.BookId] = struct{}{}
			borrowings = append(borrowings, &model.Borrowing{
				MemberID:          borrowing.MemberId,
				BookHubID:         bookHub.ID,
				BorrowDate:        timeNow,
				BorrowConditionId: bookHub.ConditionId,
				Status:            StatusBorrowed,
			})
		}
		if len(borrowings) != 0 {
			err := repo.CreateBorrowing(borrowings)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *service_) GetBorrowingsByMembershipNumber(membershipNumber string) (ResponseGetMemberBorrowingBooks, error) {
	return s.repository.GetBorrowedBooksByMembershipNumber(membershipNumber)
}

func (s *service_) ReturnBorrowing(borrowingReturn RequestReturnBorrowing) ([]string, error) {
	activeBookIdsByMember, err := s.repository.GetBookHubIdsByMemberId(borrowingReturn.MemberId)
	if err != nil {
		return nil, err
	}

	existingBooks := make(map[string]struct{}, len(activeBookIdsByMember))
	for _, activeBook := range activeBookIdsByMember {
		existingBooks[activeBook.BookHubID] = struct{}{}
	}
	message := []string{}
	return message, s.repository.TransactionReturnBookHub(func(repo Repository, repoBookHub book_hub.Repository) error {
		timeNow := time.Now()
		log.Println(existingBooks)
		for _, requestBorrowing := range borrowingReturn.ReturnBookHubs {
			_, found := existingBooks[requestBorrowing.BookHubId]
			log.Println(found)
			log.Println("booookkkhubbb", requestBorrowing.BookHubId)
			if found {
				err := repo.ReturnBorrowing(requestBorrowing.BookHubId, &model.Borrowing{
					ReturnDate:        &timeNow,
					ReturnConditionId: &requestBorrowing.ReturnConditionId,
					Status:            StatusReturned,
				})
				if err != nil {
					return err
				}
				err = repoBookHub.UpdateBookHubCondition(requestBorrowing.BookHubId, requestBorrowing.ReturnConditionId)
				if err != nil {
					return err
				}
			} else {
				log.Println("apapap")
				message = append(message, fmt.Sprintf("%v tidak dipinjam", requestBorrowing.BookHubId))
			}

		}
		return nil
	})

}

// 	returnBookHubs, err := validateReturnBookHubBorrowing(borrowingReturn)
// 	log.Println("return book HUBsss", returnBookHubs)
// 	returnBookHubIds := bookHubIdArray(borrowingReturn.ReturnBookHubs)
// 	log.Println("boookHUbbsss idddss", returnBookHubIds)
// 	if err != nil {
// 		return errors.New("gagal validasi return buku")
// 	}
// 	return s.repository.TransactionReturnBookHub(func(repo Repository, repoBookHub book_hub.Repository) error {
// 		// timeNow := time.Now()
// 		for _,requestBorrowing:=range borrowingReturn.ReturnBookHubs{
// 			if existingBooks[requestBorrowing.BookHubId]{
// 				err:=repo.
// 			}
//
//
// 		}
// 		var returnBorrowings []*model.Borrowing
// 		for _, returnBookHubId := range returnBookHubIds {
// 			returnConditionId := returnBookHubs[returnBookHubId]
// 			returnBorrowings = append(returnBorrowings, &model.Borrowing{
// 				ReturnDate:        &timeNow,
// 				ReturnConditionId: &returnConditionId,
// 				Status:            StatusReturned,
// 			})
// 		}
// 		log.Println("returnnnnnn borrooooowingss", returnBorrowings)
// 		err := repo.ReturnBorrowing(returnBookHubIds, returnBorrowings)
// 		if err != nil {
// 			return errors.New("gagal return buku")
// 		}
// 	return nil
// })

// }

func bookHubIdArray(borrowingReturn []ReturnBookHub) []string {
	if len(borrowingReturn) == 0 {
		return nil
	}
	var bookHubIds []string
	for _, bookHubId := range borrowingReturn {
		bookHubIds = append(bookHubIds, bookHubId.BookHubId)
	}
	return bookHubIds
}

func validateReturnBookHubBorrowing(returnBookHubs []ReturnBookHub) (map[string]string, error) {
	if len(returnBookHubs) == 0 {
		return nil, errors.New("tidak ada buku yang direturn")
	}
	bookHubs := make(map[string]string, len(returnBookHubs))
	for _, bookHub := range returnBookHubs {
		bookHubs[bookHub.BookHubId] = bookHub.ReturnConditionId
	}
	return bookHubs, nil
}

// func (s *service_) GetAllBorrowing() []ResponseBorrowing {
// 	borrowings := s.repository.GetAllBorrowing()
// 	return borrowings
// }

func uniqueBookHubId(books []Book) ([]string, error) {

	if len(books) == 0 {
		return nil, ErrorNoBookRequested
	}
	bookHubIds := make([]string, 0, len(books))
	seen := make(map[string]struct{}, len(books))
	for _, book := range books {
		if book.BookHubId == "" {
			return nil, ErrorEmptyBookHubId
		}
		if _, duplicate := seen[book.BookHubId]; duplicate {
			return nil, ErrorDuplicateBookHubId
		}
		seen[book.BookHubId] = struct{}{}
		bookHubIds = append(bookHubIds, book.BookHubId)
	}
	// fmt.Println("====")
	// fmt.Println(bookHubIds)
	// fmt.Println(seen)
	// fmt.Println("====")
	return bookHubIds, nil
}
