package main

import (
	"os"
	"strconv"
	_ "unicontract/src/api/routers"
	"unicontract/src/common"
	"unicontract/src/common/basic"
	"unicontract/src/config"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/engine/scanengine"
	"unicontract/src/pipelines"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	//beego.BConfig.EnableGzip = true
	beego.LoadAppConfig("ini", "../conf/app.conf")
	// logInit must following the beego.LoadAppConfig
	logInit()

	argsCount := len(os.Args)
	if argsCount == 2 && os.Args[1] == "start" {
		runStart()
	} else if argsCount == 2 && os.Args[1] == "initdb" {
		runInitDB()
	} else if argsCount == 2 && os.Args[1] == "dropdb" {
		runDropDB()
	} else if argsCount == 4 && os.Args[1] == "reconfigdb" {
		shards, error := strconv.Atoi(os.Args[2])
		if error != nil {
			logs.Error("shards should be int")
		}
		replicas, error := strconv.Atoi(os.Args[3])
		if error != nil {
			logs.Error("replicas should be int")
		}
		runReconfigDB(shards, replicas)
	} else if argsCount == 2 && os.Args[1] == "config" {
		runConfig()
	} else if argsCount == 2 && os.Args[1] == "help" {
		runHelp()
	} else {
		logs.Error("cmd should be " +
			"unicontract start|initdb|dropdb|reconfigdb $shards $replicas|config|help")
		os.Exit(2)
	}
}

func runStart() {
	config.Init()
	logs.Info("config Init")
	pipelines.Init()
	logs.Info("pipelines Init")
	go scanengine.Start()
	logs.Info("scanengine Init")
	beego.Run()
	logs.Info("beego Run")
}

func runInitDB() {
	config.Init()
	logs.Info("Database Init")
	rethinkdb.InitDatabase()
}

func runDropDB() {
	config.Init()
	logs.Info("Database Dropped")
	rethinkdb.DropDatabase()
}
func runReconfigDB(shards int, replicas int) {
	config.Init()
	logs.Info("Database Reconfigured")
	rethinkdb.Reconfig(shards, replicas)
}

func runConfig() {
	config.Init()
	config.WriteConToFile()
}

func runHelp() {
	logs.Info("cmd should be " +
		"unicontract start|initdb|dropdb|reconfigdb $shards $replicas|config|help")
}

func logInit() {
	//日志默认不输出调用的文件名和文件行号,如果你期望输出调用的文件名和文件行号,可以如下设置
	logs.SetLogFuncCall(true)
	//如果你的应用自己封装了调用 log 包,那么需要设置 SetLogFuncCallDepth,默认是 2,
	// 也就是直接调用的层级,如果你封装了多层,那么需要根据自己的需求进行调整.
	// logs 里面修改的话,此处请勿重复设置!
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	beego.BConfig.Log.AccessLogs = true

	//如果不想在控制台输出log相关的，可以打开下面设置
	//todo if u want not output to console, open following line!
	//beego.BeeLogger.DelLogger("console")

	myBeegoLogAdapterMultiFile := &basic.MyBeegoLogAdapterMultiFile{}
	myBeegoLogAdapterMultiFile.FileName = "../log/unicontract.log"
	//var levelPrefix = [LevelDebug + 1]string{"[M] ", "[A] ", "[C] ", "[E] ", "[W] ", "[N] ", "[I] ", "[D] "}
	myBeegoLogAdapterMultiFile.Level = 7
	myBeegoLogAdapterMultiFile.MaxDays = 10
	myBeegoLogAdapterMultiFile.MaxLines = 0
	myBeegoLogAdapterMultiFile.MaxSize = 0
	myBeegoLogAdapterMultiFile.Rotate = true
	myBeegoLogAdapterMultiFile.Daily = true
	//myBeegoLogAdapterMultiFile.Separate = []string{"emergency","critical","error", "warning", "debug"}
	//myBeegoLogAdapterMultiFile.Separate = []string{"emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"}
	myBeegoLogAdapterMultiFile.Separate = []string{"emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"}

	log_config := basic.NewMyBeegoLogAdapterMultiFile(myBeegoLogAdapterMultiFile)
	log_config_str := common.Serialize(log_config)
	logs.Info("beego log config: ", log_config_str)

	// order 顺序必须按照
	// 1. logs.SetLevel(level)
	// 2. logs.SetLogger(logs.AdapterMultiFile, log_config_str)
	logs.SetLevel(logs.LevelDebug)
	//logs.SetLevel(logs.LevelInfo)
	beego.SetLogger(logs.AdapterMultiFile, log_config_str)

}
