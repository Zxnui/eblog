package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *MainController) Edit(){
	c.TplName = "edit.html"
}
