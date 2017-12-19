package controllers

import (
	"github.com/astaxie/beego"
	"github.com/zxnui/eblog/model"
	"github.com/garyburd/redigo/redis"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}

func (c *MainController) Edit() {
	c.TplName = "edit.html"
}

func (c *MainController) GetNetIp(){
	// 从池里获取连接
	rc := model.RedisClient.Get()
	defer rc.Close()
	ip,_:=redis.String(rc.Do("Get","ipAddr"))
	json := model.BaseRes{Code:1,Data:ip}
	c.Data["json"] = &json
	c.ServeJSON()
}

func (c *MainController) File(){
	ip := c.GetString("ipAddr")
	// 从池里获取连接
	rc := model.RedisClient.Get()
	defer rc.Close()
	rc.Send("Set","ipAddr",ip)
	json := model.BaseRes{Code:1}
	c.Data["json"] = &json
	c.ServeJSON()
}