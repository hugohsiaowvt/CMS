package main

import (
	_ "CMS/routers"
	_ "CMS/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

