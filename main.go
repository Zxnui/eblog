package main

import (
	_ "eblog/routers"
	"github.com/astaxie/beego"
)

func main() {

	////若是登录过，则前端显示用户名
	//var finish = func(ctx *context.Context) {
	//	_, ok := ctx.
	//	if !ok && ctx.Request.RequestURI != "/login" {
	//		ctx.Redirect(302, "/login")
	//	}
	//}
	//
	//beego.InsertFilter("*",beego.FinishRouter,finish)
	beego.Run()
}

