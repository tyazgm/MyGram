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

func (ur *UserRepository) FindByUsername(username string) (model.User, error) {
	var user model.User
	err := ur.db.First(&user, "username = ?", username).Error

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
