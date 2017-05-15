package util

import "github.com/astaxie/beego"

func CheckErrorIsNotExist(err error) bool{
	if err!=nil {
		beego.Error("Error : "+err.Error())
		return false
	}else{
		return true
	}
}