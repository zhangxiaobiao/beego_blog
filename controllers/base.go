package controllers

import beego "github.com/beego/beego/v2/server/web"

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare()  {
	this.Data["islogin"] = 0
	if user := this.GetSession("user");user != nil{
		this.Data["islogin"] = 1
		this.Data["user"] = user
	}
}
