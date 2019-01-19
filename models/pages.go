package models

type Page struct {
	Tag	, LinkPage , IconPath		string
}

var (
	PageChilds []Page = []Page {
		Page{
			Tag:"進度狀況",
			LinkPage:"welcome.tpl",
			IconPath:"/static/img/report.png",
		},
		Page{
			Tag:"扛打監聽",
			LinkPage:"test_ping.tpl",
			IconPath:"/static/img/monitor.png",
		},
		Page{
			Tag:"排程設定",
			LinkPage:"ip_monitoring.tpl",
			IconPath:"/static/img/schedule.png",
		},
		Page{
			Tag:"夜間監控事項紀錄",
			LinkPage:"report_monitoring.tpl",
			IconPath:"/static/img/report.png",
		},
	}
)

func GeneratePages() []Page {
	return PageChilds;
}