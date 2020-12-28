package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"project/controllers"
	"project/controllers/admin"
)

func init() {
    beego.Include(&admin.LoginController{})
    beego.Include(&admin.AdminController{})
    beego.Include(&admin.IndexController{})
    beego.Include(&controllers.IndexController{})
    beego.Include(&controllers.CaptchaController{})
}
