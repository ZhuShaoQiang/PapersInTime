package handler

import (
	"fmt"
	"main/types"
	"main/utils"

	"github.com/gofiber/fiber/v2"
)

func GetPaperDetail(c *fiber.Ctx) error {
	var resData map[string]interface{} = make(map[string]interface{})
	resData["status"] = false
	resData["msg"] = ""

	// 拿到前端传来的内容
	var resFromWeb types.PaperDetailFunc
	err := c.BodyParser(&resFromWeb)
	if err != nil {
		resData["msg"] = "获取论文id时发生错误：" + err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(resData)
	}
	fmt.Println("获取了:", resFromWeb.CiteName)
	flag, paper := utils.IsRecordExists(resFromWeb.CiteName)
	if !flag {
		// 如果存在，才继续走
		resData["msg"] = "论文ID:" + resFromWeb.CiteName + " 不存在"
		return c.Status(fiber.StatusNotAcceptable).JSON(resData)
	}

	resData["status"] = true
	resData["data"] = paper
	return c.JSON(resData)
}

func GetPaperByID(c *fiber.Ctx) error {
	var resData map[string]interface{} = make(map[string]interface{})
	resData["status"] = false
	resData["msg"] = ""

	// 拿到前端传来的内容
	var resFromWeb types.PaperDetailFunc
	err := c.BodyParser(&resFromWeb)
	if err != nil {
		resData["msg"] = "获取论文id时发生错误：" + err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(resData)
	}
	fmt.Println("获取了:", resFromWeb.CiteName)
	flag, paper := utils.IsRecordExists(resFromWeb.CiteName)
	if !flag {
		// 如果存在，才继续走
		resData["msg"] = "论文ID:" + resFromWeb.CiteName + " 不存在"
		return c.Status(fiber.StatusNotAcceptable).JSON(resData)
	}

	resData["status"] = true
	resData["data"] = paper
	return c.JSON(resData)
}
