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
