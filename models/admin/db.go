package admin

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

var ORM orm.Ormer

func init() {
	mysqluser, _ := beego.AppConfig.String("mysqluser")
	mysqlpass, _ := beego.AppConfig.String("mysqlpass")
	mysqlurls, _ := beego.AppConfig.String("mysqlurls")
	mysqlport, _ := beego.AppConfig.String("mysqlport")
	mysqldb, _ := beego.AppConfig.String("mysqldb")
	dataSource := mysqluser + ":" + mysqlpass + "@tcp(" + mysqlurls + ":" + mysqlport + ")/" + mysqldb + "?charset=utf8mb4&loc=Local"

	orm.RegisterDriver("mysql",orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s",dataSource))

	orm.RegisterModel(new(User))

	orm.RunSyncdb("default",false,true)

	ORM = orm.NewOrm()

}
