package shelf

type RequestShelf struct{
	  ShelfNumber string `json:"shelf_number" binding:"required"`

}