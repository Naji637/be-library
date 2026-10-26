package condition

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(g *gin.RouterGroup, db *gorm.DB) {
	repository := NewRepository(db)
	service_ := NewService(repository)
	route := g.Group("/conditions")

	route.GET("/",func(c *gin.Context) {
		conditions:=service_.GetAllConditions()
		c.JSON(200,conditions)
	})
route.POST("/", func(c *gin.Context) {
		var request RequestCondition
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err := service_.CreateCondition(request)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"message": "Created"})
	})
}