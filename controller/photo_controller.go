package controller

import (
	"MyGram/helper"
	"MyGram/model"
	"MyGram/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	photoService service.PhotoService
}

func NewPhotoController(photoService service.PhotoService) *PhotoController {
	return &PhotoController{
		photoService: photoService,
	}
}

// CreatePhoto godoc
// @summary add photo
// @description add photo data to the database
// @tags Photo
// @accept json
// @produce json
// @param request body model.PhotoCreateRequest true "request is mandatory"
// @succes 200 {object} model.SuccessResponse{data=*model.PhotoCreateResponse}
// @failure	400	{object} model.ErrorResponse{errors=interface{}}
// @failure	500	{object} model.ErrorResponse{errors=interface{}}
// @security BearerAuth
// @router /photo [POST]
func (pc *PhotoController) CreatePhoto(ctx *gin.Context) {
	var request model.PhotoCreateRequest

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
	validateErrs = helper.PhotoCreateValidator(request)
	if validateErrs != nil {
		errResponse := make([]string, len(validateErrs))
		for i, err := range validateErrs {
			errResponse[i] = err.Error()
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errResponse,
		})
		return
	}

	response, err := pc.photoService.Create(request, userID.(string))
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
		Message: "Photo created successfully",
		Data:    response,
	})
}

// GetAll godoc
// @summary get all photo
// @description get all photo data to the database
// @tags Photo
// @accept json
// @produce json
// @succes 200 {object} model.SuccessResponse{data=[]model.PhotoResponse}
// @failure	400	{object} model.ErrorResponse{errors=interface{}}
// @failure	500	{object} model.ErrorResponse{errors=interface{}}
// @security BearerAuth
// @router /photo [GET]
func (pc *PhotoController) GetAll(ctx *gin.Context) {
	response, err := pc.photoService.GetAll()
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
		Message: "Get all photo successfully",
		Data:    response,
	})
}

// GetOne godoc
// @summary get photo
// @description get on photo data from photo ID
// @tags Photo
// @param photoID path string true "photoID you want to retrieve"
// @accept json
// @produce json
// @succes 200 {object} model.SuccessResponse{data=model.PhotoResponse}
// @failure	400	{object} model.ErrorResponse{errors=interface{}}
// @failure	500	{object} model.ErrorResponse{errors=interface{}}
// @security BearerAuth
// @router /photo/{photoID} [GET]
func (pc *PhotoController) GetOne(ctx *gin.Context) {
	photoID := ctx.Param("photoID")

	response, err := pc.photoService.GetOne(photoID)
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
		Message: "Get photo successfully",
		Data:    response,
	})
}

// UpdatePhoto godoc
// @summary update photo
// @description update photo by photoID
// @tags Photo
// @param photoID path string true "photoID you want to retrieve"
// @param request body model.PhotoUpdateRequest true "request is mandatory"
// @accept json
// @produce json
// @succes 200 {object} model.SuccessResponse{data=model.PhotoUpdateResponse}
// @failure	400	{object} model.ErrorResponse{errors=interface{}}
// @failure	500	{object} model.ErrorResponse{errors=interface{}}
// @security BearerAuth
// @router /photo/{photoID} [PUT]
func (pc *PhotoController) UpdatePhoto(ctx *gin.Context) {
	var request model.PhotoUpdateRequest
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
			Errors: "UserID doesn't exist",
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helper.PhotoUpdateValidator(request)
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

	response, err := pc.photoService.UpdatePhoto(request, userID.(string), photoID)
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
		Message: "Photo updated successfully",
		Data: model.PhotoUpdateResponse{
			ID: response.ID,
		},
	})
}

// DeletePhoto godoc
// @summary delete a photo
// @description delete a photo by photoID
// @tags Photo
// @param photoID path string true "photoID you want to retrieve"
// @acccept json
// @produce json
// @succes 200 {object} model.SuccessResponse{data=model.PhotoDeleteResponse}
// @failure	400	{object} model.ErrorResponse{errors=interface{}}
// @failure	500	{object} model.ErrorResponse{errors=interface{}}
// @security BearerAuth
// @router /photo/{photoID} [DELETE]
func (pc *PhotoController) DeletePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photoID")

	userID, isExist := ctx.Get("userID")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	response, err := pc.photoService.Delete(photoID, userID.(string))
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
		Message: "Photo deleted successfully",
		Data: model.PhotoDeleteResponse{
			ID: response.ID,
		},
	})
}
