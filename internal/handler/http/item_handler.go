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

// create godoc
//
//	@Summary		Create item
//	@Description	Create a new item
//	@Tags			items
//	@Accept			json
//	@Produce		json
//	@Param			body	body	itemRequest	true	"Item request"
//	@Success		200	{string}	string	"created"
//	@Failure		400	{string}	string	"bad request"
//	@Failure		500	{string}	string	"internal server error"
//	@Security		BearerAuth
//	@Router			/items [post]
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

// getAll godoc
//
//	@Summary		List items
//	@Description	Get all items
//	@Tags			items
//	@Produce		json
//	@Success		200	{array}		entity.Item
//	@Failure		500	{string}	string	"internal server error"
//	@Security		BearerAuth
//	@Router			/items [get]
func (h *ItemHandler) getAll(c *fiber.Ctx) error {

	data, err := h.usecase.GetAll(c.Context())
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(data)
}

// getByID godoc
//
//	@Summary		Get item by ID
//	@Description	Get a single item by its ID
//	@Tags			items
//	@Produce		json
//	@Param			id	path	string	true	"Item ID"
//	@Success		200	{object}	entity.Item
//	@Failure		404	{string}	string	"not found"
//	@Failure		500	{string}	string	"internal server error"
//	@Security		BearerAuth
//	@Router			/items/{id} [get]
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

// update godoc
//
//	@Summary		Update item
//	@Description	Update an item by ID
//	@Tags			items
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string		true	"Item ID"
//	@Param			body	body	itemRequest	true	"Item request"
//	@Success		200	{string}	string	"updated"
//	@Failure		400	{string}	string	"bad request"
//	@Failure		500	{string}	string	"internal server error"
//	@Security		BearerAuth
//	@Router			/items/{id} [put]
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

// delete godoc
//
//	@Summary		Delete item
//	@Description	Delete an item by ID
//	@Tags			items
//	@Produce		json
//	@Param			id	path	string	true	"Item ID"
//	@Success		200	{string}	string	"deleted"
//	@Failure		500	{string}	string	"internal server error"
//	@Security		BearerAuth
//	@Router			/items/{id} [delete]
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
