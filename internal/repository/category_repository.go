package repository

import (
	"context"

	"github.com/ghitufnine/my-go/internal/domain/entity"
)

type CategoryRepository interface {
	Create(ctx context.Context, c *entity.Category) error
	Update(ctx context.Context, c *entity.Category) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]entity.Category, error)
	GetByID(ctx context.Context, id string) (*entity.Category, error)
}
