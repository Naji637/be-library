package borrowing

type RequestBorrowing struct {
	MemberId string `json:"member_id" binding:"required"`
	Books    []Book `json:"books" validate:"dive"`
}

type Book struct {
	BookHubId string `json:"book_hub_id" binding:"required"`
}
