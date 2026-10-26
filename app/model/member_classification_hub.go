package model

type MemberClassificationHub struct {
	BaseModel
	MemberID               string `gorm:"type:varchar(50);not null"`
	MemberClassificationID string `gorm:"type:varchar(50);not null"`
	Member                 Member `gorm:"foreignKey:MemberID;references:ID"`
	MemberClassification MemberClassification `gorm:"foreignKey:MemberClassificationID;references:ID"`
	DateTime
}
