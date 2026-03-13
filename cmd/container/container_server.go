package container

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/ghitufnine/my-go/internal/handler/http"
	"github.com/ghitufnine/my-go/internal/infrastructure/rabbitmq"
	infraRepo "github.com/ghitufnine/my-go/internal/infrastructure/repository"
	"github.com/ghitufnine/my-go/internal/repository"
	"github.com/ghitufnine/my-go/internal/routes"
	"github.com/ghitufnine/my-go/internal/usecase"
)

func SetupContainerServer(
	app *fiber.App,
	pg *pgxpool.Pool,
	cache repository.CacheRepository,
	rabbit *rabbitmq.Rabbit,
	log *zap.Logger,
) {

	// Repositories
	healthRepo := infraRepo.NewHealthPostgresRepository(pg)
	userRepo := infraRepo.NewUserPostgresRepository(pg)
	refreshTokenRepo := infraRepo.NewRefreshTokenPostgresRepository(pg)
	publisher := rabbitmq.NewPublisher(rabbit)
	catRepoPG := infraRepo.NewCategoryPostgresRepository(pg)
	catRepo := infraRepo.NewCategoryCacheRepository(catRepoPG, cache)

	itemRepoPG := infraRepo.NewItemPostgresRepository(pg)
	itemRepo := infraRepo.NewItemCacheRepository(itemRepoPG, cache)

	// Usecases
	healthUC := usecase.NewHealthUsecase(healthRepo)
	authUC := usecase.NewAuthUsecase(userRepo, refreshTokenRepo)
	categoryUC := usecase.NewCategoryUsecase(
		catRepo,
		cache,
		publisher,
	)

	itemUC := usecase.NewItemUsecase(
		itemRepo,
		catRepo,
		cache,
		publisher,
	)

	// Handlers
	healthHandler := http.NewHealthHandler(healthUC)
	authHandler := http.NewAuthHandler(authUC)
	categoryHandler := http.NewCategoryHandler(categoryUC)
	itemHandler := http.NewItemHandler(itemUC)

	// Router
	router := routes.NewRouter(
		app,
		log,
		healthHandler,
		authHandler,
		categoryHandler,
		itemHandler,
	)

	router.Setup()
}
