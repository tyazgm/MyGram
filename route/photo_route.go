package route

import (
	"MyGram/controller"
	"MyGram/middleware"
	"MyGram/repository"
	"MyGram/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PhotoRoute(router *gin.Engine, db *gorm.DB) {
	photoRepository := repository.NewPhotoRepository(db)
	commentRepository := repository.NewCommentRepository(db)
	photoService := service.NewPhotoService(*photoRepository, *commentRepository)
	photoController := controller.NewPhotoController(*photoService)

	authUser := router.Group("/photos", middleware.AuthMiddleware)
	{
		authUser.POST("", photoController.CreatePhoto)
	}
}
