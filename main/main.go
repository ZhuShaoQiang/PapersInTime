package main

import (
	"log"
	"main/dbops"
	"main/dbs"
	"main/handler"
	"main/handler/api"
	"main/handler/view"
	"main/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize database
	// sqlite只支持：null, int, book, date and time, real, string, blob
	dbs.InitDatabase()

	app.Use(middleware.RootToManagement)

	// Serve static files (HTML, JS, etc.)
	app.Static("/static", "./static", fiber.Static{
		Browse: false,
		Index:  "papers/management.html",
	})

	// API routes
	app.Post("/upload", dbops.UploadFile)
	app.Get("/getpapers", dbops.GetPapers)

	// 获取论文详情
	app.Post("/papers/detail", handler.GetPaperDetail)
	app.Post("/papers/getPaperByID", handler.GetPaperByID)

	// 查看页面
	app.Get("/view/intree", view.ViewInTree)

	// 根据basepaper，获取引用它的，和它引用的
	app.Post("/papers/getCitesAndCited", api.GetPaperCitesAndCited)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
