package route

import (
	"MyGram/controller"
	"MyGram/middleware"
	"MyGram/repository"
	"MyGram/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommentRoute(router *gin.Engine, db *gorm.DB) {
	commentRepository := repository.NewCommentRepository(db)
	photoRepository := repository.NewPhotoRepository(db)
	commentService := service.NewCommentService(*commentRepository, *photoRepository)
	commentController := controller.NewCommentController(*commentService)

	authUser := router.Group("/comment", middleware.AuthMiddleware)
	{
		authUser.POST("/:photoID", commentController.CreateComment)
		authUser.GET("", commentController.GetAll)
		authUser.GET("/:commentID", commentController.GetOne)
		authUser.PUT("/:commentID", commentController.UpdateComment)
		authUser.DELETE("/:commentID", commentController.DeleteComment)
	}
}
