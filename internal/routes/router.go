package routes

import (
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"go.uber.org/zap"

	"github.com/ghitufnine/my-go/internal/handler/http"
	"github.com/ghitufnine/my-go/internal/middleware"
)

type Router struct {
	App             *fiber.App
	Log             *zap.Logger
	HealthHandler   *http.HealthHandler
	AuthHandler     *http.AuthHandler
	CategoryHandler *http.CategoryHandler
	ItemHandler     *http.ItemHandler
}

func NewRouter(
	app *fiber.App,
	log *zap.Logger,
	healthHandler *http.HealthHandler,
	authHandler *http.AuthHandler,
	categoryHandler *http.CategoryHandler,
	itemHandler *http.ItemHandler,
) *Router {
	return &Router{
		App:             app,
		Log:             log,
		HealthHandler:   healthHandler,
		AuthHandler:     authHandler,
		CategoryHandler: categoryHandler,
		ItemHandler:     itemHandler,
	}
}

func (r *Router) Setup() {
	r.App.Use(middleware.RequestID())
	r.App.Use(middleware.Logger(r.Log))

	// Swagger UI
	r.App.Get("/swagger/*", fiberSwagger.WrapHandler)

	api := r.App.Group("/api")

	// public
	r.HealthHandler.Register(api)
	r.AuthHandler.RegisterRoutes(api)

	// protected
	protected := api.Group(
		"/",
		middleware.JWT(),
	)

	r.CategoryHandler.RegisterRoutes(protected)
	r.ItemHandler.RegisterRoutes(protected)
}
