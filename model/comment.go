package model

import "time"

type Comment struct {
	ID        string `gorm:"primaryKey;type:varchar(255)"`
	Message   string `gorm:"not null;type:varchar(255)"`
	UserID    string
	PhotoID   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
