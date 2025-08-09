package main

import (
	_ "booklib/docs"
	hbook "booklib/internal/handler/http/book"
	hurlprocessor "booklib/internal/handler/http/url-processor"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func routes(srv *fiber.App, uc *UseCase) {
	srv.Get("/docs/*", swagger.HandlerDefault)

	srv.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("PONG!!")
	})

	api := srv.Group("/api")
	v1 := api.Group("/v1")

	bookRoutes(v1, uc)
	urlProcessorRoutes(v1, uc)
}

func urlProcessorRoutes(router fiber.Router, uc *UseCase) {
	handler := hurlprocessor.New(uc.UrlProcessor)

	router.Post("/process-url", handler.ProcessUrl)
}

func bookRoutes(router fiber.Router, uc *UseCase) {
	handler := hbook.New(uc.Book)

	router.Get("books", handler.GetAllBooks)
	router.Get("books/:id", handler.GetBook)
	router.Post("books", handler.AddBook)
	router.Put("books/:id", handler.UpdateBook)
	router.Delete("books/:id", handler.DeleteBook)
}
