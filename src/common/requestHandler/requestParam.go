package requestHandler

import (
	"unicontract/src/common"
)

/**
 * function : 获取ip
 * param   :
 * return : 返回ip
 */
func _GetIp(yamlConfig map[interface{}]interface{}) string{

	return common.TypeToString(yamlConfig["ip"])
}


/**
 * function : 获取Port
 * param   :
 * return : 返回Port
 */
func _GetPort(yamlConfig map[interface{}]interface{}) string{

	return common.TypeToString(yamlConfig["port"])
}


/**
 * function : 获取head
 * param   :
 * return : 返回head
 */
func _GetHead(yamlConfig map[interface{}]interface{}) map[interface{}]interface{}{

	head := common.TypeToMap(yamlConfig["head"])

	return head
}

/**
 * function : 获取url
 * param   :
 * return : 返回url
 */
func _GetUrl(yamlConfig map[interface{}]interface{},path string) string{
	url := "http://" + _GetIp(yamlConfig) + ":" + _GetPort(yamlConfig) + path
	return url
}

/**
 * function : 获取param
 * param   :
 * return : 返回param
 */
func GetParam(yamlConfig map[interface{}]interface{},apiName string) (string,string,map[interface{}]interface{},map[interface{}]interface{}){

	//获取yaml文件中对用的api信息
	api := yamlConfig[apiName]
	//断言,判断是否是map类型
	value := common.TypeToMap(api)
	//json参数
	jsonBody := value["jsonBody"]
	body := common.TypeToMap(jsonBody)
	//获取api的path路径信息
	path := common.TypeToString(value["path"])
	//获取api的method类型
	method := common.TypeToString(value["method"])
	//获取请求url
	url := _GetUrl(yamlConfig,path)
	//获取请求头
	head := _GetHead(yamlConfig)

	return method,url,head,body
}





