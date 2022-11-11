package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"my-web-test/models"
)

// YangController operations for Yang
type YangController struct {
	beego.Controller
}

type Respose struct {
	Code int64
	Msg  string
	Data interface{}
}

func (c *YangController) GetOne() {
	o := orm.NewOrm()
	devices := new(models.TNesDevice)
	qs := o.QueryTable(devices)
	filter := qs.Filter("sn", "1110424280")
	var deviceList []models.TNesDevice
	_ = filter.One(&deviceList)
	c.Data["json"] = Respose{Code: 200, Msg: "获取成功", Data: &deviceList}
	c.ServeJSON()
}

func (c *YangController) GetAll() {
	o := orm.NewOrm()
	devices := new(models.TNesDevice)
	qs := o.QueryTable(devices)
	filter := qs.Filter("sn", "1110424280")
	var deviceList []models.TNesDevice
	_ = filter.One(&deviceList)
	c.Data["json"] = Respose{Code: 200, Msg: "获取成功", Data: &deviceList}
	c.ServeJSON()
}

func (c *YangController) Get() {
	o := orm.NewOrm()
	devices := new(models.TNesDevice)
	qs := o.QueryTable(devices)
	filter := qs.Filter("sn", "1110424280")
	var deviceList []models.TNesDevice
	_ = filter.One(&deviceList)
	c.Data["json"] = Respose{Code: 200, Msg: "获取成功", Data: &deviceList}
	c.ServeJSON()
}

func (c *YangController) Post() {
	o := orm.NewOrm()
	devices := new(models.TNesDevice)
	qs := o.QueryTable(devices)
	filter := qs.Filter("sn", "1110424280")
	var deviceList []models.TNesDevice
	_ = filter.One(&deviceList)
	c.Data["json"] = Respose{Code: 200, Msg: "获取成功", Data: &deviceList}
	c.ServeJSON()
}
