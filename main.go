package main

import (
	"MyGram/database"
	"MyGram/route"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var PORT string = ":8080"

func main() {
	router := gin.Default()

	database.ConnectDB()
	db := database.GetDB()

	route.UserRoute(router, db)
	route.SocialMediaRoute(router, db)
	route.CommentRoute(router, db)
	route.PhotoRoute(router, db)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(PORT)
}
