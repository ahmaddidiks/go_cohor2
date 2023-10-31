package models

import "time"

type Product struct {
	ID int
	Name string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}