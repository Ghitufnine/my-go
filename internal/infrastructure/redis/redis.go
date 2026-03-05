package redis

import (
	"context"
	"fmt"

	goredis "github.com/redis/go-redis/v9"
)

func New(ctx context.Context, addr string, password string, db int) (*goredis.Client, error) {

	client := goredis.NewClient(&goredis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis ping failed: %w", err)
	}

	return client, nil
}
