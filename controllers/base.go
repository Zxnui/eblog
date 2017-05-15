package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//controller之前执行
func (c *BaseController) Prepare(){
	url := c.Ctx.Request.RequestURI
	name:= c.GetSession("userName")
	if url != "/" && url != "/signOut" && url != "/sign" {
		if name == nil {
			c.Ctx.Redirect(302, "/sign")
		}
	}

	c.Data["UserName"] = name
}

//method之后执行
func (c *BaseController) Finish(){

}