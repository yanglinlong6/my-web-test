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
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
	// 	"dev_user",
	// 	"df234fl",
	// 	"192.168.3.190",
	// 	3306,
	// 	"neshield",
	// 	"utf8",
	// 	"utf8mb4_general_ci",
	// 	"parseTime=True&loc=Local&timeout=10000ms",
	// )

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
		"root",
		"986203",
		"112.74.125.238",
		3306,
		"easytest",
		"utf8",
		"utf8mb4_general_ci",
		"parseTime=True&loc=Local&timeout=10000ms",
	)

	orm.RegisterDataBase("default", "mysql", dsn, 30)
	orm.RegisterModel(new(models.TNesDevice))
	orm.RegisterModel(new(models.Order))
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		beego.Error(err)
	}

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//比较常用的是 init,tidy, edit
	//go mod tidy
	//go mod vendor
	//go mod init helloworld
	//bee api testapi
	//bee generate controller hello
	//bee run
	beego.Run()
}
