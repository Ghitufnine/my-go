package repository

import (
	"context"
	"time"

	"github.com/ghitufnine/my-go/internal/domain/entity"
	"github.com/ghitufnine/my-go/internal/infrastructure/cache_keys"
	"github.com/ghitufnine/my-go/internal/repository"
)

type CategoryCacheRepository struct {
	repo  repository.CategoryRepository
	cache repository.CacheRepository
}

func NewCategoryCacheRepository(
	repo repository.CategoryRepository,
	cache repository.CacheRepository,
) *CategoryCacheRepository {
	return &CategoryCacheRepository{
		repo:  repo,
		cache: cache,
	}
}

func (r *CategoryCacheRepository) Create(
	ctx context.Context,
	c *entity.Category,
) error {
	return r.repo.Create(ctx, c)
}

func (r *CategoryCacheRepository) Update(
	ctx context.Context,
	c *entity.Category,
) error {
	return r.repo.Update(ctx, c)
}

func (r *CategoryCacheRepository) Delete(
	ctx context.Context,
	id string,
) error {
	return r.repo.Delete(ctx, id)
}

func (r *CategoryCacheRepository) GetByID(
	ctx context.Context,
	id string,
) (*entity.Category, error) {
	return r.repo.GetByID(ctx, id)
}

func (r *CategoryCacheRepository) GetAll(
	ctx context.Context,
) ([]entity.Category, error) {

	var cached []entity.Category

	err := r.cache.Get(
		ctx,
		cache_keys.CategoriesAll,
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
		cache_keys.CategoriesAll,
		data,
		60*time.Second,
	)

	return data, nil
}
