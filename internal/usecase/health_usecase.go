package usecase

import (
	"context"
	"fmt"

	"github.com/ghitufnine/my-go/internal/domain/entity"
	"github.com/ghitufnine/my-go/internal/repository"
)

type HealthUsecase struct {
	repo repository.HealthRepository
}

func NewHealthUsecase(repo repository.HealthRepository) *HealthUsecase {
	return &HealthUsecase{repo: repo}
}

func (u *HealthUsecase) Check(ctx context.Context) (*entity.Health, error) {
	result, err := u.repo.Check(ctx)
	if err != nil {
		return nil, fmt.Errorf("health check failed: %w", err)
	}
	return result, nil
}
