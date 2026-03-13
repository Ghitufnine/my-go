package repository

import (
	"context"
	"time"

	"github.com/ghitufnine/my-go/internal/domain/entity"
	"github.com/ghitufnine/my-go/internal/infrastructure/cache_keys"
	"github.com/ghitufnine/my-go/internal/repository"
)

type ItemCacheRepository struct {
	repo  repository.ItemRepository
	cache repository.CacheRepository
}

func NewItemCacheRepository(
	repo repository.ItemRepository,
	cache repository.CacheRepository,
) *ItemCacheRepository {
	return &ItemCacheRepository{
		repo:  repo,
		cache: cache,
	}
}

func (r *ItemCacheRepository) Create(
	ctx context.Context,
	i *entity.Item,
) error {
	return r.repo.Create(ctx, i)
}

func (r *ItemCacheRepository) Update(
	ctx context.Context,
	i *entity.Item,
) error {
	return r.repo.Update(ctx, i)
}

func (r *ItemCacheRepository) Delete(
	ctx context.Context,
	id string,
) error {
	return r.repo.Delete(ctx, id)
}

func (r *ItemCacheRepository) GetByID(
	ctx context.Context,
	id string,
) (*entity.Item, error) {
	return r.repo.GetByID(ctx, id)
}

func (r *ItemCacheRepository) GetAll(
	ctx context.Context,
) ([]entity.Item, error) {

	var cached []entity.Item

	err := r.cache.Get(
		ctx,
		cache_keys.ItemsAll,
		&cached,
	)

	if err == nil {
		return cached, nil
	}

	data, err := r.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	_ = r.cache.Set(
		ctx,
		cache_keys.ItemsAll,
		data,
		60*time.Second,
	)

	return data, nil
}
