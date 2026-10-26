package member_classification

import "be-library/app/model"

type ResponseMemberClassification struct{
	model.BaseModel
	Classification string `json:"classification" binding:"required"`
	MaxBooks int `json:"max_books" binding:"required"`
	MaxDays int `json:"max_days" binding:"required"`
}

func (ResponseMemberClassification) TableName()string{
	return "member_classifications"
}