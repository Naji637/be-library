package model

type Membership struct {
	BaseModel
	Year   uint8
	Month  uint8
	Number uint8
	DateTime
}

func (Membership) TableName() string {
	return "memberships"
}
