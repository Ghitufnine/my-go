package entity

import "time"

type Category struct {
	ID        string
	Name      string
	CreatedAt time.Time
}
