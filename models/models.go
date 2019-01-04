package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	// register model
	orm.RegisterModel(new(Users))

}