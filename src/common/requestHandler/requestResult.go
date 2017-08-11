package requestHandler

/**
 * function : 获取请求结果并返回
 * param   : yamlName,apiName,jsonBody
 * return : responseResult struct
 */
func GetRequestResult(yamlName string, apiName string, jsonBody interface{}, chainType string) *ResponseResult {

	//读取yaml配置文件
	yamlConfig := GetYamlConfig(yamlName)
	//获取api请求参数
	method, url, head, _ := GetParam(yamlConfig, apiName, chainType)
	//创建请求参数struct
	param := NewRequestParam(method, url, head, jsonBody)
	//获取请求参数结果
	responseBody, statusCode := RequestHandler(param)
	//对请求参数进行处理,并返回response struct
	responseResult := GetResponseData(responseBody, statusCode)
	return responseResult
}

/**
 * function :获取请求结果(string类型)以及statusCode
 * param   : yamlName,apiName,jsonBody
 * return : responseBody(string),status(int)
 */
func GetRequestResult1(yamlName string, apiName string, jsonBody interface{}, chainType string) (string, int) {

	yamlConfig := GetYamlConfig(yamlName)
	method, url, head, _ := GetParam(yamlConfig, apiName, chainType)
	param := NewRequestParam(method, url, head, jsonBody)
	responseBody, statusCode := RequestHandler(param)
	responseResult, status := GetResponseString(responseBody, statusCode)

	return responseResult, status
}
