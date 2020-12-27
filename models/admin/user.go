package admin

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

type User struct {
	Id int
	Name string
	Email string
	Password string
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func (this *User) TableName() string {
	return "user"
}


func (this *User) CreateUser()  {
	result,err := ORM.Insert(this)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v",result);
}

func (this *User) Find(email,password string) (User,error)  {
	var user User
	qs := ORM.QueryTable(user)
	err := qs.Filter("email",email).Filter("password",password).One(&user)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		logs.Error("Returned Multi Rows Not One")
		return user,err
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		logs.Error("Not row found")
		return user,err
	}
	return user,nil
}


