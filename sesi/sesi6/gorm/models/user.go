package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"not null;unique;type:varchar(191)"`
	Books     []Book 
	CreatedAt time.Time
	UpdatedAt time.Time
}
