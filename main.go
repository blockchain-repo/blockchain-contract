package main

import (
	//"flag"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "unicontract/src/api/routers"
)

func main() {
	beego.LoadAppConfig("ini", "../conf/app.conf")

	//日志默认不输出调用的文件名和文件行号,如果你期望输出调用的文件名和文件行号,可以如下设置
	//logs.SetLogFuncCall(true)
	//如果你的应用自己封装了调用 log 包,那么需要设置 SetLogFuncCallDepth,默认是 2,
	// 也就是直接调用的层级,如果你封装了多层,那么需要根据自己的需求进行调整.
	//logs.EnableFuncCallDepth(true)

	//如果不想在控制台输出log相关的，可以打开下面设置
	//todo if u want not output to console, open following line!
	//beego.BeeLogger.DelLogger("console")

	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"unicontract.log","level":7,
	"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,
	"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	logs.Warn("main start")

	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}

	beego.Run()
}
