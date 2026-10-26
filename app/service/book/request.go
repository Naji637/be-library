package book

type RequestBook struct {
	Title      string `json:"title" binding:"required"`
	CategoryId string `json:"category_id" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	Author     string `json:"author" binding:"required"`
}
