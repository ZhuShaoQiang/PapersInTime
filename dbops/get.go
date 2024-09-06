package dbops

import (
	"main/dbs"
	"main/types"

	"github.com/gofiber/fiber/v2"
)

// Get all papers
func GetPapers(c *fiber.Ctx) error {
	var papers []types.Paper
	dbs.DB.Find(&papers)
	return c.JSON(papers)
}
