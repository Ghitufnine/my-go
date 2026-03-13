package middleware

import (
	"strings"

	"github.com/ghitufnine/my-go/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWT() fiber.Handler {

	return func(c *fiber.Ctx) error {

		auth := c.Get("Authorization")

		if auth == "" {
			return c.Status(401).JSON("missing token")
		}

		token := strings.TrimPrefix(auth, "Bearer ")

		userID, err := jwt.ParseToken(token)
		if err != nil {
			return c.Status(401).JSON("invalid token")
		}

		c.Locals("user_id", userID)

		return c.Next()
	}
}
