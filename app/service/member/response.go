package member

import "be-library/app/model"

type ResponseMember struct {
	model.BaseModel
	MembershipNumber string `json:"membership_number" binding:"required"`
	Name             string `json:"name" binding:"required"`
	Alumni           bool   `json:"alumni" binding:"required"`
	model.DateTime
}

func (ResponseMember) TableName() string {
	return "members"
}

type ResponseJoinMembersToMemberClassifications struct {
	model.BaseModel
	MembershipNumber string `json:"membership_number" binding:"required"`
	Name             string `json:"name" binding:"required"`
	MaxBooks         int    `json:"max_books" binding:"required"`
	MaxDays          int    `json:"max_days" binding:"required"`
}
