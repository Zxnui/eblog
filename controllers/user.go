package controllers

import (
	"github.com/astaxie/beego"
	"eblog/model"
	"fmt"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Sign(){
	if c.Ctx.Input.IsGet() {
		c.TplName = "sign.html"
	}else if c.Ctx.Input.IsPost() {
		v:=c.GetSession("userName")
		if v!=nil {
			c.Data["json"]= model.Res{1,"",""}
			c.ServeJSON()
		}
		u := model.User{}
		if err := c.ParseForm(&u); err != nil {
			c.Data["json"]= model.Res{0,err.Error(),""}
			c.ServeJSON()
		}else{
			pwd,err:=beego.GetConfig("String","user::"+u.Name,"")
			if err!=nil&&pwd!=nil&&pwd==u.Password {
				c.SetSession("userName",u.Name)
				c.Data["json"]= model.Res{1,"",""}
				c.ServeJSON()
			}else{
				fmt.Println(pwd)
				c.Data["json"]= model.Res{0,"用户名或密码错误",""}
				c.ServeJSON()
			}
		}
	}else{
		c.TplName = "/error/404.html"
	}

}

func (c *UserController) SignOut(){
	c.DelSession("userName")
	c.TplName = "index.html"
}
