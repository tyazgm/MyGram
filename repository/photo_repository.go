package repository

import (
	"MyGram/model"
	"errors"

	"gorm.io/gorm"
)

type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{
		db: db,
	}
}

func (pr *PhotoRepository) FindByUserID(userID string) ([]model.Photo, error) {
	photos := []model.Photo{}

	err := pr.db.Where("user_id = ?", userID).Find(&photos).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Photo{}, err
		}

		return []model.Photo{}, err
	}

	return photos, nil
}
