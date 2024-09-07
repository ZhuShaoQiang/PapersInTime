package api

import (
	"fmt"
	"main/dbs"
	"main/types"
	"main/utils"

	"github.com/gofiber/fiber/v2"
)

func DeletePaperByID(c *fiber.Ctx) error {
	var resData map[string]interface{} = make(map[string]interface{})
	resData["status"] = false
	resData["msg"] = ""

	// 拿到前端传来的内容
	var resFromWeb types.PaperIDOnlyType
	err := c.BodyParser(&resFromWeb)
	if err != nil {
		resData["msg"] = "获取论文id时发生错误：" + err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(resData)
	}
	fmt.Println("获取了:", resFromWeb.CiteName)
	flag, _ := utils.IsRecordExists(resFromWeb.CiteName)
	if !flag {
		// 如果存在，才继续走
		resData["msg"] = "论文ID:" + resFromWeb.CiteName + " 不存在"
		return c.Status(fiber.StatusNotFound).JSON(resData)
	}
	// 删除论文
	err = dbs.DB.Delete(&types.Paper{}, "id = ?", resFromWeb.CiteName).Error
	if err != nil {
		resData["msg"] = "删除数据失败:" + err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(resData)
	}

	resData["status"] = true
	return c.JSON(resData)
}
