package model

type Officer struct {
	BaseModel
	Name string `gorm:"type:varchar(100);not null"`
	Username string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	DateTime
}

func (Officer) TableName() string	 {
	return "officers"
}