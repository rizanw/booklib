package book

import "github.com/gofiber/fiber/v2"

func (h *Handler) GetAllBooks(c *fiber.Ctx) error {
	books, err := h.usecase.GetAllBooks(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    books,
	})
}
