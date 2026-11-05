package member_classification

type RequestMemberClassification struct{
	Classification string `json:"classification" binding:"required"`
	MaxBooks int `json:"max_books" binding:"required"`
	MaxDays int `json:"max_days" binding:"required"`
}