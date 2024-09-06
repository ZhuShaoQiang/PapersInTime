package utils

import (
	"fmt"
	"main/dbs"
	"main/types"
	"os"

	"gorm.io/gorm"
)

// 判断主键是否存在
// 如果存在，返回true和存在的内容
// 如果不存在，返回false
func IsRecordExists(primaryKey string) (flag bool, tmp types.Paper) {
	flag = false
	tmp = types.Paper{}

	res := dbs.DB.Where("id=?", primaryKey).First(&tmp)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			flag = false
			return
		}
		fmt.Println("判断记录是否存在时出现未知错误: ", res.Error.Error())
		os.Exit(1)
	}
	flag = true
	return
}
