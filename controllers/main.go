package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *MainController) Edit() {
	c.TplName = "edit.html"
}

func (c *MainController) Wow() {
	c.TplName = "wow.html"
}
