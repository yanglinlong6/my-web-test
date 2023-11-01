package models

type Order struct {
	//gorm.Model
	Id           int    `orm:"pk;auto" json:"id"`
	OrderTitle   string `json:"order_title"`
	OrderContent string `json:"order_content"`
	Status       int    `json:"status"`
	Remark       string `orm:"column(remark);null" json:"remark"`
}

func (Order) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "t_order"
}
