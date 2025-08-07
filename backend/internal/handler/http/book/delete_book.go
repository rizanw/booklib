package book

import "github.com/gofiber/fiber/v2"

func (h *Handler) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.usecase.DeleteBook(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	if err := h.usecase.DeleteBook(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
