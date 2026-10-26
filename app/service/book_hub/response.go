package book_hub

import "be-library/app/model"

type ResponseBookHub struct {
	model.BaseModel
	Barcode     string  `json:"barcode" binding:"required"`
	BookId      string  `json:"book_id" binding:"required"`
	ConditionId string  `json:"condition_id" binding:"required"`
	MemberId    string  `json:"member_id" binding:"required"`
	ShelfId     string  `json:"shelf_id" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

func (ResponseBookHub) TableName()string{
	return "book_hubs"
}