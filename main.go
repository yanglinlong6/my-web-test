package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"my-web-test/models"
	_ "my-web-test/routers"
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		"dev_user",
		"df234fl",
		"192.168.3.190",
		3306,
		"neshield",
		"utf8",
		"utf8mb4_general_ci",
		"parseTime=True&loc=Local&timeout=10000ms",
	)

	orm.RegisterDataBase("default", "mysql", dsn, 30)
	orm.RegisterModel(new(models.TNesDevice))
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
