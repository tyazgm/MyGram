package model

import "time"

type Photo struct {
	ID        string `gorm:"primaryKey;type:varchar(255)"`
	Title     string `gorm:"not null;type:varchar(255)"`
	Caption   string `gorm:"not null;type:varchar(255)"`
	PhotoUrl  string `gorm:"not null;type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string
}

type PhotoCreateRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photoUrl" validate:"required"`
}

type PhotoCreateResponse struct {
	ID        string    `json:"photo_id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
