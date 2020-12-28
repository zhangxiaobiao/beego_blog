package admin

import (
	"project/models/admin"
)

type IndexController struct {
	BaseController
}

// 初始化后台框架接口
// @router /init [get]
func (c *IndexController) SystemInit() {
	systemInit := new(admin.SystemMenu).GetSystemInit()
	c.Data["json"] = systemInit
	c.ServeJSON()
}

