package route

import (
	"MyGram/controller"
	"MyGram/repository"
	"MyGram/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoute(router *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	photoRepository := repository.NewPhotoRepository(db)
	socialRepository := repository.NewSocialRepository(db)
	userService := service.NewUserService(*userRepository, *photoRepository, *socialRepository)
	userController := controller.NewUserController(*userService)

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

}
