package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RootToManagement(c *fiber.Ctx) error {
	if c.OriginalURL() == "/" {
		return c.Redirect("/static/")
	}
	return c.Next()
}
