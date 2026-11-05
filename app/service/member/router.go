package member

import (
	"be-library/app/service/member_classification_hub"
	util "be-library/app/util/member"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(g *gin.RouterGroup, db *gorm.DB) {
	repository := NewRepository(db)
	memberRepo := util.NewRepository(db)
	member := util.NewMember(memberRepo)
	mchRepo := member_classification_hub.NewRepository(db)
	service_ := NewService(repository, member, mchRepo)
	route := g.Group("/members")

	route.GET("", func(c *gin.Context) {
		members := service_.GetAllMembers()
		c.JSON(200, members)
	})
	route.POST("", func(c *gin.Context) {
		var request RequestMember
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err := service_.CreateMember(request)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"message": "Created"})
	})
	route.PUT("/:id", func(c *gin.Context) {
		var request RequestMember
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		id := c.Param("id")
		err := service_.UpdateMember(request, id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "Updated"})
	})
	route.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := service_.DeleteMember(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "Deleted"})
	})
}
