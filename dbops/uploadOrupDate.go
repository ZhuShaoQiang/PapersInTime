package dbops

import (
	"main/dbs"
	"main/types"
	"main/utils"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// 上传论文或者更新论文
func UploadFile(c *fiber.Ctx) error {
	var resData map[string]interface{} = make(map[string]interface{})
	resData["status"] = false
	resData["msg"] = "未知错误"

	// Get file from form
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot get file from form"})
	}

	// Create directory if not exists
	uploadDir := "./static/files"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot create upload directory"})
	}

	// Save file to uploads directory
	filePath := filepath.Join(uploadDir, filepath.Base(file.Filename))
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot save file"})
	}

	// Get other form values
	citeName := c.FormValue("citename")
	title := c.FormValue("papertitle")
	author := c.FormValue("authorname")
	paperlabel := c.FormValue("paperlabel")
	year := c.FormValue("pubyear")
	cites := c.FormValue("cites")
	cited := c.FormValue("cited")

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
