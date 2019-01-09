package main

import (
	_ "CMS/models"
	_ "CMS/routers"
	"github.com/astaxie/beego"
	"time"
)
var TimeInfo time.Time

func main() {
/*	t := time.NewTicker(3*time.Second)
	defer t.Stop()

	fmt.Println(time.Now())
	time.Sleep(4*time.Second)
	for  {
		select {
		case <-t.C:
			fmt.Println(time.Now())
			TimeInfo = time.Now()
		default:
		}
	}*/
	beego.Run()
}

