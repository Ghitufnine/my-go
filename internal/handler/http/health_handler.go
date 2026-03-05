package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ghitufnine/my-go/internal/usecase"
	"github.com/ghitufnine/my-go/pkg/utils"
)

type HealthHandler struct {
	usecase *usecase.HealthUsecase
}

func NewHealthHandler(u *usecase.HealthUsecase) *HealthHandler {
	return &HealthHandler{usecase: u}
}

func (h *HealthHandler) Register(router fiber.Router) {
	router.Get("/health", h.check)
}

func (h *HealthHandler) check(c *fiber.Ctx) error {
	ctx := c.UserContext()

	result, err := h.usecase.Check(ctx)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, result)
}
