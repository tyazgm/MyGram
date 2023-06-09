package service

import (
	"MyGram/helper"
	"MyGram/model"
	"MyGram/repository"
	"errors"
	"time"
)

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

func (cs *CommentService) Create(commentReqData model.CommentCreateRequest, userID string, photoID string) (*model.CommentResponse, error) {
	_, err := cs.photoRepository.FindByID(photoID)
	if err != nil {
		return nil, err
	}

	commentID := helper.GenerateID()
	newComment := model.Comment{
		ID:        commentID,
		Message:   commentReqData.Message,
		PhotoID:   photoID,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = cs.commentRepository.Create(newComment)
	if err != nil {
		return nil, err
	}

	return &model.CommentResponse{
		ID:        newComment.ID,
		Message:   newComment.Message,
		PhotoID:   newComment.PhotoID,
		UserID:    newComment.UserID,
		CreatedAt: newComment.CreatedAt,
		UpdatedAt: newComment.UpdatedAt,
	}, nil
}

func (cs *CommentService) GetAll() ([]model.CommentResponse, error) {
	commentsResult, err := cs.commentRepository.FindAll()
	if err != nil {
		return []model.CommentResponse{}, err
	}

	commentsResponse := []model.CommentResponse{}
	for _, commentRes := range commentsResult {
		commentsResponse = append(commentsResponse, model.CommentResponse(commentRes))
	}

	return commentsResponse, nil
}

func (cs *CommentService) GetOne(commentID string) (model.CommentResponse, error) {
	commentResult, err := cs.commentRepository.FindByID(commentID)
	if err != nil {
		return model.CommentResponse{}, err
	}

	return model.CommentResponse(commentResult), nil
}

func (cs *CommentService) UpdateComment(commentReqData model.CommentUpdateRequest, userID string, commentID string) (*model.CommentResponse, error) {
	findCommentResponse, err := cs.commentRepository.FindByID(commentID)
	if err != nil {
		return nil, err
	}

	if userID != findCommentResponse.UserID {
		return nil, errors.New("Unauthorized")
	}

	updatedComment := model.Comment{
		ID:        findCommentResponse.ID,
		Message:   commentReqData.Message,
		PhotoID:   findCommentResponse.PhotoID,
		UserID:    userID,
		UpdatedAt: time.Now(),
	}

	err = cs.commentRepository.Update(updatedComment)
	if err != nil {
		return nil, err
	}

	return &model.CommentResponse{
		ID: commentID,
	}, nil
}

func (cs *CommentService) Delete(commentID string, userID string) (model.CommentDeleteResponse, error) {
	findCommentResponse, err := cs.commentRepository.FindByID(commentID)
	if err != nil {
		return model.CommentDeleteResponse{}, err
	}

	if userID != findCommentResponse.UserID {
		return model.CommentDeleteResponse{}, errors.New("Unauthorized")
	}

	err = cs.commentRepository.Delete(model.Comment{ID: commentID})
	if err != nil {
		return model.CommentDeleteResponse{}, err
	}

	return model.CommentDeleteResponse{
		ID: commentID,
	}, nil
}
