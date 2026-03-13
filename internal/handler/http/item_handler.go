package http

import (
	"github.com/ghitufnine/my-go/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type ItemHandler struct {
	usecase *usecase.ItemUsecase
}

func NewItemHandler(u *usecase.ItemUsecase) *ItemHandler {
	return &ItemHandler{usecase: u}
}

func (h *ItemHandler) RegisterRoutes(router fiber.Router) {

	r := router.Group("/items")

	r.Post("/", h.create)
	r.Get("/", h.getAll)
	r.Get("/:id", h.getByID)
	r.Put("/:id", h.update)
	r.Delete("/:id", h.delete)
}

type itemRequest struct {
	Name       string  `json:"name"`
	CategoryID string  `json:"category_id"`
	Price      float64 `json:"price"`
}

func (h *ItemHandler) create(c *fiber.Ctx) error {

	var req itemRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	err := h.usecase.Create(
		c.Context(),
		req.CategoryID,
		req.Name,
		req.Price,
	)

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("created")
}

func (h *ItemHandler) getAll(c *fiber.Ctx) error {

	data, err := h.usecase.GetAll(c.Context())
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(data)
}

func (h *ItemHandler) getByID(c *fiber.Ctx) error {

	id := c.Params("id")

	data, err := h.usecase.GetByID(
		c.Context(),
		id,
	)

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	if data == nil {
		return c.Status(404).JSON("not found")
	}

	return c.JSON(data)
}

func (h *ItemHandler) update(c *fiber.Ctx) error {

	id := c.Params("id")

	var req itemRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	err := h.usecase.Update(
		c.Context(),
		id,
		req.CategoryID,
		req.Name,
		req.Price,
	)

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON("updated")
}

func (h *ItemHandler) delete(c *fiber.Ctx) error {

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
