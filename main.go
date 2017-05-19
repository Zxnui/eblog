package main

import (
	_"eblog/common/redis"
	_ "eblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Debug("eblog is begin")
	beego.Run()
}

