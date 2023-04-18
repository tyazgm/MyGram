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

func (pr *PhotoRepository) FindByID(photoID string) (model.Photo, error) {
	photo := model.Photo{}

	err := pr.db.Debug().Where("photo_id = ?", photoID).Take(&photo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Photo{}, err
		}

		return model.Photo{}, err
	}

	return photo, nil
}

func (pr *PhotoRepository) Create(photo model.Photo) error {
	err := pr.db.Create(&photo).Error
	if err != nil {
		return err
	}

	return nil
}

func (pr *PhotoRepository) FindAll() ([]model.Photo, error) {
	photos := []model.Photo{}

	err := pr.db.Find(&photos).Error
	if err != nil {
		return []model.Photo{}, err
	}

	return photos, nil
}
