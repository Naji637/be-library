package borrowing

import (
	"time"
)

//	type ResponseBorrowing struct {
//		model.BaseModel
//		MemberId        string    `json:"member_id" binding:"required"`
//		BookHubId       string    `json:"book_hub_id" binding:"required"`
//		BorrowDate      time.Time `json:"borrow_date" binding:"required"`
//		ReturnDate      time.Time `json:"return_date" binding:"required"`
//		BorrowOfficerId string    `json:"borrow_officer_id" binding:"required"`
//		ReturnOfficerId string    `json:"return_officer_id" binding:"required"`
//		BorrowConditionId string `json:"borrow_condition_id" binding:"required"`
//		ReturnConditionId string `json:"return_condition_id" binding:"required"`
//		Status string `json:"status" binding:"required"`
//	}
// type responseBorrowingBook struct {
// 	model.BaseModel
// 	Name             string `json:"name"`
// 	MembershipNumber string `json:"membership_number"`
// }
//
// func (responseBorrowingBook) TableName() string {
// 	return "members"
// }

// type borrowings struct {
// }

// type ResponseJoinBorrowingToBooks struct {
// 	model.BaseModel
// 	Title      string    `json:"title" binding:"required"`
// 	BorrowDate time.Time `json:"borrow_date" binding:"required"`
// 	Condition  string    `json:"condition" binding:"required"`
// }
//
// type ResponseBorrowing struct {
// 	model.BaseModel
// 	MembershipNumber string                         `json:"membership_number" binding:"required"`
// 	Name             string                         `json:"name" binding:"required"`
// 	IsAllowed        *bool                          `json:"is_allowed"`
// 	BorrowedBooks    []ResponseJoinBorrowingToBooks `json:"borrowed_books"`
// }

type ResponseGetMemberBorrowingBooks struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// MembershipNumber string                   `json:"membership_number"`
	Borrowing []ResponseBorrowingBooks `json:"borrowing" gorm:"foreignKey:MemberId;references:ID"`
}

func (ResponseGetMemberBorrowingBooks) TableName() string {
	return "members"
}

type ResponseBorrowingBooks struct {
	ID                string                       `json:"id"`
	MemberId          string                       `json:"-"`
	BookHubId         string                       `json:"-"`
	BorrowConditionId string                       `json:"-"`
	BorrowDate        time.Time                    `json:"borrow_date"`
	BookHub           ResponseJoinBorrowingBookHub `json:"book_hub" gorm:"foreignKey:ID;references:BookHubId"`
	Condition         ResponseConditionBook        `json:"condition" gorm:"foreignKey:ID;references:BorrowConditionId"`
}

func (ResponseBorrowingBooks) TableName() string {
	return "borrowings"
}

type ResponseJoinBorrowingBookHub struct {
	ID     string       `json:"-"`
	BookId string       `json:"-"`
	Book   ResponseBook `json:"book" gorm:"foreignKey:ID;references:BookId"`
}

func (ResponseJoinBorrowingBookHub) TableName() string {
	return "book_hubs"
}

type ResponseBook struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (ResponseBook) TableName() string {
	return "books"
}

type ResponseConditionBook struct {
	ID        string `json:"-"`
	Condition string `json:"condition"`
}

func (ResponseConditionBook) TableName() string {
	return "conditions"
}
