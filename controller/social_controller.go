package controller

import (
	"MyGram/helper"
	"MyGram/model"
	"MyGram/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SocialController struct {
	socialService service.SocialService
}

func NewSocialController(socialService service.SocialService) *SocialController {
	return &SocialController{
		socialService: socialService,
	}
}

func (sc *SocialController) CreateSocial(ctx *gin.Context) {
	var request model.SocialCreateRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	userID, IsExist := ctx.Get("userID")
	if !IsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "User does not exist",
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helper.SocialCreateValidator(request)
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

	response, err := sc.socialService.Create(request, userID.(string))
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
		Message: "Social media created successfully",
		Data:    response,
	})
}
