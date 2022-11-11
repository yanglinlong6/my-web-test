package models

type TNesDevice struct {
	//gorm.Model
	Sn        string `orm:"pk" json:"sn"`
	ModelName string `json:"modelName"`
}

func (TNesDevice) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "t_nes_device"
}
