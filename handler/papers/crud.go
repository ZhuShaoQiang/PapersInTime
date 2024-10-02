package papers

import (
	"fmt"
	"main/dbs"
	"main/types"
	"main/utils"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Get all papers
func GetPapers(c *fiber.Ctx) error {
	var papers []types.Paper
	dbs.DB.Find(&papers)
	return c.JSON(papers)
}

// 上传论文或者更新论文
func UploadFile(c *fiber.Ctx) error {
	fmt.Println("uploadfile")
	var resData map[string]interface{} = make(map[string]interface{})
	resData["status"] = false
	resData["msg"] = "未知错误"

	// Get file from form
	file, err := c.FormFile("file")
	if err != nil {
		resData["msg"] = "取出文件失败"
		return c.Status(fiber.StatusBadRequest).JSON(resData)
	}

	// Create directory if not exists
	uploadDir := "./static/files"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		resData["msg"] = "创建文件的目录失败"
		return c.Status(fiber.StatusInternalServerError).JSON(resData)
	}

	// Save file to uploads directory
	filePath := filepath.Join(uploadDir, filepath.Base(file.Filename))
	if err := c.SaveFile(file, filePath); err != nil {
		resData["msg"] = "保存文件失败"
		return c.Status(fiber.StatusInternalServerError).JSON(resData)
	}

	// Get other form values
	citeName := c.FormValue("citename")
	title := c.FormValue("papertitle")
	author := c.FormValue("authorname")
	paperlabel := c.FormValue("paperlabel")
	year := c.FormValue("pubyear")
	cites := c.FormValue("cites")
	cited := c.FormValue("cited")
	// fmt.Println("拿到的数据:")
	// fmt.Println(citeName)
	// fmt.Println(title)
	// fmt.Println(author)
	// fmt.Println(paperlabel)
	// fmt.Println(year)
	// fmt.Println(cites)
	// fmt.Println(cited)
	// fmt.Println(filePath)

	flag, record := utils.IsRecordExists(citeName)

	var opRes *gorm.DB
	if flag {
		// 如果记录存在，更新记录
		record.Title = title
		record.ID = citeName
		record.Author = author
		record.Labels = paperlabel
		record.Year = year
		record.Cite = cites
		record.Cited = cited
		record.Path = filePath
		opRes = dbs.DB.Save(&record)
	} else {
		// Save paper info to database
		paper := types.Paper{
			ID:     citeName,
			Title:  title,
			Author: author,
			Labels: paperlabel,
			Year:   year,
			Path:   filePath,
			Cite:   cites,
			Cited:  cited,
		}
		opRes = dbs.DB.Create(&paper)
	}
	if opRes.Error != nil {
		resData["msg"] = "保存数据时出错：" + opRes.Error.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(resData)
	}

	resData["status"] = true
	return c.JSON(resData)
}

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
