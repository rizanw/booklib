package main

import (
	hbook "booklib/internal/handler/http/book"
	"github.com/gofiber/fiber/v2"
)

func routes(srv *fiber.App, uc *UseCase) {
	api := srv.Group("/api")
	v1 := api.Group("/v1")

	srv.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	bookRoutes(v1, uc)
}

func bookRoutes(router fiber.Router, uc *UseCase) {
	handler := hbook.New(uc.Book)

	router.Get("books", handler.GetAllBooks)
	router.Get("books/:id", handler.GetBook)
	router.Post("books", handler.AddBook)
	router.Put("books/:id", handler.UpdateBook)
	router.Delete("books/:id", handler.DeleteBook)
}
