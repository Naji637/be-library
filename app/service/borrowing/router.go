package borrowing

import (
	"be-library/app/service/book_hub"
	"be-library/app/service/member"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(g *gin.RouterGroup, db *gorm.DB) {
	repository := NewRepository(db)
	memberRepo := member.NewRepository(db)
	bookHubRepo := book_hub.NewRepository(db)
	service_ := NewService(repository, memberRepo, bookHubRepo)
	route := g.Group("/borrowings")

	// route.GET("", func(c *gin.Context) {
	// 	borrowings := service_.GetAllBorrowing()
	// 	c.JSON(200, borrowings)
	// })

	route.POST("", func(c *gin.Context) {
		var request RequestBorrowing
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err := service_.CreateBorrowing(request)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"message": "Created"})
	})

	route.GET("/:membership-number", func(c *gin.Context) {
		membershipNumber := c.Param("membership-number")
		borrowing, err := service_.GetBorrowingsByMembershipNumber(membershipNumber)
		if err != nil {
			c.JSON(500, gin.H{"message": "ups"})
			return
		}
		c.JSON(200, borrowing)

	})
	route.PUT("", func(c *gin.Context) {
		var request []RequestReturnBorrowing
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"message": "ups"})
			return
		}

		err := service_.ReturnBorrowing(request)
		if err != nil {
			c.JSON(500, gin.H{"message": "ups"})
			return
		}
		c.JSON(201, gin.H{"message": "updated"})

	})
}
