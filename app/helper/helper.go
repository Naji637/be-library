package helper

import "github.com/gin-gonic/gin"
type TypeOfMessage struct {
	Create gin.H
	Update gin.H
	Delete gin.H
}

var GetMessage = TypeOfMessage{
	Create: gin.H{"message": "saved"},
	Update: gin.H{"message": "updated"},
	Delete: gin.H{"message": "deleted"},
}