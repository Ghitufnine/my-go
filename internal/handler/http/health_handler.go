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

// check godoc
//
//	@Summary		Health check
//	@Description	Returns the health status of the service and its dependencies
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"success: true, data: {status}"
//	@Failure		500	{object}	map[string]interface{}	"success: false, error: message"
//	@Router			/health [get]
func (h *HealthHandler) check(c *fiber.Ctx) error {
	ctx := c.UserContext()

	result, err := h.usecase.Check(ctx)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, result)
}
