package model

type Fine struct {
	BaseModel
	BorrowingID string `gorm:"type:varchar(50);not null"`

	TotalDays int       `gorm:"type:int;not null"`
	TotalFine float64   `gorm:"type:decimal(10,2);not null"`
	Status    string    `gorm:"type:varchar(50);not null"`
	Borrowing Borrowing `gorm:"foreignKey:BorrowingID;references:ID"`
	DateTime
}

func (Fine) TableName() string {
	return "fines"
}