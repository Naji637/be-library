package model

type BookHub struct {
	BaseModel
	Barcode     string    `gorm:"type:varchar(100);not null"`
	BookID      string    `gorm:"type:varchar(50);not null"`
	ConditionID string    `gorm:"type:varchar(50);not null"`
	MemberID    *string    `gorm:"type:varchar(50)"`
	ShelfId     string    `gorm:"type:varchar(50);not null"`
	Price       float64   `gorm:"type:decimal(10,2);not null"`
	Book        Book      `gorm:"foreignKey:BookID;references:ID"`
	Condition   Condition `gorm:"foreignKey:ConditionID;references:ID"`
	Member      Member    `gorm:"foreignKey:MemberID;references:ID"`
	Shelf       Shelf     `gorm:"foreignKey:ShelfId;references:ID"`
	DateTime
}

func (BookHub) TableName() string {
	return "book_hubs"
}
