package book

type RequestBook struct {
	Title      string `json:"title" binding:"required"`
	CategoryId string `json:"category_id" binding:"required"`
	// Stock      int    `json:"stock" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type RequestBarcode struct {
	BookId      string  `json:"book_id" binding:"required"`
	ConditionId string  `json:"condition_id" binding:"required"`
	MemberId    *string `json:"member_id"`
	ShelfId     string  `json:"shelf_id" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}
