package model

type Shelf struct {
	BaseModel
	ShelfNumber string `gorm:"type:varchar(100);not null"`
	DateTime
}

func (Shelf) TableName() string {
	return "shelves"
}
