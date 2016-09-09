package main

import (
	_ "gluster-api/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Info(beego.BConfig.AppName, "1.0")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
