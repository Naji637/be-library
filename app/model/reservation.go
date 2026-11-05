package model

import "time"

type Reservation struct {	
BaseModel
	MemberID string `gorm:"type:varchar(50);not null"`
	BookHubID string `gorm:"type:varchar(50);not null"`
	PickupDeadline time.Time `gorm:"type:datetime;not null"`
	Member Member `gorm:"foreignKey:MemberID;references:ID"`
	BookHub BookHub `gorm:"foreignKey:BookHubID;references:ID"`
	DateTime
}