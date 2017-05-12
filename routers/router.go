package routers

import (
	"eblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/edit", &controllers.MainController{},"*:Edit")
	beego.Router("/sign",&controllers.UserController{},"*:Sign")
	beego.Router("/signOut",&controllers.UserController{},"*:SignOut")
}