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

// CreateSocial godoc
// @summary add social media
// @description add social media data to the database
// @tags Social Media
// @accept json
// @produce json
// @param request body model.SocialCreateRequest true "request is mandatory"
// @succes 200 {object} model.SuccessResponse{data=*model.SocialResponse}
// @failure	400	{object} model.ErrorResponse{errors=interface{}}
// @failure	500	{object} model.ErrorResponse{errors=interface{}}
// @security BearerAuth
// @router /socialmedia [POST]
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

// GetAll godoc
// @summary get all social media
// @description get all social media data to the database
// @tags Social Media
// @accept json
// @produce json
// @succes 200 {object} model.SuccessResponse{data=model.[]SocialResponse}
// @failure	400	{object} model.ErrorResponse{errors=interface{}}
// @failure	500	{object} model.ErrorResponse{errors=interface{}}
// @security BearerAuth
// @router /socialmedia [GET]
func (sc *SocialController) GetAll(ctx *gin.Context) {
	response, err := sc.socialService.GetAll()
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
		Message: "All social media info has been successfully retrieved",
		Data:    response,
	})
}

// GetOne godoc
// @summary get social media
// @description get on social media data from socialmedia ID
// @tags Social Media
// @param socialmediaID path string true "socialmediaID you want to retrieve"
// @accept json
// @produce json
// @succes 200 {object} model.SuccessResponse{data=model.SocialResponse}
// @failure	400	{object} model.ErrorResponse{errors=interface{}}
// @failure	500	{object} model.ErrorResponse{errors=interface{}}
// @security BearerAuth
// @router /socialmedia/{socialmediaID} [GET]
func (sc *SocialController) GetOne(ctx *gin.Context) {
	socialMediaID := ctx.Param("socialmediaID")

	response, err := sc.socialService.GetOne(socialMediaID)
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
		Message: "Social media data has been successfully retrieved",
		Data:    response,
	})
}

// UpdateSocialMedia godoc
// @summary update social media
// @description update social media by socialmediaID
// @tags Social Media
// @param socialmediaID path string true "socialmediaID you want to retrieve"
// @param request body model.SocialUpdateRequest true "request is mandatory"
// @accept json
// @produce json
// @succes 200 {object} model.SuccessResponse{data=model.SocialUpdateResponse}
// @failure	400	{object} model.ErrorResponse{errors=interface{}}
// @failure	500	{object} model.ErrorResponse{errors=interface{}}
// @security BearerAuth
// @router /socialmedia/{socialmediaID} [PUT]
func (sc *SocialController) UpdateSocialMedia(ctx *gin.Context) {
	var request model.SocialUpdateRequest
	socialID := ctx.Param("socialmediaID")

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
	validateErrs = helper.SocialUpdateValidator(request)
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

	response, err := sc.socialService.UpdateSocialMedia(request, userID.(string), socialID)
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
		Message: "Social media has been successfully updated",
		Data: model.SocialUpdateResponse{
			ID: response.ID,
		},
	})
}

// DeleteSocialMedia godoc
// @summary delete a social media
// @description delete a social media by socialmediaID
// @tags Social Media
// @param socialmediaID path string true "socialmediaID you want to retrieve"
// @accept json
// @produce json
// @succes 200 {object} model.SuccessResponse{data=model.SocialDeleteResponse}
// @failure	400	{object} model.ErrorResponse{errors=interface{}}
// @failure	500	{object} model.ErrorResponse{errors=interface{}}
// @security BearerAuth
// @router /socialmedia/{socialmediaID} [DELETE]
func (sc *SocialController) DeleteSocialMedia(ctx *gin.Context) {
	socialID := ctx.Param("socialmediaID")

	userID, isExist := ctx.Get("userID")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	response, err := sc.socialService.Delete(socialID, userID.(string))
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
		Message: "Social media has been successfully deleted",
		Data: model.SocialDeleteResponse{
			ID: response.ID,
		},
	})
}
