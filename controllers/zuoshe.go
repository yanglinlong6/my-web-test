package controllers

import (
	"my-web-test/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// ZuosheController operations for Zuoshe
type ZuosheController struct {
	beego.Controller
}

// URLMapping ...
func (c *ZuosheController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Zuoshe
// @Param	body		body 	models.Zuoshe	true		"body for Zuoshe content"
// @Success 201 {object} models.Zuoshe
// @Failure 403 body is empty
// @router / [post]
func (c *ZuosheController) Post() {
	o := orm.NewOrm()
	devices := new(models.TNesDevice)
	qs := o.QueryTable(devices)
	filter := qs.Filter("sn", "1110424280")
	var deviceList []models.TNesDevice
	_ = filter.One(&deviceList)
	c.Data["json"] = Respose{Code: 200, Msg: "获取成功Post", Data: &deviceList}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Zuoshe by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Zuoshe
// @Failure 403 :id is empty
// @router /getone [get]
func (c *ZuosheController) GetOne() {
	o := orm.NewOrm()
	devices := new(models.TNesDevice)
	qs := o.QueryTable(devices)
	filter := qs.Filter("sn", "1110424280")
	var deviceList []models.TNesDevice
	_ = filter.One(&deviceList)
	c.Data["json"] = Respose{Code: 200, Msg: "获取成功GetOne", Data: &deviceList}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Zuoshe
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Zuoshe
// @Failure 403
// @router / [get]
func (c *ZuosheController) GetAll() {
	o := orm.NewOrm()
	devices := new(models.TNesDevice)
	qs := o.QueryTable(devices)
	filter := qs.Filter("sn", "1110424280")
	var deviceList []models.TNesDevice
	_ = filter.One(&deviceList)
	c.Data["json"] = Respose{Code: 200, Msg: "获取成功GetAll", Data: &deviceList}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Zuoshe
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Zuoshe	true		"body for Zuoshe content"
// @Success 200 {object} models.Zuoshe
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ZuosheController) Put() {
	o := orm.NewOrm()
	devices := new(models.TNesDevice)
	qs := o.QueryTable(devices)
	filter := qs.Filter("sn", "1110424280")
	var deviceList []models.TNesDevice
	_ = filter.One(&deviceList)
	c.Data["json"] = Respose{Code: 200, Msg: "获取成功Put", Data: &deviceList}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Zuoshe
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ZuosheController) Delete() {
	o := orm.NewOrm()
	devices := new(models.TNesDevice)
	qs := o.QueryTable(devices)
	filter := qs.Filter("sn", "1110424280")
	var deviceList []models.TNesDevice
	_ = filter.One(&deviceList)
	c.Data["json"] = Respose{Code: 200, Msg: "获取成功Delete", Data: &deviceList}
	c.ServeJSON()
}
