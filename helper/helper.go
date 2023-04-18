package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateID() string {
	return uuid.New().String()
}

func GenerateToken(userID string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})

	var tokenString string

	tokenString, err := jwtToken.SignedString([]byte("4l0h4m0r4"))

	return tokenString, err
}

func Hash(plain string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(plain), 8)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func PasswordIsMatch(plain string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	if err != nil {
		return false
	}

	return true
}
