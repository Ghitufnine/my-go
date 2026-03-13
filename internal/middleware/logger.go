package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Logger(log *zap.Logger) fiber.Handler {

	return func(c *fiber.Ctx) error {

		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		reqID := c.Locals("request_id")

		log.Info(
			"http_request",
			zap.String("request_id", toString(reqID)),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
		)

		return err
	}
}

func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	return v.(string)
}
