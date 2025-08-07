package main

import "github.com/gofiber/fiber/v2"

func routes(srv *fiber.App) {
	api := srv.Group("/api")
	v1 := api.Group("/v1")

	srv.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	bookRoutes(v1)
}

func bookRoutes(router fiber.Router) {
	// TODO: TBD
}
