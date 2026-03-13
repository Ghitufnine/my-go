package container

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ghitufnine/my-go/internal/handler/http"
	infraRepo "github.com/ghitufnine/my-go/internal/infrastructure/repository"
	"github.com/ghitufnine/my-go/internal/routes"
	"github.com/ghitufnine/my-go/internal/usecase"
)

func SetupContainerServer(
	app *fiber.App,
	pg *pgxpool.Pool,
) {

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
}
