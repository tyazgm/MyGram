package service

import (
	"MyGram/helper"
	"MyGram/model"
	"MyGram/repository"
	"errors"
	"fmt"
)

type UserService struct {
	userRepository   repository.UserRepository
	photoRepository  repository.PhotoRepository
	socialRepository repository.SocialRepository
}

func NewUserService(userRepository repository.UserRepository, photoRepository repository.PhotoRepository, socialRepository repository.SocialRepository) *UserService {
	return &UserService{
		userRepository:   userRepository,
		photoRepository:  photoRepository,
		socialRepository: socialRepository,
	}
}

func (us *UserService) Register(userRegisterRequest model.UserRegisterRequest) (*model.UserRegisterResponse, error) {
	userID := helper.GenerateID()
	hashedPassword, err := helper.Hash(userRegisterRequest.Password)
	if err != nil {
		return nil, err
	}

	newUser := model.User{
		ID:       userID,
		Username: userRegisterRequest.Username,
		Email:    userRegisterRequest.Email,
		Password: hashedPassword,
		Age:      userRegisterRequest.Age,
	}

	err = us.userRepository.Create(newUser)
	if err != nil {
		return nil, err
	}

	return &model.UserRegisterResponse{
		ID:       newUser.ID,
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: newUser.Password,
		Age:      newUser.Age,
	}, nil
}

func (us *UserService) Login(userLoginRequest model.UserLoginRequest) (*string, error) {
	userResponse, err := us.userRepository.FindByUsername(userLoginRequest.Username)
	if err != nil {
		return nil, err
	}

	isMatch := helper.PasswordIsMatch(userLoginRequest.Password, userResponse.Password)
	if isMatch == false {
		return nil, errors.New(fmt.Sprintf("Invalid username or password"))
	}

	token, err := helper.GenerateToken(userResponse.ID)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (us *UserService) GetProfile(userID string) (model.User, error) {
	user, err := us.userRepository.FindByID(userID)
	if err != nil {
		return model.User{}, err
	}

	photos, err := us.photoRepository.FindByUserID(userID)
	if err != nil {
		return model.User{}, err
	}

	socials, err := us.socialRepository.FindByUserID(userID)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:           userID,
		Username:     user.Username,
		Email:        user.Email,
		Age:          user.Age,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Photos:       photos,
		SocialMedias: socials,
	}, nil
}
