package repository

import (
	"context"
	"errors"

	"github.com/ghitufnine/my-go/internal/domain/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ItemPostgresRepository struct {
	db *pgxpool.Pool
}

func NewItemPostgresRepository(
	db *pgxpool.Pool,
) *ItemPostgresRepository {
	return &ItemPostgresRepository{
		db: db,
	}
}

func (r *ItemPostgresRepository) Create(
	ctx context.Context,
	i *entity.Item,
) error {

	query := `
	INSERT INTO items (
		id,
		category_id,
		name,
		price,
		created_at
	)
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		i.ID,
		i.CategoryID,
		i.Name,
		i.Price,
		i.CreatedAt,
	)

	return err
}

func (r *ItemPostgresRepository) Update(
	ctx context.Context,
	i *entity.Item,
) error {

	query := `
	UPDATE items
	SET name = $1,
	    price = $2,
	    category_id = $3
	WHERE id = $4
	`

	_, err := r.db.Exec(
		ctx,
		query,
		i.Name,
		i.Price,
		i.CategoryID,
		i.ID,
	)

	return err
}

func (r *ItemPostgresRepository) Delete(
	ctx context.Context,
	id string,
) error {

	query := `
	DELETE FROM items
	WHERE id = $1
	`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
	)

	return err
}

func (r *ItemPostgresRepository) GetAll(
	ctx context.Context,
) ([]entity.Item, error) {

	query := `
	SELECT id, category_id, name, price, created_at
	FROM items
	ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []entity.Item

	for rows.Next() {

		var i entity.Item

		err := rows.Scan(
			&i.ID,
			&i.CategoryID,
			&i.Name,
			&i.Price,
			&i.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, i)
	}

	return result, nil
}

func (r *ItemPostgresRepository) GetByID(
	ctx context.Context,
	id string,
) (*entity.Item, error) {

	query := `
	SELECT id, category_id, name, price, created_at
	FROM items
	WHERE id = $1
	`

	var i entity.Item

	err := r.db.QueryRow(ctx, query, id).
		Scan(
			&i.ID,
			&i.CategoryID,
			&i.Name,
			&i.Price,
			&i.CreatedAt,
		)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &i, nil
}
