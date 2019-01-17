package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type ReportCategory struct {
	Id			int64	`orm:"column(id)"`
	Title		string	`orm:"column(title)"`
}

func (u *ReportCategory) TableName() string {
	return "reportcategory"
}

type ReportItem struct {
	Id			int64	`orm:"column(id)"`
	CategoryId	int		`orm:"column(category_id)"`
	Title		string	`orm:"column(title)"`
}

func (u *ReportItem) TableName() string {
	return "reportitem"
}

type ReportResult struct {
	Id			int64		`orm:"column(result_id)"`
	CategoryId	int64	`orm:"column(category_id)"`
	ItemId		int64	`orm:"column(item_id)"`
	Category	string	`orm:"column(category)"`
	Item		string	`orm:"column(item)"`
	Date		string	`orm:"column(date)"`
	Note		string	`orm:"column(note)"`
	Remark		string	`orm:"column(remark)"`
}

func (u *ReportResult) TableName() string {
	return "reportresult"
}


type MonitoringSchema struct {
	CategoryId	int64	`orm:"column(category_id)"`
	ItemId		int64	`orm:"column(item_id)"`
	Category	string	`orm:"column(category)"`
	Item		string	`orm:"column(item)"`
}

func GetBaseAllMonitoringSchema(o orm.Ormer, data *[]MonitoringSchema) (int64, error) {
	return o.Raw("SELECT ri.category_id, ri.id AS item_id, rc.title AS category, ri.title AS item FROM reportcategory AS rc, reportitem AS ri WHERE rc.id = ri.category_id ORDER BY rc.id, ri.id;").QueryRows(data)
}

func InsertMonitoringResult(o orm.Ormer, date string, data []ReportResult) error {
	s := &ReportResult{}
	sql := "INSERT INTO " + s.TableName() + " (category_id, item_id, category, item, date) VALUES "
	for k, v := range data {
		if k < len(data) - 1 {
			sql += fmt.Sprintf("('%d', '%d', '%s', '%s', '%s'), ", v.CategoryId, v.ItemId, v.Category, v.Item, date)
		} else {
			sql += fmt.Sprintf("('%d', '%d', '%s', '%s', '%s');", v.CategoryId, v.ItemId, v.Category, v.Item, date)
		}

	}
	_, err := o.Raw(sql).Exec()
	return err

}