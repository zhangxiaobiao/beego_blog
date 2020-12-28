package admin

import (
	"project/models"
)

type AdminController struct {
	BaseController
}


//func (this *AdminController) Index()  {
//	this.TplName = "admin/index.html"
//}


// @router /user [get]
func (this *AdminController) GetUser()  {
	var result models.Response
	user := this.GetSession("user")
	if user == nil {
		result = R.Error("您已退出登录",nil)
	} else {
		result = R.Success("您已登录", this.GetSession("user"))
	}
	this.Data["json"] = result
	this.ServeJSON()
}

//后台框架首页
//@router /admin/index [get]
func (this *AdminController) Index()  {
	this.TplName = "admin/index.html"
}

//默认显示
//@router /page/index
func (this *AdminController) PageIndex()  {
	this.TplName = "admin/page/welcome.html"
}
