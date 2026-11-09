package borrowing

import (
	"be-library/app/service/member"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(g *gin.RouterGroup, db *gorm.DB) {
	repository := NewRepository(db)
	memberRepo := member.NewRepository(db)
	service_ := NewService(repository, memberRepo)
	route := g.Group("/borrowings")

	// route.GET("", func(c *gin.Context) {
	// 	borrowings := service_.GetAllBorrowing()
	// 	c.JSON(200, borrowings)
	// })

	// 	route.POST("", func(c *gin.Context) {
	// 		var request RequestBorrowing
	// 		if err := c.ShouldBindJSON(&request); err != nil {
	// 			c.JSON(400, gin.H{"error": err.Error()})
	// 			return
	// 		}
	//
	// 		err := service_.CreateBorrowing(request)
	// 		if err != nil {
	// 			c.JSON(500, gin.H{"error": err.Error()})
	// 			return
	// 		}
	//
	// 		c.JSON(201, gin.H{"message": "Created"})
	// 	})

	route.GET("/:membersip-number", func(c *gin.Context) {
		membershipNumber := c.Param("membership-number")

		borrowing := service_.GetBorrowingsByMembershipNumber(membershipNumber)
		c.JSON(201, borrowing)

	})
}
