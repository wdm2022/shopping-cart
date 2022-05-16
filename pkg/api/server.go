package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RunHttpServer(port *int, prefork *bool) error {
	server := fiber.New(fiber.Config{
		Prefork: *prefork,
	})

	// Middleware
	server.Use(recover.New())
	server.Use(logger.New())

	return server.Listen(fmt.Sprintf(":%d", *port))
}
