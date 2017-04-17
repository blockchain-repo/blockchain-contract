package requestHandler

import (
	"encoding/json"
	"unicontract/src/common"
	"github.com/astaxie/beego"
)

/**
 * 响应结果struct
 */
type ResponseResult struct {
	Status       string
	Code         int
	Message      string
	Data         interface{}
}

/**
 * function: 初始化响应结果
 * param : status,code,message,data
 * return: 返回response result
 */
func _NewResponseResult(status string,code int,message string,data interface{}) *ResponseResult{

	result := &ResponseResult{
		status,
		code,
		message,
		data,
	}
	return result
}

/**
 * function: 获取response string
 * param : jsonBody
 * return: 返回ResponseResult string
 */
func GetResponseString(responseBody string,statusCode int) (string,int){

	message := _StatusHandler(responseBody,statusCode)

	return message,statusCode
}

/**
 * function: 获取response struct
 * param : jsonBody
 * return: 返回ResponseResult struct
 */
func GetResponseData(responseBody string,statusCode int) *ResponseResult{

	var responseData *ResponseResult

	message := _StatusHandler(responseBody,statusCode)
	if statusCode >=200 && statusCode < 300{
		responseData = _ResponseDataHandler(message)
	}else {
		responseData = _ResponseFailDataHandler(statusCode,message)
	}

	return responseData
}

/**
 * function: 处理失败请求,并返回相应的数据
 * param : responseBody
 * return: 返回ResponseResult struct
 */
func _StatusHandler(responseBody string,statusCode int) string{

	var message string

	if statusCode >=100 && statusCode < 200{

		//信息，服务器收到请求，需要请求者继续执行操作
		message = "Extraordinary Response"
	}else if statusCode >=200 && statusCode < 300{

		//成功，操作被成功接收并处理
		message = responseBody
	}else if statusCode >=300 && statusCode < 400{

		////重定向，需要进一步的操作以完成请求
		message = "Re-order New Destination Site"
	}else if statusCode >=400 && statusCode < 500{

		//客户端错误，请求包含语法错误或无法完成请求
		message = "Customer Service request Error"
	}else if statusCode >=500 && statusCode < 600{

		//服务器错误，服务器在处理请求的过程中发生了错误
		message = "Server Service Internal Error"
	}else {

		//其它错误
		message = "Unexecption Error"
	}
	return message
}


/**
 * function: 处理失败请求,并返回相应的数据
 * param : responseBody
 * return: 返回ResponseResult struct
 */
func _ResponseFailDataHandler(statusCode int,message string) *ResponseResult{

	status := "Fail"
	data := "null"

	return _NewResponseResult(status,statusCode,message,data)
}

/**
 * function: 请求成功,解析相应数据
 * param : responseBody
 * return: 返回ResponseResult struct
 */
func _ResponseDataHandler(responseBody string) *ResponseResult{

	//创建map,将返回的jsonString解析到map中
	responseMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(responseBody),&responseMap)
	if err != nil{
		beego.Error(err.Error())
	}
	//将数据取出来放到responseResult struct中
	status := common.TypeToString(responseMap["status"])

	//处理接收数据为int
	code := responseMap["code"]
	statusCode,ok:= code.(float64)
	var intCode int
	if !ok{
		intCode = common.TypeToInt(code)
	}
	intCode = int(statusCode)

	//对message,data断言处理
	message := common.TypeToString(responseMap["message"])
	data := responseMap["data"]

	return _NewResponseResult(status,intCode,message,data)
}
