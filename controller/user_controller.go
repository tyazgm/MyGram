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

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Register godoc
// @summary register account
// @description register your account to MyGram App
// @tags User
// @accept json
// @produce json
// @param data body model.UserRegisterRequest true "data is mandatory"
// @succes 200 {object} model.SuccessResponse
// @router /register [POST]
func (uc *UserController) Register(ctx *gin.Context) {
	request := model.UserRegisterRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helper.UserRegisterValidator(request)
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

	response, err := uc.userService.Register(request)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: "Bad Request",
				Errors: errors.New("This email or username is already used").Error(),
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
		Message: "Register success!",
		Data:    response,
	})
}

// Login godoc
// @summary login account
// @description login to MyGram App using registered account
// @tags User
// @accept json
// @produce json
// @param data body model.UserLoginRequest true "data is mandatory"
// @succes 200 {object} model.SuccessResponse
// @router /login [POST]
func (uc *UserController) Login(ctx *gin.Context) {
	request := model.UserLoginRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helper.UserLoginValidator(request)
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

	response, err := uc.userService.Login(request)
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
		Message: "Login success!",
		Data: model.UserLoginResponse{
			Token: *response,
		},
	})
}

// GetProfile godoc
// @summary get user profile
// @description get user profile using token info
// @tags User
// @param userID path string true "userID you want to retrieve"
// @produce json
// @succes 200 {object} model.UserProfileResponse
// @router /user/profile [GET]
func (uc *UserController) GetProfile(ctx *gin.Context) {
	userID, isExist := ctx.Get("userID")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	user, err := uc.userService.GetProfile(userID.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.UserProfileResponse{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		Age:          user.Age,
		Photos:       user.Photos,
		SocialMedias: user.SocialMedias,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	})
}
