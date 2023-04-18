package service

import "MyGram/repository"

type CommentService struct {
	commentRepository repository.CommentRepository
	photoRepository   repository.PhotoRepository
}

func NewCommentService(commentRepository repository.CommentRepository, photoRepository repository.PhotoRepository) *CommentService {
	return &CommentService{
		commentRepository: commentRepository,
		photoRepository:   photoRepository,
	}
}
