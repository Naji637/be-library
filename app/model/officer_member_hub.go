package model

type OfficerMemberHub struct {	
	MemberID  string `gorm:"type:varchar(50);not null"`
	OfficerID string `gorm:"type:varchar(50);not null"`
	Member    Member `gorm:"foreignKey:MemberID;references:ID"`
	Officer   Officer `gorm:"foreignKey:OfficerID;references:ID"`
	DateTime
}

func (OfficerMemberHub) TableName() string {
	return "officer_member_hubs"
}