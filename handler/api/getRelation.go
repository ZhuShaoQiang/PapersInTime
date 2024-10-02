package api

import (
	"fmt"
	"main/dbs"
	"main/types"
	"main/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetPaperCitesAndCited(c *fiber.Ctx) error {
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
	fmt.Println("citeName", resFromWeb.CiteName)
	flag, tmp := utils.IsRecordExists(resFromWeb.CiteName)
	if !flag {
		// 如果存在，才继续走
		resData["msg"] = "论文ID:" + resFromWeb.CiteName + " 不存在"
		return c.Status(fiber.StatusNotAcceptable).JSON(resData)
	}
	// 此时拿到了这个数据
	var paper types.BasepaperCitesAndCited
	var papersCitesFind []types.Paper
	var papersCitedFind []types.Paper
	paper.Author = tmp.Author
	paper.ID = tmp.ID
	paper.Path = tmp.Path
	paper.Title = tmp.Title
	paper.Year = tmp.Year
	if tmp.Cite != "" {
		cites_tmp := strings.ReplaceAll(tmp.Cite, " ", "")  // 去除空格
		cites_tmp = strings.ReplaceAll(cites_tmp, ";", ",") // 可以使用;分开
		cites := strings.Split(cites_tmp, ",")              // 拿到引用的论文列表了
		dbs.DB.Where("id IN ?", cites).Find(&papersCitesFind)
		paper.Cite = papersCitesFind // 赋值
	} else {
		paper.Cite = []types.Paper{}
	}
	if tmp.Cited != "" {
		cited_tmp := strings.ReplaceAll(tmp.Cited, " ", "") // 去除空格
		cited_tmp = strings.ReplaceAll(cited_tmp, ";", ",") // 可以使用;分开
		cited := strings.Split(cited_tmp, ",")              // 拿到被引用的列表了
		fmt.Println("自己的cited:", cited)
		// 找到手动记录的被引用的列表对象
		dbs.DB.Where("id IN ?", cited).Find(&papersCitedFind)
	} else {
		paper.Cited = []types.Paper{}
	}
	// 上面只是处理了他自己中记录的自己被引用的情况，有可能别人引用他他不知道。

	var papersCitePaper []types.Paper // 查一下谁引用自己了，然后送过去
	dbs.DB.Where("cite LIKE ?", "%"+paper.ID+"%").Find(&papersCitePaper)
	paper.Cited = append(paper.Cited, papersCitePaper...)

	// fmt.Println("查询结果:", papersCitePaper)

	resData["status"] = true
	resData["data"] = paper
	return c.JSON(resData)
}
