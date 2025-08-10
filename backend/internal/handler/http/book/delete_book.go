package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizanw/go-log"
)

// DeleteBook godoc
// @Summary Delete a book by ID
// @Description Deletes a book with the given ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 204 "No Content"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{id} [delete]
func (h *Handler) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "id cannot be empty",
		})
	}

	if err := h.usecase.DeleteBook(c.UserContext(), id); err != nil {
		log.Error(c.UserContext(), err, nil, "failed to delete book")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
	})
}
