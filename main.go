package main

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "project/models"
	_ "project/routers"
)

func main() {
	beego.BConfig.WebConfig.DirectoryIndex=true
	beego.SetStaticPath("/admin","static/admin")
	orm.Debug = true
	beego.Run()
}

