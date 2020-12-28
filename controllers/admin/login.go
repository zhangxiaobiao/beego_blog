package admin

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"project/models"
	"project/models/admin"
	"strings"
)

type LoginController struct {
	beego.Controller
}

//@router /login [get]
func (this *LoginController) Login(){
	this.TplName = "admin/login.html"
}

//@router /dologin [post]
func (this *LoginController) GoLogin()  {
	captcha := this.GetString("code")
	captchaId := this.GetString("cid")
	email := this.GetString("email")
	password := this.GetString("password")
	logs.Info(captcha,captchaId)
	var result models.Response
	verify := this.Verify(captcha, captchaId)
	logs.Info(verify)
	if verify == false {
		result = R.Error("captcha error",nil)
	} else {
		var model admin.User
		user,err := model.Find(email,password)
		if err != nil {
			result = R.Error("login error",nil)
		} else {
			logs.Info(user)
			this.SetSession("user",user)
			url := "/admin/index.html"
			result = R.Success("login success",map[string]string{"url":url})
		}
	}
	this.Data["json"] = result
	this.ServeJSON()
}

//@router /logout [get]
func (this *LoginController) Logout()  {
	this.DestroySession();
	this.Data["json"] = R.Success("logout success", nil)
	this.ServeJSON()
}

func (this *LoginController) Verify(captcha string,id string) bool  {
	result := this.GetSession(id)
	if result == nil {
		return false
	}
	if strings.ToLower(fmt.Sprintf("%v", result)) != strings.ToLower(captcha) {
		return false
	}
	this.DelSession(id)
	return true
}
