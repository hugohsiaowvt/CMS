package models

import (
	"github.com/astaxie/beego/orm"
)

type TestPingCategory struct {
	Id			int		`orm:"column(id)"`
	Category	string	`orm:"column(category)"`
}

func (u *TestPingCategory) TableName() string {
	return "testpingcategory"
}

type TestPingItem struct {
	Id			int		`orm:"column(id)"`
	CategoryId	int		`orm:"column(category_id)"`
	Item		string	`orm:"column(item)"`
}

func (u *TestPingItem) TableName() string {
	return "testpingitem"
}



type TestPingData struct {
	CategoryId	int		`orm:"column(category_id)"`
	ItemId		int		`orm:"column(item_id)"`
	Category	string	`orm:"column(category)"`
	Item		string	`orm:"column(item)"`
	Ip			string	`orm:"column(ip)"`
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

func GetBaseAllTestPingData(o orm.Ormer, data *[]TestPingData) (int64, error) {
	return o.Raw("SELECT ti.category_id, ti.id AS item_id, tc.title AS category, ti.title AS item, ti.ip as ip FROM testpingcategory AS tc, testpingitem AS ti WHERE tc.id = ti.category_id ORDER BY ti.id;").QueryRows(data)
}

func GetBaseCategoryCount(o orm.Ormer, data *[]TestPingCategoryCount) (int64, error) {
	return o.Raw("SELECT category_id AS id, count(category_id) AS count  FROM testpingitem GROUP BY category_id;").QueryRows(data)
}

func GetTestPingResultByDate(o orm.Ormer, data *[]TestPingResultData, date1, date2 string) (int64, error) {
	return o.Raw("SELECT item_id, id AS result_id, time, status FROM testpingresult WHERE (date=? AND time >= '1830') OR (date=? AND time <= '0600') ORDER BY category_id, item_id;", date1, date2).QueryRows(data)
}

func GetPreviousPingData(o orm.Ormer, data *[]TestPingData, date1, date2 string) (int64, error) {
	return o.Raw("SELECT distinct item, category_id, item_id, category, ip FROM CMS.testpingresult WHERE (date = ? AND time >= '0600') OR (date = ? AND time < '0600') ORDER BY category_id, item_id;", date1, date2).QueryRows(data)
}

func GetPreviousCategoryCount(o orm.Ormer, data *[]TestPingCategoryCount, date1, date2 string) (int64, error) {
	return o.Raw("SELECT category_id AS id, count(distinct item) FROM CMS.testpingresult WHERE (date = ? AND time >= '0600') OR (date = ? AND time < '0600') GROUP BY category_id;", date1, date2).QueryRows(data)
}

func InsertTestPingResult(o orm.Ormer, data TestPingData, date, time string, status int) error {
	_, err := o.Raw("INSERT INTO `CMS`.`testpingresult` (`category_id`, `item_id`, `date`, `time`, `category`, `item`, `ip`, `status`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		data.CategoryId, data.ItemId, date, time, data.Category, data.Item, data.Ip, status).Exec()
	return err
}

func AddCategory(o orm.Ormer, title string) error {
	_, err := o.Raw("INSERT INTO `CMS`.`testpingcategory` (`title`) VALUES (?);", title).Exec()
	return err
}

func AddItem(o orm.Ormer, c_id int, title, ip string) error {
	_, err := o.Raw("INSERT INTO `CMS`.`testpingitem` (`category_id`, `title`, `ip`) VALUES (?, ?, ?);", c_id, title, ip).Exec()
	return err
}

func DelItem(o orm.Ormer, id int) error {
	_, err := o.Raw("DELETE FROM `CMS`.`testpingitem` WHERE `id` = ?;", id).Exec()
	return err
}

func AddPingResult(o orm.Ormer, categoryId, itemId, status int, date, time, category, item, ip string) error {
	_, err := o.Raw("INSERT INTO `CMS`.`testpingresult` (`category_id`, `item_id`, `date`, `time`, `category`, `item`, `ip`, `status`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		categoryId, itemId, date, time, category, item, ip, status).Exec()
	return err
}
