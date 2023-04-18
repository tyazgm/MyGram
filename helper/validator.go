package helper

import (
	"MyGram/model"

	"github.com/go-playground/validator/v10"
)

func UserRegisterValidator(requestData model.UserRegisterRequest) []error {
	validate := validator.New()

	err := validate.Struct(requestData)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}

func UserLoginValidator(userLoginRequest model.UserLoginRequest) []error {
	validate := validator.New()

	err := validate.Struct(userLoginRequest)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}

func SocialCreateValidator(requestData model.SocialCreateRequest) []error {
	validate := validator.New()

	err := validate.Struct(requestData)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}

func SocialUpdateValidator(requestData model.SocialUpdateRequest) []error {
	validate := validator.New()

	err := validate.Struct(requestData)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}

func CommentCreateValidator(commentCreateRequest model.CommentCreateRequest) []error {
	validate := validator.New()

	err := validate.Struct(commentCreateRequest)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}

func CommentUpdateValidator(commentUpdateRequest model.CommentUpdateRequest) []error {
	validate := validator.New()

	err := validate.Struct(commentUpdateRequest)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}
