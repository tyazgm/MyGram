package route

import (
	"MyGram/controller"
	"MyGram/middleware"
	"MyGram/repository"
	"MyGram/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SocialMediaRoute(router *gin.Engine, db *gorm.DB) {
	socialRepository := repository.NewSocialRepository(db)
	socialService := service.NewSocialService(*socialRepository)
	socialController := controller.NewSocialController(*socialService)

	authUser := router.Group("/social-media", middleware.AuthMiddleware)
	{
		authUser.POST("", socialController.CreateSocial)
	}
}
