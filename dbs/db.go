package dbs

import (
	"fmt"
	"main/types"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("papers.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("打开数据库失败!")
		os.Exit(1)
	}
	DB.AutoMigrate(&types.Paper{})
}
