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
