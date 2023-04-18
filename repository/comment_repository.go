package repository

import (
	"MyGram/model"

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
