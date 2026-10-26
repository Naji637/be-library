package book

import "be-library/app/model"

type ResponseBook struct {
	model.BaseModel
	Title      string `json:"title" binding:"required"`
	CategoryId string `json:"category_id" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	Author     string `json:"author" binding:"required"`
}

func (ResponseBook) TableName() string {
	return "books"
}
