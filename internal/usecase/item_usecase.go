package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/ghitufnine/my-go/internal/domain/entity"
	"github.com/ghitufnine/my-go/internal/repository"
)

type ItemUsecase struct {
	repo         repository.ItemRepository
	categoryRepo repository.CategoryRepository
	cache        repository.CacheRepository
	publisher    repository.EventPublisher
}

func NewItemUsecase(
	repo repository.ItemRepository,
	categoryRepo repository.CategoryRepository,
	cache repository.CacheRepository,
	publisher repository.EventPublisher,
) *ItemUsecase {
	return &ItemUsecase{
		repo:         repo,
		categoryRepo: categoryRepo,
		cache:        cache,
		publisher:    publisher,
	}
}

func (u *ItemUsecase) Create(
	ctx context.Context,
	categoryID string,
	name string,
	price float64,
) error {

	cat, err := u.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		return err
	}

	if cat == nil {
		return fmt.Errorf("category not found")
	}

	item := &entity.Item{
		ID:         uuid.New().String(),
		CategoryID: categoryID,
		Name:       name,
		Price:      price,
		CreatedAt:  time.Now(),
	}

	err = u.repo.Create(ctx, item)
	if err != nil {
		return err
	}

	_ = u.cache.Delete(ctx, "items:all")

	_ = u.publisher.Publish(
		ctx,
		"item.created",
		[]byte(item.ID),
	)

	return nil
}

func (u *ItemUsecase) Update(
	ctx context.Context,
	id string,
	categoryID string,
	name string,
	price float64,
) error {

	item, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if item == nil {
		return fmt.Errorf("item not found")
	}

	item.Name = name
	item.Price = price
	item.CategoryID = categoryID

	err = u.repo.Update(ctx, item)
	if err != nil {
		return err
	}

	_ = u.cache.Delete(ctx, "items:all")

	_ = u.publisher.Publish(
		ctx,
		"item.updated",
		[]byte(id),
	)

	return nil
}

func (u *ItemUsecase) Delete(
	ctx context.Context,
	id string,
) error {

	err := u.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	_ = u.cache.Delete(ctx, "items:all")

	_ = u.publisher.Publish(
		ctx,
		"item.deleted",
		[]byte(id),
	)

	return nil
}

func (u *ItemUsecase) GetAll(
	ctx context.Context,
) ([]entity.Item, error) {

	return u.repo.GetAll(ctx)
}

func (u *ItemUsecase) GetByID(
	ctx context.Context,
	id string,
) (*entity.Item, error) {

	return u.repo.GetByID(ctx, id)
}
