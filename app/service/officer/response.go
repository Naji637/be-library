package officer

import "be-library/app/model"

type ResponseOfficer struct {
	model.BaseModel
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ResponseOfficer) TableName() string{
	return "officers"
}