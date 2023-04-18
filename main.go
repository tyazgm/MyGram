package main

import (
	"MyGram/database"

	"github.com/gin-gonic/gin"
)

var PORT string = ":8080"

func main() {
	router := gin.Default()

	database.ConnectDB()
	// db := database.GetDB()

	router.Run(PORT)
}
