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

	beego.Router("/eblog/getMenu",&controllers.BlogController{},"*:GetMenu")
	beego.Router("/eblog/delete",&controllers.BlogController{},"*:DeleteMenu")
	beego.Router("/eblog/move",&controllers.BlogController{},"*:MoveMenu")
	beego.Router("/eblog/save",&controllers.BlogController{},"*:SaveBlog")
}