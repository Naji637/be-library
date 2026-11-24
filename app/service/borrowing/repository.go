package borrowing

import (
	"be-library/app/model"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	CreateBorrowing(borrowings []*model.Borrowing) error
	GetBorrowedBooksByMembershipNumber(membershipNumber string) (ResponseGetMemberBorrowingBooks, error)
	GetActiveBorrowedBookIdsByMemberId(memberId string) ([]string, error)
	Transaction(fn func(repo Repository) error) error
	GetBookHubForBorrowing(bookHubIds []string) ([]BookHubForBorrowing, error)
	// GetAllBorrowing() []ResponseBorrowing
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Transaction(fn func(repo Repository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return fn(&repository{db: tx})
	})
}

func (r *repository) CreateBorrowing(borrowings []*model.Borrowing) error {
	err := r.db.Create(&borrowings).Error
	return err
}

func (r *repository) GetBorrowedBooksByMembershipNumber(membershipNumber string) (ResponseGetMemberBorrowingBooks, error) {
	var borrowedBooks ResponseGetMemberBorrowingBooks
	err := r.db.Preload("Borrowing", func(db *gorm.DB) *gorm.DB {
		return db.Joins("BookHub.Book").Joins("Condition")
	}).Where("membership_number=?", membershipNumber).Find(&borrowedBooks).Error
	return borrowedBooks, err
}

// "select bh.book_id from book_hubs as bh inner join borrowings as b on b.book_hub_id = bh.id where b.member_id=?"
func (r *repository) GetActiveBorrowedBookIdsByMemberId(memberId string) ([]string, error) {
	var bookIds []string
	err := r.db.Table("borrowings br").
		Joins("join book_hubs as bh on bh.id=br.book_hub_id and bh.deleted_at is null").
		Where("br.member_id=? and br.return_date is null and br.deleted_at is null", memberId).
		Pluck("bh.book_id", &bookIds).Error
	return bookIds, err

}

func (r *repository) GetBookHubForBorrowing(bookHubIds []string) ([]BookHubForBorrowing, error) {
	var bookHubs []BookHubForBorrowing
	err := r.db.Debug().Table("book_hubs bh").Select("bh.id, bh.book_id, bh.condition_id,(br.id is not null) as is_borrowed").
		Joins("left join borrowings as br on br.book_hub_id=bh.id and br.return_date is null and br.deleted_at is null").
		Where("bh.id in ? and bh.deleted_at is null", bookHubIds).
		Clauses(clause.Locking{Strength: "UPDATE"}).Scan(&bookHubs).Error
	return bookHubs, err
}

// func (r *repository) GetAllBorrowing() []ResponseBorrowing {
// 	var borrowings []ResponseBorrowing
// 	r.db.Find(&borrowings)
// 	return borrowings
// }
