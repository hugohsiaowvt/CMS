package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	// register model
	orm.RegisterModel(new(Users))
}

//===LOG===

type EditIPMonitoringAfter struct {
	ID int64
	Category int
	Title string
	IP string
	Type int
}

type AddIPMonitoringResultAfter struct {
	CategoryId int
	Status int
	ItemId int64
	Date string
	Time string
	Category string
	Item string
	IP string
}

type EditStatusBefore struct {
	Id int
	Status int
}

type EditStatusAfter struct {
	Id int
	Status int
}

type EditNoteBefore struct {
	Id int
	Note string
}

type EditNoteAfter struct {
	Id int
	Note string
}

//===LOG===

func InsertOperationLog(o orm.Ormer, action, operator, before, after interface{}) error {
	_, err := o.Raw("INSERT INTO `CMS`.`operation_log` (`action`, `operator`, `before`, `after`) VALUES (?, ?, ?, ?);", action, operator, before, after).Exec()
	return err
}