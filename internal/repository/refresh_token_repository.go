package repository

import (
	"context"
	"time"
)

type RefreshTokenRepository interface {
	Store(ctx context.Context, userID string, token string, expires time.Time) error
	Delete(ctx context.Context, token string) error
	Exists(ctx context.Context, token string) (bool, error)
}
