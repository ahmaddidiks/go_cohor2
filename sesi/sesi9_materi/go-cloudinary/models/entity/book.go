package entity

import "time"

type Book struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Title     string     `gorm:"not null" json:"title"`
	Author    string     `gorm:"not null" json:"author"`
	Stock     int        `gorm:"not null" json:"stock"`
	ImageURL  string     `json:"image_url"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
