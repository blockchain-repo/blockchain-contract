package requestHandler

import (
	"os"
	"unicontract/src/common/yaml"
	"github.com/astaxie/beego"
)


var config map[interface{}]interface{}

/**
 * function : function: 获取yamlconfig中内容
 * param   :
 * return : 将数据写到map中并返回
 */
func GetYamlConfig(yamlName string) map[interface{}]interface{}{

	//获取环境变量
	requestPath := os.Getenv("CONFIGPATH")
	requestPath = requestPath + "/" + yamlName
	config = make(map[interface{}]interface{})
	err := yaml.Read(requestPath,config)
	if err != nil{
		beego.Error(err.Error())
	}

	return config
}