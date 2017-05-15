package util

import (
	"os"
	"github.com/astaxie/beego"
)

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileIsExist(filename string) (bool) {
	var exist = true;
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false;
		beego.Error(err.Error())
	}
	return exist;
}
