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
