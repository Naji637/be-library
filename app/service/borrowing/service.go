package borrowing

import (
	"be-library/app/model"
	"be-library/app/service/book"
	"be-library/app/service/book_hub"
	"be-library/app/service/member"
	"errors"
	"fmt"
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
	bookRepo    book.Repository
}

func NewService(repository Repository, memberRepo member.Repository, bookHubRepo book_hub.Repository, bookRepo book.Repository) *service_ {
	return &service_{repository, memberRepo, bookHubRepo, bookRepo}
}
func (s *service_) CreateBorrowing(borrowing RequestBorrowing) error {
	bookHubIds, err := uniqueBookHubId(borrowing.Books)
	// fmt.Println(bookHubIds)
	if err != nil {
		return err
	}
	memberClassification, err := s.memberRepo.JoinMembersToMemberClassifications(borrowing.MemberId)
	if err != nil {
		return errors.New("tidak ada member id")
	}
	// fmt.Println("member class", memberClassification)
	if memberClassification.MembershipNumber == "" {
		return errors.New("member not found")
	}
	return s.repository.Transaction(func(repo Repository) error {
		activeBookIds, err := repo.GetActiveBorrowedBookIdsByMemberId(borrowing.MemberId)
		if err != nil {
			return errors.New("gagal mengambil data transaksi terakhir")
		}
		if len(activeBookIds)+len(bookHubIds) > memberClassification.MaxBooks {
			return errors.New("melebihi batas pinjam buku")
		}
		heldBookIds := make(map[string]struct{}, len(activeBookIds))
		for _, bookId := range activeBookIds {
			heldBookIds[bookId] = struct{}{}
		}
		bookHubs, err := repo.GetBookHubForBorrowing(bookHubIds)
		if err != nil {
			return errors.New("gagal get book hub")
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
				return errors.New("buku tidak ditemukan")
			}
			if bookHub.IsBorrowed {
				return errors.New("buku sudah dipinjam")
			}
			if _, held := heldBookIds[bookHub.BookId]; held {
				return errors.New("buku sudah dipinjam oleh member ini")
			}
			if _, duplicate := requestBookIds[bookHub.BookId]; duplicate {
				return errors.New("judul yg sama tidak boleh dipinjam")
			}
			requestBookIds[bookHub.BookId] = struct{}{}
			borrowings = append(borrowings, &model.Borrowing{
				MemberID:          borrowing.MemberId,
				BookHubID:         bookHub.ID,
				BorrowDate:        timeNow,
				BorrowConditionId: bookHub.ConditionId,
				Status:            "dipinjam",
			})
		}
		if len(borrowings) != 0 {
			fmt.Println(len(borrowings))
			fmt.Println("=============")
			fmt.Println(borrowings)
			err := repo.CreateBorrowing(borrowings)
			if err != nil {
				return errors.New("gagal create pinjaman")
			}
		}
		return nil
	})
}

func (s *service_) GetBorrowingsByMembershipNumber(membershipNumber string) (ResponseGetMemberBorrowingBooks, error) {
	return s.repository.GetBorrowedBooksByMembershipNumber(membershipNumber)
}

// func (s *service_) GetAllBorrowing() []ResponseBorrowing {
// 	borrowings := s.repository.GetAllBorrowing()
// 	return borrowings
// }

func uniqueBookHubId(books []Book) ([]string, error) {
	if len(books) == 0 {
		return nil, errors.New("tidak ada request buku")
	}
	bookHubIds := make([]string, 0, len(books))
	seen := make(map[string]struct{}, len(books))
	for _, book := range books {
		if book.BookHubId == "" {
			return nil, errors.New("empty book hub id")
		}
		if _, duplicate := seen[book.BookHubId]; duplicate {
			return nil, errors.New("duplicate book")
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
