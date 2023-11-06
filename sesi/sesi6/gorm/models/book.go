package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Author    string `gorm:"not null"`
	Stock     int    `gorm:"not null"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Book befroe create")

	if len(b.Title) < 3 {
		err = errors.New("Book title is too sort")
	}
	return
}
