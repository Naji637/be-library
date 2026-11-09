package borrowing

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateBorrowing(borrowing *model.Borrowing) error
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

// func (r *repository) GetBorrowedBooksByMemberId(memberId string) (ResponseJoinBorrowingToBooks, error) {
// 	var borrowedBooks ResponseJoinBorrowingToBooks
//
// 	query :=
// 		`select br.id,b.title,br.borrow_date,bh.condition as `
// }

// func (r *repository) GetAllBorrowing() []ResponseBorrowing {
// 	var borrowings []ResponseBorrowing
// 	r.db.Find(&borrowings)
// 	return borrowings
// }
