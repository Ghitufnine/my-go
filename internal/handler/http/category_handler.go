package http

import (
	"github.com/ghitufnine/my-go/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	usecase *usecase.CategoryUsecase
}

func NewCategoryHandler(u *usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{usecase: u}
}

func (h *CategoryHandler) RegisterRoutes(router fiber.Router) {

	r := router.Group("/categories")

	r.Post("/", h.create)
	r.Get("/", h.getAll)
	r.Put("/:id", h.update)
	r.Delete("/:id", h.delete)
}

type categoryRequest struct {
	Name string `json:"name"`
}

func (h *CategoryHandler) create(c *fiber.Ctx) error {

	var req categoryRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	err := h.usecase.Create(
		c.Context(),
		req.Name,
	)

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("created")
}

func (h *CategoryHandler) getAll(c *fiber.Ctx) error {

	data, err := h.usecase.GetAll(c.Context())
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(data)
}

func (h *CategoryHandler) update(c *fiber.Ctx) error {

	id := c.Params("id")

	var req categoryRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	err := h.usecase.Update(
		c.Context(),
		id,
		req.Name,
	)

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("updated")
}

func (h *CategoryHandler) delete(c *fiber.Ctx) error {

	id := c.Params("id")

	err := h.usecase.Delete(
		c.Context(),
		id,
	)

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("deleted")
}
