package model


type MemberClassification struct {
	BaseModel
	Classification string `gorm:"type:varchar(100);not null"`
	MaxBooks int `gorm:"type:int;not null"`
	MaxDays int `gorm:"type:int;not null"`
	DateTime
}

func (MemberClassification) TableName() string {
	return "member_classifications"
}	