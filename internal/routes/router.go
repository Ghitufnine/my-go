package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ghitufnine/my-go/internal/handler/http"
	"github.com/ghitufnine/my-go/internal/middleware"
)

type Router struct {
	App             *fiber.App
	HealthHandler   *http.HealthHandler
	AuthHandler     *http.AuthHandler
	CategoryHandler *http.CategoryHandler
	ItemHandler     *http.ItemHandler
}

func NewRouter(
	app *fiber.App,
	healthHandler *http.HealthHandler,
	authHandler *http.AuthHandler,
	categoryHandler *http.CategoryHandler,
	itemHandler *http.ItemHandler,
) *Router {
	return &Router{
		App:             app,
		HealthHandler:   healthHandler,
		AuthHandler:     authHandler,
		CategoryHandler: categoryHandler,
		ItemHandler:     itemHandler,
	}
}

func (r *Router) Setup() {

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
