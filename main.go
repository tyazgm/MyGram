package main

import (
	"MyGram/database"
	"MyGram/route"

	"github.com/gin-gonic/gin"
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

	router.Run(PORT)
}
