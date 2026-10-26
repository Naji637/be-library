package condition

import "be-library/app/model"

type ResponseCondition struct {
	model.BaseModel
	Condition string `json:"condition" binding:"required"`
}

func (ResponseCondition) TableName() string {
	return "conditions"
}