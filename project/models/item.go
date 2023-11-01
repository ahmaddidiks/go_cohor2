package models

import (
	"time"
)

type Item struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Quantity    int    `gorm:"not null"`
	OrderID     int    `gorm:"not null"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type ItemOrderRecv struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type ItemRecv struct {
	OrderAt      string          `json:"orderedAt"`
	CustomerName string          `json:"customerName"`
	Items        []ItemOrderRecv `json:"items"`
}

// func (b *Item) BeforeCreate(tx *gorm.DB) (err error) {
// 	fmt.Println("Book befroe create")
// 	return
// }
