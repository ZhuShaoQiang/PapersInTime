package main

import (
	"log"
	"main/dbs"
	"main/handler"
	"main/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // 100MB的最大大小
	})

	// Initialize database
	// sqlite只支持：null, int, book, date and time, real, string, blob
	dbs.InitDatabase()

	app.Use(middleware.RootToManagement)

	// Serve static files (HTML, JS, etc.)
	app.Static("/static", "./static", fiber.Static{
		Browse: false,
		Index:  "papers/management.html",
	})

	handler.InitAllRoute(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
