package entity

import "time"

type TransactionLog struct {
	ID        string
	Topic     string
	Payload   string
	CreatedAt time.Time
}
