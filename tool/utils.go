package tool

import (
	"os"
	"path/filepath"
	"path"
	"strings"
)

//首字母转大写
func StrFirstToUpper(str string) string {
    if len(str) < 1 {
        return ""
    }
    strArry := []rune(str)
    if strArry[0] >= 97 && strArry[0] <= 122  {
        strArry[0] -= 32
    }
    return string(strArry)
}

//获取文件名称 没有后缀
func GetFileName(src string) (string, string) {
	filenameWithSuffix := filepath.Base(src)
    fileSuffix := path.Ext(filenameWithSuffix)
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)

	return filenameWithSuffix,filenameOnly
}

//检查文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}
