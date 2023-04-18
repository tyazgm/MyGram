package model

import "time"

type SocialMedia struct {
	ID             string `gorm:"primaryKey;type:varchar(255)"`
	Name           string `gorm:"not null;type:varchar(255)"`
	SocialMediaUrl string `gorm:"not null;type:varchar(255)"`
	UserID         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type SocialCreateRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
}

type SocialResponse struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         string    `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialUpdateRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
}

type SocialUpdateResponse struct {
	ID string `json:"id"`
}
