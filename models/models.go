package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	// register model
	orm.RegisterModel(new(Users))
}

func InsertOperationLog(o orm.Ormer, action, operator, before, after interface{}) error {
	_, err := o.Raw("INSERT INTO `CMS`.`operation_log` (`action`, `operator`, `before`, `after`) VALUES (?, ?, ?, ?);", action, operator, before, after).Exec()
	return err
}