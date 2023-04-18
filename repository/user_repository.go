package repository

import (
	"MyGram/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user model.User) error {
	err := ur.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}
