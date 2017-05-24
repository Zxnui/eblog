package main

import (
	_ "eblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Debug("eblog is begin")
	beego.Run()
}

