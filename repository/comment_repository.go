package repository

import (
	"MyGram/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (cr *CommentRepository) Create(comment model.Comment) error {
	err := cr.db.Create(&comment).Error
	if err != nil {
		return err
	}

	return nil
}

func (cr *CommentRepository) FindAll() ([]model.Comment, error) {
	comments := []model.Comment{}

	err := cr.db.Find(&comments).Error
	if err != nil {
		return []model.Comment{}, err
	}

	return comments, nil
}

func (cr *CommentRepository) FindByID(commentID string) (model.Comment, error) {
	comment := model.Comment{}

	err := cr.db.Debug().Where("comment_id = ?", commentID).Take(&comment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Comment{}, err
		}

		return model.Comment{}, err
	}

	return comment, nil
}

func (cr *CommentRepository) Update(comment model.Comment) error {
	err := cr.db.Save(&model.Comment{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		UpdatedAt: comment.UpdatedAt,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *CommentRepository) Delete(comment model.Comment) error {
	err := cr.db.Delete(&comment).Error
	if err != nil {
		return err
	}

	return nil
}

func (pr *CommentRepository) FindByPhotoID(photoID string) ([]model.Comment, error) {
	comments := []model.Comment{}

	err := pr.db.Where("photo_id = ?", photoID).Find(&comments).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Comment{}, err
		}

		return []model.Comment{}, err
	}

	fmt.Println("comments: ", comments)

	return comments, nil
}
