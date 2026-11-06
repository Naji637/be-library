package model

type Barcode struct {
	BaseModel
	Year   uint8
	Month  uint8
	Number uint8
	DateTime
}

func (Barcode) TableName() string {
	return "barcode"
}
