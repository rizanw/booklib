package middleware

import (
	"bytes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rizanw/go-log"
)

func AccessLogMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// --- Request ID ---
		reqID := c.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}
		c.Set("X-Request-ID", reqID)
		c.SetUserContext(log.SetCtxRequestID(c.UserContext(), reqID))

		// --- Request data ---
		queryParams := c.Queries()
		routeParams := c.AllParams()
		reqBody := string(c.Body())

		// Log request
		log.Infof(
			c.UserContext(),
			nil,
			log.KV{
				"ip":      c.IP(),
				"ua":      c.Get("User-Agent"),
				"query":   queryParams,
				"route":   routeParams,
				"request": reqBody,
			},
			"[REQ] %s %s %s",
			start.Format("2006-01-02 15:04:05"),
			c.Method(),
			c.OriginalURL(),
		)

		// --- Capture response ---
		var resBody bytes.Buffer
		c.Response().SetBodyStream(&resBody, -1)

		// Process the request
		err := c.Next()

		// --- Response data ---
		duration := time.Since(start)
		status := c.Response().StatusCode()

		result := "SUCCESS"
		if status >= 400 {
			result = "ERROR"
		}

		// Log response
		log.Infof(
			c.UserContext(),
			nil,
			log.KV{
				"ip":       c.IP(),
				"ua":       c.Get("User-Agent"),
				"duration": duration,
				"status":   status,
				"response": resBody.String(),
				"error":    err,
			},
			"[RES] %s %s %s %d %s",
			start.Format("2006-01-02 15:04:05"),
			c.Method(),
			c.OriginalURL(),
			status,
			result,
		)

		return err
	}
}
