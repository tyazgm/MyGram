package model

import "time"

type User struct {
	ID           string `gorm:"primaryKey;type:varchar(255)"`
	Username     string `gorm:"not null;uniqueIndex;type:varchar(255)"`
	Email        string `gorm:"not null;uniqueIndex;type:varchar(255)"`
	Password     string `gorm:"not null;type:varchar(255)"`
	Age          int    `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Photos       []Photo
	SocialMedias []SocialMedia
	Comments     []Comment
}

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Age      int    `json:"age" validate:"required,gt=8"`
}

type UserRegisterResponse struct {
	ID        string    `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
