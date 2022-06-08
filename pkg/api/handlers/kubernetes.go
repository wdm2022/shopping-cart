package handlers

import "github.com/gofiber/fiber/v2"

// Liveness endpoint is used by Kubernetes to determine if the application
// after startup is still able to serve the requests
func Liveness(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

// Readiness endpoint is used by Kubernetes to determine if the application
// finished startup and can serve requests
func Readiness(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
