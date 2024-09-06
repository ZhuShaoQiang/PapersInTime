package utils_test

import (
	"fmt"
	"main/dbs"
	"main/utils"
	"os"
	"testing"
)

func Test_IsRecordExist(t *testing.T) {
	os.Chdir("../")
	dbs.InitDatabase()
	fmt.Println(utils.IsRecordExists("test1"))
}
