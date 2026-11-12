package borrowing

import (
	"be-library/app/model"
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	CreateBorrowing(borrowing *model.Borrowing) error
	GetBorrowedBooksByMembershipNumber(membershipNumber string) (ResponseGetMemberBorrowingBooks, error)
	ReturnBorrowing(id string, borrowingReturn *model.Borrowing) error
	BorrowedBooks(bookId string) (int, error)
	// GetAllBorrowing() []ResponseBorrowing
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateBorrowing(borrowing *model.Borrowing) error {
	err := r.db.Create(borrowing).Error
	return err
}

func (r *repository) GetBorrowedBooksByMembershipNumber(membershipNumber string) (ResponseGetMemberBorrowingBooks, error) {
	var borrowedBooks ResponseGetMemberBorrowingBooks
	err := r.db.Preload("Borrowing", func(db *gorm.DB) *gorm.DB {
		return db.Joins("BookHub.Book").Joins("Condition")
	}).Where("membership_number=?", membershipNumber).Find(&borrowedBooks).Error
	return borrowedBooks, err
}

func (r *repository) ReturnBorrowing(id string, borrowingReturn *model.Borrowing) error {
	err := r.db.Model(&model.Borrowing{}).Where("id=?", id).Updates(&borrowingReturn).Error
	return err
}

func (r *repository) BorrowedBooks(bookId string) (int, error) {
	var borrowings []model.Borrowing
	var count int
	err := r.db.Debug().Model(&model.Borrowing{}).
		// Select("book_hubs.*, borrowings.*").
		Joins("join book_hubs on book_hubs.id = borrowings.book_hub_id").
		Where("book_hubs.book_id = ? and borrowings.return_date is null", bookId).
		Find(&borrowings).Error
	log.Println("%%%%%", borrowings)
	if err != nil {
		return count, err
	}
	count = len(borrowings)
	log.Println("count: ", count)
	return count, err
}

// func (r *repository) GetAllBorrowing() []ResponseBorrowing {
// 	var borrowings []ResponseBorrowing
// 	r.db.Find(&borrowings)
// 	return borrowings
// }
