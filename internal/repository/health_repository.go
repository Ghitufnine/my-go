package repository

import (
	"context"

	"github.com/ghitufnine/my-go/internal/domain/entity"
)

type HealthRepository interface {
	Check(ctx context.Context) (*entity.Health, error)
}
