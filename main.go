package main

import (
	_ "github.com/zxnui/eblog/routers"
	"github.com/astaxie/beego"
	_ "github.com/zxnui/eblog/model"
)

func main() {
	beego.Debug("eblog is begin")
	beego.Run()
}

