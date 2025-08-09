package book

import "github.com/gofiber/fiber/v2"

// GetAllBooks godoc
// @Summary Get all books
// @Description Returns a list of all books
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /books [get]
func (h *Handler) GetAllBooks(c *fiber.Ctx) error {
	books, err := h.usecase.GetAllBooks(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   books,
	})
}
