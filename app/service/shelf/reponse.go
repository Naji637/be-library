package shelf

import "be-library/app/model"

type ResponseShelf struct{
 model.BaseModel
 ShelfNumber string `json:"shelf_number" binding:"required"`
}

func (ResponseShelf) TableName() string {
	return "shelves"
}