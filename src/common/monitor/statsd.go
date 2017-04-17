package monitor

import (
	"github.com/alexcesaro/statsd"
	"os"
	"log"
	"sync"
	"unicontract/src/common/yaml"
	"unicontract/src/common"
)

var (
	//monitor   *statsd.Client
	//err error
	once sync.Once
	config map[interface{}]interface{}
)

/**
 * function: 初始化config
 */
func _GetConfig() map[interface{}]interface{}{
	//单例模式
	once.Do(func(){
		//获取环境变量
		requestPath := os.Getenv("CONFIGPATH")
		requestPath = requestPath + "/statsdConfig.yaml"
		config = make(map[interface{}]interface{})
		err := yaml.Read(requestPath,config)
		if err != nil{
			log.Fatal(err.Error())
		}
	})

	return config
}

/**
 * function :
 * param   :
 * return :
 */
func GetMonitor(){

	//add,pre,rate,flus := _GetParam()
	//
	//ftime := 1*time.Millisecond
	//address := statsd.Address(add)
	//prefix := statsd.Prefix(pre)
	//simpleSate := statsd.SampleRate(rate)
	//flushTime := statsd.FlushPeriod(flus*time.Millisecond)
	//
	//monitor,err = statsd.New(address,prefix,simpleSate,flushTime)
}

func _GetParam() (string,string,float32,int){

	//读取配置文件
	statsdConfig := _GetConfig()
	ip :=common.TypeToString(statsdConfig["ip"])
	port :=common.TypeToString(statsdConfig["port"])
	simpleRate :=common.TypeToFloat32(statsdConfig["simpleRate"])
	flushTime := common.TypeToInt(statsdConfig["flushTime"])

	address := ip + ":" + port

	return address,"",simpleRate,flushTime
}
