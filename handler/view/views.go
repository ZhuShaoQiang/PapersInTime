package view

import "github.com/gofiber/fiber/v2"

// 这个地方是跳转到这个页面，我们应该返回html
func ViewInTree(c *fiber.Ctx) error {
	return c.Redirect("/static/view/intree.html")
}
