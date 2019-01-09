package models

type Page struct {
	Tag	, LinkPage		string
}

var (
	PageChilds []Page = []Page {
		Page{
			Tag:"Page1",
			LinkPage:"testPage1.tpl",
		},

		Page{
			Tag:"Page2",
			LinkPage:"testPage2.tpl",
		},

		Page{
			Tag:"Page3",
			LinkPage:"welcome.tpl",
		},

		Page{
			Tag:"扛打排程",
			LinkPage:"ip_monitoring.tpl",
		},

		Page{
			Tag:"Page5",
			LinkPage:"testPage1.tpl",
		},
	}
)

func GeneratePages() []Page {
	return PageChilds;
}