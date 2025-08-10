package urlprocessor

import (
	"errors"
	"github.com/rizanw/go-log"

	"github.com/gofiber/fiber/v2"
)

// ProcessUrlRequest represents the request payload for adding a book
type ProcessUrlRequest struct {
	Url       string `json:"url"`
	Operation string `json:"operation"`
}

func (req *ProcessUrlRequest) validate() error {
	if req.Url == "" {
		return errors.New("url cannot be empty")
	}
	if req.Operation == "" {
		return errors.New("operation cannot be empty")
	}
	return nil
}

// ProcessUrl godoc
// @Summary Clean and process a URL
// @Description Cleans or modifies a URL based on the specified operation.
// @Tags URLProcessor
// @Accept json
// @Produce json
// @Param request body ProcessUrlRequest true "URL processor request payload"
// @Success 201 {object} map[string]string "Processed URL returned"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /process-url [post]
func (h *Handler) ProcessUrl(c *fiber.Ctx) error {
	var req ProcessUrlRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if err := req.validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	res, err := h.usecase.CleanURL(c.UserContext(), req.Operation, req.Url)
	if err != nil {
		log.Error(c.UserContext(), err, nil, "failed to clean url")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"processed_url": res,
	})
}
