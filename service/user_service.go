package service

import (
	"MyGram/helper"
	"MyGram/model"
	"MyGram/repository"
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
