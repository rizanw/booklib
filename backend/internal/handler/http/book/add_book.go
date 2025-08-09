package book

import (
	"booklib/internal/usecase/book"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type AddBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func (req *AddBookRequest) parseValidateRequest() (book.AddBookInput, error) {
	if req.Title == "" {
		return book.AddBookInput{}, errors.New("title cannot be empty")
	}
	if req.Author == "" {
		return book.AddBookInput{}, errors.New("author cannot be empty")
	}
	if req.Year == 0 {
		return book.AddBookInput{}, errors.New("year cannot be empty")
	}

	return book.AddBookInput{
		Title:  req.Title,
		Author: req.Author,
		Year:   req.Year,
	}, nil
}

func (h *Handler) AddBook(c *fiber.Ctx) error {
	var req AddBookRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	in, err := req.parseValidateRequest()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	if err := h.usecase.AddBook(c.Context(), in); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
	})
}
