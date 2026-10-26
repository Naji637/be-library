package category

type RequestCategory struct {
	Category string `json:"category" binding:"required"`
}