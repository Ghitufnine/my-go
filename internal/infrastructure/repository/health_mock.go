package repository

import (
	"context"

	"github.com/ghitufnine/my-go/internal/domain/entity"
)

type HealthMockRepository struct{}

func NewHealthMockRepository() *HealthMockRepository {
	return &HealthMockRepository{}
}

func (r *HealthMockRepository) Check(ctx context.Context) (*entity.Health, error) {
	return &entity.Health{
		Status: "ok",
	}, nil
}
