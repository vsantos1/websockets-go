package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)


func Sameorigin(c *fiber.Ctx) error{
	
	if c.Get("host") == "localhost:3000" {
		c.Locals("Host", "Localhost:3000")
		return c.Next()
	}
	return c.Status(http.StatusForbidden).SendString("Request origin not allowed")
}

