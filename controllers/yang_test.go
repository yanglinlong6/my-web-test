package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"my-web-test/models"
	"testing"
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

func TestYangController_GetOne(t *testing.T) {
	o := orm.NewOrm()
	devices := new(models.TNesDevice)
	qs := o.QueryTable(devices)
	filter := qs.Filter("sn", "1110424280")
	var deviceList []*models.TNesDevice
	_ = filter.One(&deviceList)
	for _, value := range deviceList {
		fmt.Println("value===", value)
	}
	fmt.Println("deviceList===", deviceList)

	var list []orm.Params
	_, _ = o.Raw("select * from t_nes_device").Values(&list)
	for _, value := range list {
		fmt.Println("valueall===", value)
		fmt.Println("valuekey===", value["sn"])
	}
}
