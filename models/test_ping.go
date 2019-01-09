package models

import (
	"github.com/astaxie/beego/orm"
)

type TestPingCategory struct {
	Id			int		`orm:"column(id)"`
	Category	string	`orm:"column(category)"`
	Status		int		`orm:"column(status)"`
}

func (u *TestPingCategory) TableName() string {
	return "testpingcategory"
}

type TestPingItem struct {
	Id			int		`orm:"column(id)"`
	CategoryId	int		`orm:"column(category_id)"`
	Item		string	`orm:"column(item)"`
	Status		int		`orm:"column(status)"`
}

func (u *TestPingItem) TableName() string {
	return "testpingitem"
}



type TestPingData struct {
	CategoryId	int		`orm:"column(category_id)"`
	ItemId		int		`orm:"column(item_id)"`
	Category	string	`orm:"column(category)"`
	Item		string	`orm:"column(item)"`
}

type TestPingCategoryCount struct {
	Id			int		`orm:"column(id)"`
	Count		int		`orm:"column(count)"`
}

type TestPingResultData struct {
	ItemId		int		`orm:"column(item_id)"`
	ResultId	int		`orm:"column(result_id)"`
	Time		string	`orm:"column(time)"`
	Status		int		`orm:"column(status)"`
}

func GetAllTestPingData(o orm.Ormer, data *[]TestPingData) (int64, error) {
	return o.Raw("SELECT ti.category_id, ti.id AS item_id, tc.title AS category, ti.title AS item FROM testpingcategory AS tc, testpingitem AS ti WHERE tc.id = ti.category_id ORDER BY ti.id;").QueryRows(data)
}

func GetCategoryCount(o orm.Ormer, data *[]TestPingCategoryCount) (int64, error) {
	return o.Raw("SELECT category_id AS id, count(category_id) AS count  FROM testpingitem GROUP BY category_id;").QueryRows(data)
}

func GetTestPingResultByDate(o orm.Ormer, data *[]TestPingResultData, date string) (int64, error) {
	return o.Raw("SELECT ti.id AS item_id, tr.id AS result_id, tr.time, tr.status FROM testpingcategory AS tc, testpingitem AS ti, testpingresult AS tr WHERE tc.id = ti.category_id AND ti.id = tr.item_id AND tr.date=?;", date).QueryRows(data)
}