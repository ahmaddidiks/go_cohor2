package models

import (
	"time"
)

type Order struct {
	ID           int    `gorm:"primaryKey"`
	Customername string `gorm:"not null;unique;type:varchar(191)"`
	OrderAt      *time.Time
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

// func (b *Order) BeforeCreate(tx *gorm.DB) (err error) {
// 	fmt.Println("Book befroe create")

// 	return
// }
