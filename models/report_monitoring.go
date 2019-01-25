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

type MonitoringCategoryCount struct {
	Id			int		`orm:"column(id)"`
	Count		int		`orm:"column(count)"`
}

type MonitoringResultData struct {
	Id			int		`json:"ResultId" orm:"column(result_id)"`
	ItemId		int		`orm:"column(item_id)"`
	Status		int		`orm:"column(status)"`
	Note		string	`orm:"column(note)"`
}

func GetBaseAllReportMonitoringSchema(o orm.Ormer, data *[]MonitoringSchema) (int64, error) {
	return o.Raw("SELECT ri.category_id, ri.id AS item_id, rc.title AS category, ri.title AS item FROM reportcategory AS rc, reportitem AS ri WHERE rc.id = ri.category_id ORDER BY rc.id, ri.id;").QueryRows(data)
}

func GetBaseReportMonitoringCategoryCount(o orm.Ormer, data *[]MonitoringCategoryCount) (int64, error) {
	return o.Raw("SELECT category_id AS id, count(category_id) AS count  FROM reportitem GROUP BY category_id;").QueryRows(data)
}

func GetReportMonitoringResultByDate(o orm.Ormer, data *[]MonitoringResultData, date string) (int64, error) {
	return o.Raw("SELECT item_id, id AS result_id, status, note FROM CMS.reportresult WHERE date = ?;", date).QueryRows(data)
}

func GetPreviousReportMonitoringSchema(o orm.Ormer, data *[]MonitoringSchema, date string) (int64, error) {
	return o.Raw("SELECT category_id, item_id, category, item FROM CMS.reportresult WHERE date = ?;", date).QueryRows(data)
}

func GetPreviousReportMonitoringCategoryCount(o orm.Ormer, data *[]MonitoringCategoryCount, date string) (int64, error) {
	return o.Raw("SELECT category_id AS id, count(category_id) FROM CMS.reportresult WHERE date = ? GROUP BY category_id;", date).QueryRows(data)
}

func GetReportNonitoringResultById(o orm.Ormer, data *MonitoringResultData, id int) error {
	return o.Raw("SELECT * FROM CMS.reportresult WHERE id = ?;", id).QueryRow(data)
}

func EditReportMonitoringStatus(o orm.Ormer, id, status int) error {
	_, err := o.Raw("UPDATE `CMS`.`reportresult` SET `status`= ? WHERE `id`=?;", status, id).Exec()
	return err
}

func EditReportMonitoringNote(o orm.Ormer, id int, note string) error {
	_, err := o.Raw("UPDATE `CMS`.`reportresult` SET `note`= ? WHERE `id`=?;", note, id).Exec()
	return err
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