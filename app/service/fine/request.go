package fine

type RequestFine struct {
	BorrowingId string  `json:"borrowing_id" binding:"required"`
	TotalDays   int     `json:"total_day" binding:"required"`
	TotalFine   float64 `json:"total_fine" binding:"required"`
	Status      string  `json:"statu" binding:"required"`
}
