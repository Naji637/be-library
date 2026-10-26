package book_hub

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(g *gin.RouterGroup, db *gorm.DB) {
	repository := NewRepository(db)
	service_ := NewService(repository)
	route := g.Group("/book_hubs")

	route.GET("/", func(c *gin.Context) {
		bookHubs := service_.GetAllBookHubs()
		c.JSON(200, bookHubs)
	})
	route.POST("/", func(c *gin.Context) {
		var request RequestBookHub
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err := service_.CreateBookHub(request)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "Created"})
	})
}
