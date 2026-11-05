package fine

import "be-library/app/model"

type ResponseFine struct {
	model.BaseModel
	BorrowingId string  `json:"borrowing_id" binding:"required"`
	TotalDays   int     `json:"total_day" binding:"required"`
	TotalFine   float64 `json:"total_fine" binding:"required"`
	Status      string  `json:"statu" binding:"required"`
}

func (ResponseFine) TableName() string {
	return "fines"
}
