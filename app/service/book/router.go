package book

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(g *gin.RouterGroup, db *gorm.DB) {
	repository := NewRepository(db)
	service_ := NewService(repository)
	route := g.Group("/books")

	route.GET("", func(c *gin.Context) {
		books := service_.GetAllBooks()
		c.JSON(200, books)
	})
	route.POST("", func(c *gin.Context) {
		var request RequestBook
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err := service_.CreateBook(request)
		if err != nil {
			c.JSON(500, gin.H{"error": "ups"})
			return
		}
		c.JSON(201, gin.H{"message": "Created"})
	})
	route.PUT("/:id", func(c *gin.Context) {
		var request RequestBook
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		id := c.Param("id")
		fmt.Println(id)
		err := service_.UpdateBook(request, id)
		if err != nil {
			c.JSON(500, gin.H{"error": "ups"})
			return
		}
		c.JSON(201, gin.H{"message": "Updated"})
	})
	route.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := service_.DeleteBook(id)
		if err != nil {
			c.JSON(500, gin.H{"error": "ups"})
			return
		}
		c.JSON(201, gin.H{"message": "Deleted"})
	})
}
