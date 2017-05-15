package controllers

import (
	"github.com/astaxie/beego"
	"eblog/model"
	"fmt"
)

type UserController struct {
	BaseController
}

func (c *UserController) Sign(){
	if c.Ctx.Input.IsGet() {
		c.TplName = "sign.html"
	}else if c.Ctx.Input.IsPost() {
		v:=c.GetSession("userName")
		if v!=nil {//登录过，则直接跳转
			c.Data["json"]= model.Res{1,"",""}
			c.ServeJSON()
		}
		u := model.User{}
		if err := c.ParseForm(&u); err != nil {//获取请求参数，用户名和密码
			c.Data["json"]= model.Res{0,err.Error(),""}
			c.ServeJSON()
		}else{
			//获取配置文件中的用户名和密码
			pwd,_:=beego.GetConfig("String","user::"+u.Name,"")
			if pwd!=nil&&pwd==u.Password {
				c.SetSession("userName",u.Name)
				c.Data["json"]= model.Res{1,"",""}
				c.ServeJSON()
			}else{
				fmt.Println(pwd)
				fmt.Println(u.Password)
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
	c.Data["UserName"]=nil
	c.TplName = "index.html"
}
