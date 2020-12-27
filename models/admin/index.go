package admin

import (
	beego "github.com/beego/beego/v2/server/web"
)

type IndexController struct {
	beego.Controller
}

// 初始化后台框架接口
func (c *IndexController) SystemInit() {
	systemInit := new(SystemMenu).GetSystemInit()
	c.Data["json"] = systemInit
	c.ServeJSON()
}
