package member

import "be-library/app/model"

type ResponseMember struct {
	model.BaseModel
	MembershipNumber string `json:"membership_number"`
	Name             string `json:"name"`
	Alumni           bool   `json:"alumni"`
	model.DateTime
}

func (ResponseMember) TableName() string {
	return "members"
}

type ResponseJoinMembersToMemberClassifications struct {
	model.BaseModel
	MembershipNumber string `json:"membership_number"`
	Name             string `json:"name"`
	MaxBooks         int    `json:"max_books"`
	MaxDays          int    `json:"max_days"`
}

func (ResponseJoinMembersToMemberClassifications) TableName() string {
	return "member_classification_hubs"
}
