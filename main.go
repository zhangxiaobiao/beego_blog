package main

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "project/models"
	_ "project/routers"
)

func main() {
	beego.BConfig.WebConfig.DirectoryIndex=true
	orm.Debug = true
	beego.Run()
}

