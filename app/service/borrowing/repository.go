package borrowing

import (
	"be-library/app/model"

	"gorm.io/gorm"
)

type Repository interface {
	CreateBorrowing(borrowing *model.Borrowing) error
	GetBorrowedBooksByMembershipNumber(membershipNumber string) (ResponseGetMemberBorrowingBooks, error)
	ReturnBorrowing(id string, borrowingReturn *model.Borrowing) error
	JoinBookToBookHub(barcode string) (ResponseJoinBookHubToBook, error)
	GetTotalBorrowingByMemberId(memberId string) (int64, error)
	GetBorrowingsByBookHubIds(bookHubIds []string) []model.Borrowing
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
func (r *repository) JoinBookToBookHub(barcode string) (ResponseJoinBookHubToBook, error) {
	var book ResponseJoinBookHubToBook
	err := r.db.Table("book_hubs").
		Select("book_hubs.*,books.*,conditions.*").
		Joins("join books on books.id = book_hubs.book_id").
		Joins("join conditions on conditions.id = book_hubs.condition_id").
		Where("book_hubs.barcode =?", barcode).Scan(&book).Error
	return book, err
}
func (r *repository) GetBorrowingsByBookHubIds(bookHubIds []string) []model.Borrowing {
	var borrowings []model.Borrowing

	r.db.Where("book_hub_id in ? and return_date is null", bookHubIds).Find(&borrowings)
	return borrowings
}
func (r *repository) GetTotalBorrowingByMemberId(memberId string) (int64, error) {
	var borrowedBooks int64
	err := r.db.Model(&model.Borrowing{}).Where("member_id=?", memberId).Count(&borrowedBooks).Error
	return borrowedBooks, err

}

// func (r *repository) BorrowedBooks(bookId string) {
//
// }

// func (r *repository) GetAllBorrowing() []ResponseBorrowing {
// 	var borrowings []ResponseBorrowing
// 	r.db.Find(&borrowings)
// 	return borrowings
// }
