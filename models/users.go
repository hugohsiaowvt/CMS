package models

import (
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id			int		`orm:"column(id)"`
	Account		string	`orm:"column(account)"`
	Password	string	`orm:"column(password)"`
	Name		string	`orm:"column(name)"`
	Email		string	`orm:"column(email)"`
	LoginTimes	int		`orm:"column(login_times)"`
	Status		int		`orm:"column(status)"`
}

func (u *Users) TableName() string {
	return "users"
}

func GetUser(o orm.Ormer, user, query *Users) {
	o.Raw("SELECT * FROM users WHERE account = ? and password = ?", query.Account, query.Password).QueryRow(user)
}