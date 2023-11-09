package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Book struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UUID      string `gorm:"not null" json:uuid`
	Title     string `gorm:"not null" json:"title" form:"title" valid:"required~Title of book is required"`
	Author    string `gorm:"not null" json:"author" form:"author" valid:"required~Author of book is required"`
	Stock     int    `gorm:"not null" json:"stock" form:"stock" valid:"required~Stock of book is required, numeric~Invalid stock format"`
	UserID    uint
	User      *User
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(b)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (b *Book) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(b)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
