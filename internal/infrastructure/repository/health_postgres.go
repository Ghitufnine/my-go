package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ghitufnine/my-go/internal/domain/entity"
)

type HealthPostgresRepository struct {
	db *pgxpool.Pool
}

func NewHealthPostgresRepository(db *pgxpool.Pool) *HealthPostgresRepository {
	return &HealthPostgresRepository{
		db: db,
	}
}

func (r *HealthPostgresRepository) Check(ctx context.Context) (*entity.Health, error) {

	var result int

	err := r.db.QueryRow(ctx, "SELECT 1").Scan(&result)
	if err != nil {
		return nil, err
	}

	return &entity.Health{
		Status: "ok",
	}, nil
}
