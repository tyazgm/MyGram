package service

import (
	"MyGram/helper"
	"MyGram/model"
	"MyGram/repository"
	"time"
)

type PhotoService struct {
	photoRepository   repository.PhotoRepository
	commentRepository repository.CommentRepository
}

func NewPhotoService(photoRepository repository.PhotoRepository, commentRepository repository.CommentRepository) *PhotoService {
	return &PhotoService{
		photoRepository:   photoRepository,
		commentRepository: commentRepository,
	}
}

func (ps *PhotoService) Create(photoCreateRequest model.PhotoCreateRequest, userID string) (*model.PhotoCreateResponse, error) {
	photoID := helper.GenerateID()
	newPhoto := model.Photo{
		ID:        photoID,
		Title:     photoCreateRequest.Title,
		Caption:   photoCreateRequest.Caption,
		PhotoUrl:  photoCreateRequest.PhotoUrl,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := ps.photoRepository.Create(newPhoto)
	if err != nil {
		return nil, err
	}

	return &model.PhotoCreateResponse{
		ID:        newPhoto.ID,
		Title:     newPhoto.Title,
		Caption:   newPhoto.Caption,
		PhotoUrl:  newPhoto.PhotoUrl,
		UserID:    newPhoto.UserID,
		CreatedAt: newPhoto.CreatedAt,
		UpdatedAt: newPhoto.UpdatedAt,
	}, nil
}

func (ps *PhotoService) GetAll() ([]model.PhotoResponse, error) {
	photosResult, err := ps.photoRepository.FindAll()
	if err != nil {
		return []model.PhotoResponse{}, err
	}

	photosResponse := []model.PhotoResponse{}
	for _, photoRes := range photosResult {
		photosResponse = append(photosResponse, model.PhotoResponse(photoRes))
	}

	return photosResponse, nil
}

func (ps *PhotoService) GetOne(photoID string) (model.PhotoResponse, error) {
	photosResult, err := ps.photoRepository.FindByID(photoID)
	if err != nil {
		return model.PhotoResponse{}, err
	}

	comments := []model.Comment{}
	commentsResponse, err := ps.commentRepository.FindByPhotoID(photoID)
	for _, comment := range commentsResponse {
		comments = append(comments, model.Comment(comment))
	}
	if err != nil {
		return model.PhotoResponse{}, err
	}

	return model.PhotoResponse{
		ID:        photosResult.ID,
		Title:     photosResult.Title,
		Caption:   photosResult.Caption,
		PhotoUrl:  photosResult.PhotoUrl,
		UserID:    photosResult.UserID,
		Comments:  comments,
		CreatedAt: photosResult.CreatedAt,
		UpdatedAt: photosResult.UpdatedAt,
	}, nil
}
