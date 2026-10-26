package condition

type RequestCondition struct {
	Condition string `json:"condition" binding:"required"`
}