package repository

import (
	"MyGram/model"
	"errors"

	"gorm.io/gorm"
)

type SocialRepository struct {
	db *gorm.DB
}

func NewSocialRepository(db *gorm.DB) *SocialRepository {
	return &SocialRepository{
		db: db,
	}
}

func (sr *SocialRepository) FindByUserID(userID string) ([]model.SocialMedia, error) {
	socials := []model.SocialMedia{}

	err := sr.db.Debug().Where("user_id = ?", userID).Find(&socials).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.SocialMedia{}, err
		}

		return []model.SocialMedia{}, err
	}

	return socials, nil
}

func (sr *SocialRepository) Create(socialMedia model.SocialMedia) error {
	err := sr.db.Create(&socialMedia).Error
	if err != nil {
		return err
	}

	return nil
}

func (sr *SocialRepository) FindAll() ([]model.SocialMedia, error) {
	socials := []model.SocialMedia{}

	err := sr.db.Find(&socials).Error
	if err != nil {
		return []model.SocialMedia{}, err
	}

	return socials, nil
}

func (sr *SocialRepository) FindByID(socialMediaID string) (model.SocialMedia, error) {
	social := model.SocialMedia{}

	err := sr.db.Debug().Where("id = ?", socialMediaID).Take(&social).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SocialMedia{}, err
		}

		return model.SocialMedia{}, err
	}

	return social, nil
}
