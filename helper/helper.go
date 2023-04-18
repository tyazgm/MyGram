package helper

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateID() string {
	return uuid.New().String()
}

func Hash(plain string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(plain), 8)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
