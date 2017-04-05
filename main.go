package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "unicontract/src/api/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,
	"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.Warn("main start")

	beego.Run()
}
