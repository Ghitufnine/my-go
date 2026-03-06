package http

import (
	"strings"

	"github.com/ghitufnine/my-go/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(u *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

func (h *AuthHandler) RegisterRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/register", h.register)
	auth.Post("/login", h.login)
	auth.Post("/logout", h.logout)
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) register(c *fiber.Ctx) error {

	var req registerRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	err := h.usecase.Register(c.Context(), req.Email, req.Password)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("registered")
}

func (h *AuthHandler) login(c *fiber.Ctx) error {

	var req registerRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	access, refresh, err := h.usecase.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return c.Status(401).JSON(err.Error())
	}

	return c.JSON(fiber.Map{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func (h *AuthHandler) logout(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON("missing authorization header")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	err := h.usecase.Logout(c.Context(), token)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("logout success")
}
