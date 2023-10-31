package models

import "time"

type Variant struct {
	ID          int
	VariantName string
	Quantity    int
	ProductID   int
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
