package controller

import (
	"MyGram/helper"
	"MyGram/model"
	"MyGram/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct {
	commentService service.CommentService
}

func NewCommentController(commentService service.CommentService) *CommentController {
	return &CommentController{
		commentService: commentService,
	}
}

// CreateComment godoc
// @summary add comment
// @description add comment data to the database
// @tags Comment
// @accept json
// @produce json
// @Param photoID path string true "photoID you want to retrieve"
// @param request body model.CommentCreateRequest true "request is mandatory"
// @succes 200 {object} model.SuccessResponse{data=*model.CommentResponse}
// @failure	400	{object} model.ErrorResponse{errors=interface{}}
// @failure	500	{object} model.ErrorResponse{errors=interface{}}
// @security BearerAuth
// @router /comment/{photoID} [POST]
func (cc *CommentController) CreateComment(ctx *gin.Context) {
	var request model.CommentCreateRequest
	photoID := ctx.Param("photoID")

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	userID, isExist := ctx.Get("userID")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "User does not exist",
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helper.CommentCreateValidator(request)
	if validateErrs != nil {
		errResponseStr := make([]string, len(validateErrs))
		for i, err := range validateErrs {
			errResponseStr[i] = err.Error()
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errResponseStr,
		})
		return
	}

	response, err := cc.commentService.Create(request, userID.(string), photoID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: "Bad Request",
				Errors: errors.New("Photo not found").Error(),
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Comment has been created successfully",
		Data:    response,
	})
}

// GetAll godoc
// @summary get all comment
// @description get all comment data to the database
// @tags Comment
// @accept json
// @produce json
// @succes 200 {object} model.SuccessResponse{data=[]model.CommentResponse}
// @failure	400	{object} model.ErrorResponse{errors=interface{}}
// @failure	500	{object} model.ErrorResponse{errors=interface{}}
// @security BearerAuth
// @router /comment [GET]
func (cc *CommentController) GetAll(ctx *gin.Context) {
	response, err := cc.commentService.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Get all comment successfully",
		Data:    response,
	})
}

// GetOne godoc
// @summary get comment
// @description get on comment data from comment ID
// @tags Comment
// @param commentID path string true "commentID you want to retrieve"
// @accept json
// @produce json
// @Param commentID path string true "commentID you want to retrieve"
// @succes 200 {object} model.SuccessResponse{data=model.CommentResponse}
// @failure	400	{object} model.ErrorResponse{errors=interface{}}
// @failure	500	{object} model.ErrorResponse{errors=interface{}}
// @security BearerAuth
// @router /comment/{commentID} [GET]
func (cc *CommentController) GetOne(ctx *gin.Context) {
	commentID := ctx.Param("commentID")

	response, err := cc.commentService.GetOne(commentID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Get comment successfully",
		Data:    response,
	})
}

// UpdateComment godoc
// @summary update comment
// @description update comment by commentID
// @tags Comment
// @param commentID path string true "commentID you want to retrieve"
// @param request body model.CommentUpdateRequest true "request is mandatory"
// @accept json
// @produce json
// @succes 200 {object} model.SuccessResponse{data=model.CommentUpdateResponse}
// @failure	400	{object} model.ErrorResponse{errors=interface{}}
// @failure	500	{object} model.ErrorResponse{errors=interface{}}
// @security BearerAuth
// @router /comment/{commentID} [PUT]
func (cc *CommentController) UpdateComment(ctx *gin.Context) {
	var request model.CommentUpdateRequest
	commentID := ctx.Param("commentID")

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	userID, isExist := ctx.Get("userID")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helper.CommentUpdateValidator(request)
	if validateErrs != nil {
		errResponseStr := make([]string, len(validateErrs))
		for i, err := range validateErrs {
			errResponseStr[i] = err.Error()
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errResponseStr,
		})
		return
	}

	response, err := cc.commentService.UpdateComment(request, userID.(string), commentID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Comment updated successfully",
		Data: model.CommentUpdateResponse{
			ID: response.ID,
		},
	})
}

// DeleteComment godoc
// @summary delete a comment
// @description delete a comment by commentID
// @tags Comment
// @param commentID path string true "commentID you want to retrieve"
// @accept json
// @produce json
// @succes 200 {object} model.SuccessResponse{data=model.CommentDeleteResponse}
// @failure	400	{object} model.ErrorResponse{errors=interface{}}
// @failure	500	{object} model.ErrorResponse{errors=interface{}}
// @security BearerAuth
// @router /comment/{commentID} [DELETE]
func (cc *CommentController) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("commentID")

	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	response, err := cc.commentService.Delete(commentID, userID.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Comment deleted successfully",
		Data: model.CommentDeleteResponse{
			ID: response.ID,
		},
	})
}
