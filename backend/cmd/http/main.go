package main

import (
	"context"
	"fmt"

	"booklib/internal/infra"
	"booklib/internal/infra/config"
	"github.com/gofiber/fiber/v2"
	"github.com/rizanw/go-log"
)

const (
	appName = "booklib"
)

func main() {
	var (
		ctx = context.Background()
	)

	// load config
	conf, err := config.New(appName)
	if err != nil {
		log.Fatal(ctx, err, nil, "failed to load config")
	}

	// build infra and resource
	resources, err := infra.NewResources(ctx, conf)
	if err != nil {
		log.Fatal(ctx, err, nil, "failed to build resources")
	}

	repo := newRepo(resources)
	uc := newUseCase(repo)

	srv := fiber.New(fiber.Config{
		AppName: appName,
	})
	routes(srv, uc)

	log.Infof(ctx, nil, nil, "⚡️server started on :%d", conf.Server.Port)
	if err = srv.Listen(fmt.Sprintf(":%d", conf.Server.Port)); err != nil {
		log.Fatal(ctx, err, nil, "failed to start server")
	}
}
