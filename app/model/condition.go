package model

type Condition struct {
	BaseModel
	Condition string `gorm:"type:varchar(100);not null"`
	DateTime
}

func (Condition) TableName() string {
	return "conditions"
}