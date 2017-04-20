package monitor

import (
	"os"
	"sync"

	"unicontract/src/common/yaml"
	"unicontract/src/common"

	"github.com/astaxie/beego"
	"github.com/alexcesaro/statsd"
)

var (
	once sync.Once
	config map[interface{}]interface{}
	Monitor *statsd.Client
	monitor *statsd.Client
)

/**
 * function : init函数
 * param   :
 * return : statsd.Client客户端
 */
func init(){
	monitor := _GetMonitor()
	Monitor = monitor
}

/**
 * function : 单例模式
 * param   :
 * return : statsd.Client客户端
 */
func _GetMonitorClient() *statsd.Client{

	once.Do(func(){
		//获取monitorConfig信息
		requestPath := os.Getenv("CONFIGPATH")
		requestPath = requestPath + "/monitorConfig.yaml"
		config := make(map[interface{}]interface{})
		err := yaml.Read(requestPath,config)
		if err != nil{
			beego.Error(err.Error())
		}
		//获取配置信息中内容
		ip :=common.TypeToString(config["ip"])
		port :=common.TypeToString(config["port"])
		maxPacketSize := common.TypeToInt(config["maxPacketSize"])

		add := ip + ":" + port
		//获取系统主机名
		pre,err := os.Hostname()
		if err != nil{
			beego.Error(err.Error())
		}
		//获取monitorClient客户端
		//准备statsd option
		address := statsd.Address(add)
		prefix := statsd.Prefix(pre)
		packetSize := statsd.MaxPacketSize(maxPacketSize)
		//创建monitor.client对象
		monitor,err = statsd.New(address,prefix,packetSize)
		if err != nil{
			beego.Error(err.Error())
		}
	})
	return monitor

}

/**
 * function: 获取monitorConfig信息
 * param :
 * return : 返回map[interface{}]interface{}
 */
func _GetMonitorConfig() map[interface{}]interface{}{

	//获取环境变量
	requestPath := os.Getenv("CONFIGPATH")
	requestPath = requestPath + "/monitorConfig.yaml"
	config = make(map[interface{}]interface{})
	err := yaml.Read(requestPath,config)
	if err != nil{
		beego.Error(err.Error())
	}

	return config
}

/**
 * function : 获取配置信息中内容
 * param   :
 * return : address,prefix,simpleRate,flushTime
 */
func _GetMonitorParam() (string,string,int){

	//获取MonitorConfig信息
	statsdConfig := _GetMonitorConfig()
	//对获取数据断言处理
	ip :=common.TypeToString(statsdConfig["ip"])
	port :=common.TypeToString(statsdConfig["port"])
	maxPacketSize := common.TypeToInt(statsdConfig["maxPacketSize"])

	address := ip + ":" + port
	//获取系统主机名
	prefix,err := os.Hostname()
	if err != nil{
		beego.Error(err.Error())
	}

	return address,prefix,maxPacketSize
}


/**
 * function : 获取monitorClient客户端
 * param   :
 * return :
 */
func _GetMonitor() *statsd.Client{

	add,pre,maxPacketSize := _GetMonitorParam()
	//准备statsd option
	address := statsd.Address(add)
	prefix := statsd.Prefix(pre)
	packetSize := statsd.MaxPacketSize(maxPacketSize)
	//创建monitor.client对象
	monitor,err := statsd.New(address,prefix,packetSize)
	if err != nil{
		beego.Error(err.Error())
	}
	Monitor = monitor
	return Monitor
}

