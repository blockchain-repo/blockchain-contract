package requestHandler

import (
	"sync"
	"os"
	"log"
	"unicontract/src/common/yaml"
)

var config map[interface{}]interface{}
var once sync.Once

/**
 * function: 初始化config
 */
func _GetConfig() map[interface{}]interface{}{
	//单例模式
	once.Do(func(){
		//获取环境变量
		requestPath := os.Getenv("CONFIGPATH")
		requestPath = requestPath + "/requestConfig.yaml"
		config = make(map[interface{}]interface{})
		err := yaml.Read(requestPath,config)
		if err != nil{
			log.Fatal(err.Error())
		}
	})

	return config
}

/**
 * function : 断言-类型转换string
 * param   :
 * return : 返回string
 */

func TypeToString(name interface{}) string{

	value,ok := name.(string)
	if !ok {
		log.Fatal("Type conversion error")
	}
	return value
}


/**
 * function : 获取ip
 * param   :
 * return : 返回ip
 */

func _GetIp() string{

	return TypeToString(config["ip"])
}


/**
 * function : 获取Port
 * param   :
 * return : 返回Port
 */

func _GetPort() string{

	return TypeToString(config["port"])
}


/**
 * function : 获取head
 * param   :
 * return : 返回head
 */

func _GetHead() map[interface{}]interface{}{

	head,ok := config["head"].(map[interface{}]interface{})
	if !ok{
		log.Fatal("Type conversion error")
	}

	return head
}

/**
 * function : 获取url
 * param   :
 * return : 返回url
 */

func _GetUrl(path string) string{
	url := "http://" + _GetIp() + ":" + _GetPort() + path
	return url
}

/**
 * function : 获取param
 * param   :
 * return : 返回param
 */

func GetParam(apiName string) (string,string,map[interface{}]interface{},map[interface{}]interface{}){

	config := _GetConfig()
	api := config[apiName]
	value,ok := api.(map[interface{}]interface{})
	if !ok {
		log.Fatal("Type conversion error")
	}
	jsonBody := value["jsonBody"]

	body,ok1 := jsonBody.(map[interface{}]interface{})
	if !ok1 {
		log.Fatal("Type conversion error")
	}

	path := TypeToString(value["path"])
	method := TypeToString(value["method"])

	url := _GetUrl(path)
	head := _GetHead()

	return method,url,head,body
}





