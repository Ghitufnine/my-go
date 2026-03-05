package server

import (
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "clean-arch-service",
	})
	return app
}
