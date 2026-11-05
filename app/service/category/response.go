package category
import "be-library/app/model"

type ResponseCategory struct {
	model.BaseModel
	Category string `json:"category"`
}

func (ResponseCategory) TableName() string {
	return "categories"
}