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
