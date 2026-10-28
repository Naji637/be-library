package member_classification

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(g *gin.RouterGroup, db *gorm.DB) {
	repository := NewRepository(db)
	service_ := NewService(repository)
	route := g.Group("/member_classifications")

	route.GET("",func(c *gin.Context) {
		memberClassification:=service_.GetAllMemberClassifications()
		c.JSON(200,memberClassification)
	})
	route.POST("",func(c *gin.Context) {
		var request RequestMemberClassification
		if err:=c.ShouldBindJSON(&request);err!=nil{
			c.JSON(400,gin.H{"error":err.Error()})
			return
		}
		err:=service_.CreateMemberClassification(request)
		if err!=nil{
			c.JSON(500,gin.H{"error":err.Error()})
			return
		}
		c.JSON(201,gin.H{"message":"Created"})
	})
}