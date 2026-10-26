package borrowing

import (
	"be-library/app/model"
	"time"
)

type ResponseBorrowing struct {
	model.BaseModel
	MemberId        string    `json:"member_id" binding:"required"`
	BookHubId       string    `json:"book_hub_id" binding:"required"`
	BorrowDate      time.Time `json:"borrow_date" binding:"required"`
	ReturnDate      time.Time `json:"return_date" binding:"required"`
	BorrowOfficerId string    `json:"borrow_officer_id" binding:"required"`
	ReturnOfficerId string    `json:"return_officer_id" binding:"required"`
	BorrowConditionId string `json:"borrow_condition_id" binding:"required"`
	ReturnConditionId string `json:"return_condition_id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

func (ResponseBorrowing) TableName() string {
	return "borrowing"
}
