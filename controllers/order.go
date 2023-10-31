package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"my-web-test/models"
)

// OrderController operations for Order
type OrderController struct {
	beego.Controller
}

// URLMapping ...
func (c *OrderController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Order
// @Param	body		body 	models.Order	true		"body for Order content"
// @Success 201 {object} models.Order
// @Failure 403 body is empty
// @router / [post]
func (c *OrderController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Order by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Order
// @Failure 403 :id is empty
// @router /:id [get]
func (c *OrderController) GetOne() {
	o := orm.NewOrm()
	orders := new(models.Order)
	qs := o.QueryTable(orders)
	filter := qs.Filter("id", "1")
	var orderList []models.Order
	_ = filter.One(&orderList)
	c.Data["json"] = Respose{Code: 200, Msg: "获取成功", Data: &orderList}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Order
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Order
// @Failure 403
// @router / [get]
func (c *OrderController) GetAll() {
	o := orm.NewOrm()
	order := models.Order{}
	//order.Id = 1
	order.OrderTitle = "kaishi"
	order.OrderContent = "美好"
	order.Status = 1

	beego.Info(order)
	insert, err := o.Insert(&order)
	if err != nil {
		beego.Error(err)
		return
	}
	c.Data["json"] = Respose{Code: 200, Msg: "获取成功", Data: insert}
	beego.Info(c)
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Order
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Order	true		"body for Order content"
// @Success 200 {object} models.Order
// @Failure 403 :id is not int
// @router /:id [put]
func (c *OrderController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Order
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *OrderController) Delete() {

}
