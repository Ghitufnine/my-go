package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/ghitufnine/my-go/internal/domain/entity"
	"github.com/ghitufnine/my-go/internal/repository"
)

type CategoryUsecase struct {
	repo      repository.CategoryRepository
	cache     repository.CacheRepository
	publisher repository.EventPublisher
}

func NewCategoryUsecase(
	repo repository.CategoryRepository,
	cache repository.CacheRepository,
	publisher repository.EventPublisher,
) *CategoryUsecase {
	return &CategoryUsecase{
		repo:      repo,
		cache:     cache,
		publisher: publisher,
	}
}

func (u *CategoryUsecase) Create(
	ctx context.Context,
	name string,
) error {

	c := &entity.Category{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
	}

	err := u.repo.Create(ctx, c)
	if err != nil {
		return err
	}

	// invalidate cache
	_ = u.cache.Delete(ctx, "categories:all")

	// publish event
	_ = u.publisher.Publish(
		ctx,
		"category.created",
		[]byte(c.ID),
	)

	return nil
}

func (u *CategoryUsecase) Update(
	ctx context.Context,
	id string,
	name string,
) error {

	c, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if c == nil {
		return fmt.Errorf("category not found")
	}

	c.Name = name

	err = u.repo.Update(ctx, c)
	if err != nil {
		return err
	}

	_ = u.cache.Delete(ctx, "categories:all")

	_ = u.publisher.Publish(
		ctx,
		"category.updated",
		[]byte(id),
	)

	return nil
}

func (u *CategoryUsecase) Delete(
	ctx context.Context,
	id string,
) error {

	err := u.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	_ = u.cache.Delete(ctx, "categories:all")

	_ = u.publisher.Publish(
		ctx,
		"category.deleted",
		[]byte(id),
	)

	return nil
}

func (u *CategoryUsecase) GetAll(
	ctx context.Context,
) ([]entity.Category, error) {

	return u.repo.GetAll(ctx)
}
