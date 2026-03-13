package container

import (
	"context"
	"log"

	"github.com/ghitufnine/my-go/internal/infrastructure/config"
	"github.com/ghitufnine/my-go/internal/infrastructure/logger"
	"github.com/ghitufnine/my-go/internal/infrastructure/mongo"
	"github.com/ghitufnine/my-go/internal/infrastructure/postgres"
	"github.com/ghitufnine/my-go/internal/infrastructure/redis"
	cache "github.com/ghitufnine/my-go/internal/infrastructure/redis_cache"
	"github.com/ghitufnine/my-go/internal/infrastructure/server"
	"go.uber.org/zap"
)

func Container() {

	ctx := context.Background()

	cfg := config.Load()

	logg, err := logger.New()
	if err != nil {
		log.Fatal(err)
	}
	defer logg.Sync()

	// Postgres
	pg, err := postgres.New(
		ctx,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPass,
		cfg.PostgresDB,
	)
	if err != nil {
		logg.Fatal("postgres connection failed", zap.Error(err))
	}

	// Mongo
	mongoDB, err := mongo.New(
		ctx,
		cfg.MongoURI,
		cfg.MongoDB,
	)
	if err != nil {
		logg.Fatal("mongo connection failed", zap.Error(err))
	}

	_ = mongoDB

	// Redis
	redisClient, err := redis.New(
		ctx,
		cfg.RedisAddr,
		cfg.RedisPassword,
		cfg.RedisDB,
	)
	if err != nil {
		logg.Fatal("redis connection failed", zap.Error(err))
	}

	redisCache := cache.NewRedisCache(redisClient)

	_ = redisCache

	// Fiber
	app := server.New()

	// Setup router + handlers
	SetupContainerServer(
		app,
		pg,
	)

	logg.Info("server starting on port " + cfg.AppPort)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
