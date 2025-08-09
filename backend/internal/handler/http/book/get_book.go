package book

import "github.com/gofiber/fiber/v2"

// GetBook godoc
// @Summary Get a book by ID
// @Description Returns a single book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]string
// @Router /books/{id} [get]
func (h *Handler) GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "id cannot be empty",
		})
	}

	res, err := h.usecase.GetBook(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    res,
	})
}
