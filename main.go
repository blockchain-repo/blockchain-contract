package main

import (
	//"flag"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "unicontract/src/api/routers"
)

func main() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,
	"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,
	"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	logs.Warn("main start")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()

	//var command *string = flag.String("c", "command", "Start beego")
	//flag.Parse()
	//fmt.Println("param command is \t", *command)
	//if command != nil {
	//
	//	logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,
	//"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	//	logs.Warn("main start")
	//	if *command == "api" {
	//		if beego.BConfig.RunMode == "dev" {
	//			beego.BConfig.WebConfig.DirectoryIndex = true
	//			beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//		}
	//
	//		beego.Run()
	//	} else if *command == "test" {
	//		fmt.Println("test...")
	//	}

	//}

}
