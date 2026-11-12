package borrowing

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateBorrowing(borrowing *model.Borrowing) error
	GetBorrowedBooksByMembershipNumber(membershipNumber string) (ResponseGetMemberBorrowingBooks, error)
	ReturnBorrowing(id string, borrowingReturn *model.Borrowing) error
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
	err := r.db.Debug().Preload("Borrowing", func(db *gorm.DB) *gorm.DB {
		return db.Joins("BookHub.Book").Joins("Condition")
	}).Where("membership_number=?", membershipNumber).Find(&borrowedBooks).Error
	return borrowedBooks, err
}

func (r *repository) ReturnBorrowing(id string, borrowingReturn *model.Borrowing) error {
	err := r.db.Model(&model.Borrowing{}).Where("id=?", id).Updates(&borrowingReturn).Error
	return err
}

// func (r *repository) BorrowedBooks(bookId string) {
//
// }

// func (r *repository) GetAllBorrowing() []ResponseBorrowing {
// 	var borrowings []ResponseBorrowing
// 	r.db.Find(&borrowings)
// 	return borrowings
// }
