package main

import (
	"context"
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

	err := log.SetConfig(
		&log.Config{
			AppName:     appName,
			Environment: "dev",
		},
	)
	if err != nil {
		log.Fatal(ctx, err, nil, "failed to set log config")
	}

	srv := fiber.New(fiber.Config{
		AppName: appName,
	})

	routes(srv)

	if err := srv.Listen(":8080"); err != nil {
		log.Fatal(ctx, err, nil, "failed to start server")
	}
}
