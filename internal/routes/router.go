package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ghitufnine/my-go/internal/handler/http"
)

type Router struct {
	App           *fiber.App
	HealthHandler *http.HealthHandler
	AuthHandler   *http.AuthHandler
}

func NewRouter(
	app *fiber.App,
	healthHandler *http.HealthHandler,
	authHandler *http.AuthHandler,
) *Router {
	return &Router{
		App:           app,
		HealthHandler: healthHandler,
		AuthHandler:   authHandler,
	}
}

func (r *Router) Setup() {

	api := r.App.Group("/api")

	r.HealthHandler.Register(api)
	r.AuthHandler.RegisterRoutes(api)
}
