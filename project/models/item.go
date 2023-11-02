package models

import "time"

type Item struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Quantity    int    `gorm:"not null"`
	OrderID     uint
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
