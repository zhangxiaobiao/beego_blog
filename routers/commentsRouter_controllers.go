package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["project/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["project/controllers/admin:AdminController"],
        beego.ControllerComments{
            Method: "Index",
            Router: "/admin/index",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["project/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["project/controllers/admin:AdminController"],
        beego.ControllerComments{
            Method: "PageIndex",
            Router: "/page/index",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["project/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["project/controllers/admin:AdminController"],
        beego.ControllerComments{
            Method: "GetUser",
            Router: "/user",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["project/controllers/admin:IndexController"] = append(beego.GlobalControllerRouter["project/controllers/admin:IndexController"],
        beego.ControllerComments{
            Method: "SystemInit",
            Router: "/init",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["project/controllers/admin:LoginController"] = append(beego.GlobalControllerRouter["project/controllers/admin:LoginController"],
        beego.ControllerComments{
            Method: "GoLogin",
            Router: "/dologin",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["project/controllers/admin:LoginController"] = append(beego.GlobalControllerRouter["project/controllers/admin:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["project/controllers/admin:LoginController"] = append(beego.GlobalControllerRouter["project/controllers/admin:LoginController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["project/controllers:CaptchaController"] = append(beego.GlobalControllerRouter["project/controllers:CaptchaController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/captcha",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["project/controllers:IndexController"] = append(beego.GlobalControllerRouter["project/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
