package controllers

import (
	"eblog/model"
	"eblog/common/util"
	"os"
	"encoding/json"
	"io/ioutil"
	"github.com/astaxie/beego"
)

type BlogController struct {
	BaseController
}

//获取博客菜单
func (c *BlogController) GetMenu(){
	uname := c.GetSession("userName").(string)
	path:=beego.AppConfig.String("eblog_pwd")+"static_blog/"+uname
	if util.CheckFileIsExist(path) {//之前有博客，读取博客信息
		beego.Debug("file is exit")
		file,err:=os.Open(path+"/base.json")
		defer file.Close()
		if util.CheckErrorIsNotExist(err) {
			jsonString,err := ioutil.ReadAll(file)
			menu:=model.Menu{}
			json.Unmarshal(jsonString,&menu)
			if util.CheckErrorIsNotExist(err) {
				c.Data["json"]= model.Res{1,"",menu}
				c.ServeJSON()
				return
			}
		}

		c.Data["json"]= model.Res{0,err.Error(),""}
		c.ServeJSON()
	}else{//之前没有博客，创建博客
		beego.Debug("file is not exit")
		err:=os.MkdirAll(path,os.ModeDir)
		if util.CheckErrorIsNotExist(err) {
			file,err:=os.Create(path+"/base.json")
			defer file.Close()
			if util.CheckErrorIsNotExist(err) {
				jsonString,err:=json.Marshal(model.Menu{nil,0})
				if util.CheckErrorIsNotExist(err) {
					_, err := file.Write(jsonString)  //写入文件(字节数组)
					file.Sync()
					if util.CheckErrorIsNotExist(err) {
						c.Data["json"]= model.Res{1,"",model.Menu{nil,0}}
						c.ServeJSON()
						return
					}
				}
			}
		}
		c.Data["json"]= model.Res{0,err.Error(),""}
		c.ServeJSON()
	}
}

//菜单删除
func (c *BlogController) DeleteMenu(){

}

//文件归类
func (c *BlogController) MoveMenu(){

}
