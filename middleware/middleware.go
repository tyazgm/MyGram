package middleware

import (
	"MyGram/helper"
	"MyGram/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")

	bearerIsExist := strings.Contains(auth, "Bearer")
	if !bearerIsExist {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: "Unauthorized",
		})
		return
	}

	token := strings.Split(auth, " ")[1]

	jwtToken, err := helper.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	ctx.Set("username", claims["username"])
	ctx.Set("userID", claims["user_id"])
	ctx.Next()
}
