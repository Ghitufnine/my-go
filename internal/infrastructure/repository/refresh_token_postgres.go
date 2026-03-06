package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RefreshTokenPostgresRepository struct {
	db *pgxpool.Pool
}

func NewRefreshTokenPostgresRepository(db *pgxpool.Pool) *RefreshTokenPostgresRepository {
	return &RefreshTokenPostgresRepository{
		db: db,
	}
}

func (r *RefreshTokenPostgresRepository) Store(ctx context.Context, userID string, token string, expires time.Time) error {
	query := `
		INSERT INTO refresh_tokens (id, user_id, token, expires_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(ctx, query, uuid.New().String(), userID, token, expires)
	return err
}

func (r *RefreshTokenPostgresRepository) Delete(ctx context.Context, token string) error {
	query := `
		DELETE FROM refresh_tokens
		WHERE token = $1
	`
	_, err := r.db.Exec(ctx, query, token)
	return err
}

func (r *RefreshTokenPostgresRepository) Exists(ctx context.Context, token string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM refresh_tokens
			WHERE token = $1 AND expires_at > NOW()
		)
	`

	var exists bool
	err := r.db.QueryRow(ctx, query, token).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
