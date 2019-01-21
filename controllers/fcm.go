package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type FCMController struct {
	beego.Controller
}

type FCM struct {
	Id	int		`json:"id"`
	Fcm	string	`json:"fcm"`
}

func (this *FCMController) Get() {

	res := &ResponseStatus{}
	res.Status = -1

	o := orm.NewOrm()

	fcm := &[]FCM{}

	if _, err := o.Raw("SELECT * FROM CMS.fcm;").QueryRows(fcm); err != nil {
		res.Msg = err.Error()
	} else {
		res.Status = 1
		res.Ext = fcm
	}

	this.Data["json"] = res
	this.ServeJSON()

}

func (this *FCMController) Post() {

	fcm := this.GetString("fcm")

	res := &ResponseStatus{}
	res.Status = -1

	o := orm.NewOrm()

	if _, err := o.Raw("INSERT INTO `CMS`.`fcm` (`fcm`) VALUES (?);", fcm).Exec(); err != nil {
		res.Msg = err.Error()
	} else {
		res.Status = 1
	}

	this.Data["json"] = res
	this.ServeJSON()

}

