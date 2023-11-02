package models

import (
	"time"
)

type Order struct {
	ID           uint   `gorm:"primaryKey"`
	Customername string `gorm:"not null;type:varchar(191)"`
	OrderAt      string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
