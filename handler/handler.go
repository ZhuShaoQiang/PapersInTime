package handler

import (
	"main/handler/api"
	"main/handler/papers"
	"main/handler/view"

	"github.com/gofiber/fiber/v2"
)

func initPostRoute(app *fiber.App) {
	// 获取论文详情
	app.Post("/papers/detail", papers.GetPaperDetail)
	app.Post("/papers/getPaperByID", papers.GetPaperByID)
	// 上传论文
	app.Post("/upload", papers.UploadFile)
	// 根据basepaper，获取引用它的，和它引用的
	app.Post("/papers/getCitesAndCited", api.GetPaperCitesAndCited)
	// 删除论文
	app.Post("/papers/delete", api.DeletePaperByID)
}

func initGetRoute(app *fiber.App) {
	// 查看页面
	app.Get("/view/intree", view.ViewInTree)

	app.Get("/getpapers", papers.GetPapers)
}

// 初始化所有的路由，包括get路由和post路由
func InitAllRoute(app *fiber.App) {
	initPostRoute(app)
	initGetRoute(app)
}
