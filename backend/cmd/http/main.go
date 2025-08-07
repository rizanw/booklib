package main

import (
	"context"

	"booklib/internal/infra"
	"booklib/internal/infra/config"
	"github.com/gofiber/fiber/v2"
	"github.com/rizanw/go-log"
)

const (
	appName = "booklib-http"
)

func main() {
	var (
		ctx = context.Background()
	)

	// load config
	conf, err := config.New(appName)
	if err != nil {
		return
	}

	// build infra and resource
	resources, err := infra.NewResources(ctx, conf)
	if err != nil {
		return
	}

	repo := newRepo(resources)
	uc := newUseCase(repo)
	handler := newHandler(uc)

	srv := fiber.New(fiber.Config{
		AppName: appName,
	})
	routes(srv, handler)

	if err = srv.Listen(":8080"); err != nil {
		log.Fatal(ctx, err, nil, "failed to start server")
	}
}
