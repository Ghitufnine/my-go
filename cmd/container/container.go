package container

import (
	"context"
	"log"

	"github.com/ghitufnine/my-go/internal/handler/http"
	"github.com/ghitufnine/my-go/internal/infrastructure/config"
	"github.com/ghitufnine/my-go/internal/infrastructure/logger"
	"github.com/ghitufnine/my-go/internal/infrastructure/mongo"
	"github.com/ghitufnine/my-go/internal/infrastructure/postgres"
	"github.com/ghitufnine/my-go/internal/infrastructure/redis"
	cache "github.com/ghitufnine/my-go/internal/infrastructure/redis_cache"
	infraRepo "github.com/ghitufnine/my-go/internal/infrastructure/repository"
	"github.com/ghitufnine/my-go/internal/infrastructure/server"
	"github.com/ghitufnine/my-go/internal/routes"
	"github.com/ghitufnine/my-go/internal/usecase"
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

	// PostgreSQL
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

	// MongoDB
	mongoDB, err := mongo.New(
		ctx,
		cfg.MongoURI,
		cfg.MongoDB,
	)
	if err != nil {
		logg.Fatal("mongo connection failed", zap.Error(err))
	}

	_ = mongoDB // used later

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

	// Fiber server
	app := server.New()

	// Repositories
	healthRepo := infraRepo.NewHealthPostgresRepository(pg)
	userRepo := infraRepo.NewUserPostgresRepository(pg)
	refreshTokenRepo := infraRepo.NewRefreshTokenPostgresRepository(pg)

	// Usecases
	healthUC := usecase.NewHealthUsecase(healthRepo)
	authUC := usecase.NewAuthUsecase(userRepo, refreshTokenRepo)

	// Handlers
	healthHandler := http.NewHealthHandler(healthUC)
	authHandler := http.NewAuthHandler(authUC)

	// Router
	router := routes.NewRouter(
		app,
		healthHandler,
		authHandler,
	)

	router.Setup()

	logg.Info("server starting on port " + cfg.AppPort)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
