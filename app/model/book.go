package model

type Book struct {
	BaseModel
	Title      string `gorm:"type:varchar(100);not null"`
	CategoryId string `gorm:"type:varchar(50);not null"`
	Stock    int      `gorm:"type:int;not null"`
	Author   string   `gorm:"type:varchar(100);not null"`
	Category Category `gorm:"foreignKey:CategoryId;references:ID"`
	DateTime
}

func (Book) TableName() string {
	return "books"
}
