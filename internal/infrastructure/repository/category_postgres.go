package repository

import (
	"context"
	"errors"

	"github.com/ghitufnine/my-go/internal/domain/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryPostgresRepository struct {
	db *pgxpool.Pool
}

func NewCategoryPostgresRepository(db *pgxpool.Pool) *CategoryPostgresRepository {
	return &CategoryPostgresRepository{
		db: db,
	}
}

func (r *CategoryPostgresRepository) Create(
	ctx context.Context,
	c *entity.Category,
) error {

	query := `
	INSERT INTO categories (id, name, created_at)
	VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		c.ID,
		c.Name,
		c.CreatedAt,
	)

	return err
}

func (r *CategoryPostgresRepository) Update(
	ctx context.Context,
	c *entity.Category,
) error {

	query := `
	UPDATE categories
	SET name = $1
	WHERE id = $2
	`

	_, err := r.db.Exec(
		ctx,
		query,
		c.Name,
		c.ID,
	)

	return err
}

func (r *CategoryPostgresRepository) Delete(
	ctx context.Context,
	id string,
) error {

	query := `
	DELETE FROM categories
	WHERE id = $1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
	)

	return err
}

func (r *CategoryPostgresRepository) GetAll(
	ctx context.Context,
) ([]entity.Category, error) {

	query := `
	SELECT id, name, created_at
	FROM categories
	ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []entity.Category

	for rows.Next() {

		var c entity.Category

		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, c)
	}

	return result, nil
}

func (r *CategoryPostgresRepository) GetByID(
	ctx context.Context,
	id string,
) (*entity.Category, error) {

	query := `
	SELECT id, name, created_at
	FROM categories
	WHERE id = $1
	`

	var c entity.Category

	err := r.db.QueryRow(ctx, query, id).
		Scan(
			&c.ID,
			&c.Name,
			&c.CreatedAt,
		)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &c, nil
}
