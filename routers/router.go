package routers

import (
	"github.com/zxnui/eblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/edit", &controllers.MainController{}, "*:Edit")
	beego.Router("/wow", &controllers.MainController{}, "*:Wow")
}
