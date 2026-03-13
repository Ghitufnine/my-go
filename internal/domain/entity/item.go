package entity

import "time"

type Item struct {
	ID         string
	CategoryID string
	Name       string
	Price      float64
	CreatedAt  time.Time
}
