package model

type OfficerHubMember struct {	
	MemberID  string `gorm:"type:varchar(50);not null"`
	OfficerID string `gorm:"type:varchar(50);not null"`
	Member    Member `gorm:"foreignKey:MemberID;references:ID"`
	Officer   Officer `gorm:"foreignKey:OfficerID;references:ID"`
	DateTime
}

func (OfficerHubMember) TableName() string {
	return "officer_hub_members"
}