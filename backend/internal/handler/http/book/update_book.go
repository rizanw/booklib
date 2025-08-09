package book

import (
	"booklib/internal/usecase/book"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type UpdateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func (req *UpdateBookRequest) parseValidateRequest() (book.UpdateBookInput, error) {
	if req.Title == "" {
		return book.UpdateBookInput{}, errors.New("title cannot be empty")
	}
	if req.Author == "" {
		return book.UpdateBookInput{}, errors.New("author cannot be empty")
	}
	if req.Year == 0 {
		return book.UpdateBookInput{}, errors.New("year cannot be empty")
	}

	return book.UpdateBookInput{
		Title:  req.Title,
		Author: req.Author,
		Year:   req.Year,
	}, nil
}

func (h *Handler) UpdateBook(c *fiber.Ctx) error {
	var req UpdateBookRequest

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "id cannot be empty",
		})
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Cannot parse JSON",
		})
	}

	in, err := req.parseValidateRequest()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err,
		})
	}

	if err = h.usecase.UpdateBook(c.Context(), id, in); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
	})
}
