package book

import "github.com/gofiber/fiber/v2"

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
