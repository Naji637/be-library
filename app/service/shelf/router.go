package shelf

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(g *gin.RouterGroup, db *gorm.DB) {
	repository := NewRepository(db)
	service_ := NewService(repository)
	route := g.Group("/shelves")

	route.GET("",func(c *gin.Context) {
		shelves:=service_.GetAllShelves()
		c.JSON(200,shelves)
	})
	route.POST("",func(c *gin.Context) {
		var request RequestShelf
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err := service_.CreateShelf(request)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"message": "Created"})
	})

}