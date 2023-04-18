package service

import "MyGram/repository"

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
