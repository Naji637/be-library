package book

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(g *gin.RouterGroup, db *gorm.DB) {
	repository := NewRepository(db)
	service_ := NewService(repository)
	route := g.Group("/books")

	route.GET("/",func(c *gin.Context) {
		books:=service_.GetAllBooks()
		c.JSON(200,books)
	})
	route.POST("/", func(c *gin.Context) {
		var request RequestBook
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err := service_.CreateBook(request)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"message": "Created"})
	})

}