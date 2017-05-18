package controllers

import (
	"eblog/model"
	"eblog/common/util"
	"os"
	"encoding/json"
	"io/ioutil"
	"github.com/astaxie/beego"
	"strconv"
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
				menuDetail := []model.MenuDetail{model.MenuDetail{"默认",1,0}}
				menu:=model.Menu{menuDetail,0}
				jsonString,err:=json.Marshal(menu)
				if util.CheckErrorIsNotExist(err) {
					_, err := file.Write(jsonString)  //写入文件(字节数组)
					file.Sync()
					if util.CheckErrorIsNotExist(err) {
						c.Data["json"]= model.Res{1,"",menu}
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

//菜单增加
func (c *BlogController) addMenu(){

}

//文件归类
func (c *BlogController) MoveMenu(){

}

//保存博客
func (c *BlogController) SaveBlog(){
	parentId:=c.Input().Get("parentId")
	name:=c.Input().Get("name")
	text:=c.Input().Get("text")
	parentName:=c.Input().Get("parentName")

	path:=beego.AppConfig.String("eblog_pwd")+"static_blog/"+c.GetSession("userName").(string)

	file_base, _ := os.Open(path + "/base.json")
	defer file_base.Close()

	jsonString, _ := ioutil.ReadAll(file_base)
	menu := model.Menu{}
	json.Unmarshal(jsonString, &menu)

	pid,_:=strconv.Atoi(parentId)
	menu.MaxId = menu.MaxId+1
	item:=model.MenuDetail{name,menu.MaxId,pid}
	beego.Debug(len(menu.Menu_list))
	menu.Menu_list = append(menu.Menu_list,item)
	jsonString,_ = json.Marshal(menu)

	beego.Debug(string(jsonString))
	ioutil.WriteFile(path+"/base.json",jsonString,os.ModeType)

	path=path+"/md/"+parentName
	err := ioutil.WriteFile(path+"/"+name+".md",[]byte(text),os.ModeType)
	if util.CheckErrorIsNotExist(err) {
		c.Data["json"] = model.Res{1, "", menu}
		c.ServeJSON()
		return
	}

	c.Data["json"]= model.Res{0,err.Error(),""}
	c.ServeJSON()
}
