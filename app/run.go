package app

import (
	"be-library/app/db"
	"be-library/app/service/book"
	"be-library/app/service/category"
	"be-library/app/service/condition"
	"be-library/app/service/member"
	"be-library/app/service/shelf"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	db_ := db.Connection()
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(cors.Default())
	api := router.Group("/api")
	category.Router(api, db_)
	book.Router(api, db_)
	condition.Router(api, db_)
	member.Router(api, db_)
	shelf.Router(api,db_)
	//import your service in here...
port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	router.Run(fmt.Sprintf(":%v", port))
}
