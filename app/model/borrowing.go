package model

import "time"

type Borrowing struct {
	BaseModel
	MemberID          string    `gorm:"type:varchar(50);not null"`
	BookHubID         string    `gorm:"type:varchar(50);not null"`
	BorrowDate        time.Time `gorm:"type:datetime;not null"`
	ReturnDate        time.Time `gorm:"type:datetime;not null"`
	BorrowOfficerId   string    `gorm:"type:varchar(50);not null"`
	ReturnOfficerId   string    `gorm:"type:varchar(50);not null"`
	BorrowConditionId string    `gorm:"type:varchar(50);not null"`
	ReturnConditionId string    `gorm:"type:varchar(50);not null"`
	Status            string    `gorm:"type:varchar(50);not null"`
	Member            Member    `gorm:"foreignKey:MemberID;references:ID"`
	BookHub           BookHub   `gorm:"foreignKey:BookHubID;references:ID"`
	BorrowOfficer     Officer   `gorm:"foreignKey:BorrowOfficerId;references:ID"`
	ReturnOfficer     Officer   `gorm:"foreignKey:ReturnOfficerId;references:ID"`
	BorrowCondition   Condition `gorm:"foreignKey:BorrowConditionId;references:ID"`
	ReturnCondition   Condition `gorm:"foreignKey:ReturnConditionId;references:ID"`

	DateTime
}

func (Borrowing) TableName() string {
	return "borrowings"
}
