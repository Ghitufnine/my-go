package repository

import (
	"context"

	"github.com/ghitufnine/my-go/internal/domain/entity"
)

type ItemRepository interface {
	Create(ctx context.Context, i *entity.Item) error
	Update(ctx context.Context, i *entity.Item) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]entity.Item, error)
	GetByID(ctx context.Context, id string) (*entity.Item, error)
}
