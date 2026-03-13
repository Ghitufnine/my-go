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

// create godoc
//
//	@Summary		Create category
//	@Description	Create a new category
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			body	body	categoryRequest	true	"Category request"
//	@Success		200	{string}	string	"created"
//	@Failure		400	{string}	string	"bad request"
//	@Failure		500	{string}	string	"internal server error"
//	@Security		BearerAuth
//	@Router			/categories [post]
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

// getAll godoc
//
//	@Summary		List categories
//	@Description	Get all categories
//	@Tags			categories
//	@Produce		json
//	@Success		200	{array}		entity.Category
//	@Failure		500	{string}	string	"internal server error"
//	@Security		BearerAuth
//	@Router			/categories [get]
func (h *CategoryHandler) getAll(c *fiber.Ctx) error {

	data, err := h.usecase.GetAll(c.Context())
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(data)
}

// update godoc
//
//	@Summary		Update category
//	@Description	Update a category by ID
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string			true	"Category ID"
//	@Param			body	body	categoryRequest	true	"Category request"
//	@Success		200	{string}	string	"updated"
//	@Failure		400	{string}	string	"bad request"
//	@Failure		500	{string}	string	"internal server error"
//	@Security		BearerAuth
//	@Router			/categories/{id} [put]
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

// delete godoc
//
//	@Summary		Delete category
//	@Description	Delete a category by ID
//	@Tags			categories
//	@Produce		json
//	@Param			id	path	string	true	"Category ID"
//	@Success		200	{string}	string	"deleted"
//	@Failure		500	{string}	string	"internal server error"
//	@Security		BearerAuth
//	@Router			/categories/{id} [delete]
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
