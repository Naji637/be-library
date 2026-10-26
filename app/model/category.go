package model

type Category struct {
	BaseModel
	Category string `gorm:"type:varchar(100);not null"`
	DateTime
}

func (Category) TableName() string {
	return "categories"
}	