package model

type Member struct {
	BaseModel
	MembershipNumber string `gorm:"type:varchar(50);not null"`
	Name string `gorm:"type:varchar(100);not null"`
	Alumni bool `gorm:"type:boolean;not null"`
	DateTime
}

func (Member) TableName() string {
	return "members"
}
