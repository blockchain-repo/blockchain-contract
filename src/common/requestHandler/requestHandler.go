package requestHandler

import (
	"strings"
	"github.com/astaxie/beego/httplib"
	"unicontract/src/common"
	"github.com/astaxie/beego"
)


/**
 * 请求参数struct
 */
type RequestParam struct {
	Method       string
	URL          string
	Heads        map[interface{}]interface{}
	JsonBody     interface{}
}

/**
 * function: 初始化RequestParam
 * param : method,url,headKey,headValue,jsonBody
 * return: 返回request param参数struct
 */
func NewRequestParam(method string,url string, head map[interface{}]interface{},jsonBody interface{}) *RequestParam{

	param := &RequestParam{
		method,
		url,
		head,
		jsonBody,
	}
	return param
}

/**
 * function: 判定选择哪种请求方式get,post,put,delete
 * param : requestParam结构体
 * return: *http.Response,error
 */
func RequestHandler(requestParam *RequestParam) (string,int){
	method := strings.ToUpper(requestParam.Method)
	requestParam.Method = method
	if method == "GET" {
		return _Get(requestParam)
	} else if method == "POST" {
		return _Post(requestParam)
	} else if method == "PUT" {
		return _Put(requestParam)
	}else if method == "DELETE" {
		return _Delete(requestParam)
	}else {
		return "",0
	}
}

/**
 * function: get方法
 * param : requestParam结构体
 * return: *http.Response,error
 */
func _Get(requestParam *RequestParam) (string,int) {

	request := httplib.Get(requestParam.URL)

	return _GetResponse(request)
}

/**
 * function: post方法
 * param : requestParam结构体
 * return: *http.Response,error
 */
func _Post(requestParam *RequestParam) (string,int) {

	request := httplib.Post(requestParam.URL)
	_SetParam(request,requestParam)

	return _GetResponse(request)
}

/**
 * function: put方法
 * param : requestParam结构体
 * return: *http.Response,error
 */
func _Put(requestParam *RequestParam) (string,int){

	request := httplib.Put(requestParam.URL)
	_SetParam(request,requestParam)

	return _GetResponse(request)

}

/**
 * function: delete方法
 * param : requestParam结构体
 * return: *http.Response,error
 */
func _Delete(requestParam *RequestParam) (string,int){

	request := httplib.Delete(requestParam.URL)
	_SetParam(request,requestParam)

	return _GetResponse(request)
}

/**
 * function: 设置参数
 * param : request,requestParam结构体
 * return: *http.Response,error
 */
func _SetParam(request *httplib.BeegoHTTPRequest,requestParam *RequestParam){

	//设置请求头
	heads := requestParam.Heads
	if heads != nil{
		for k,v := range requestParam.Heads{

			request.Header(common.TypeToString(k),common.TypeToString(v))
		}
	}

	//设置参数
	jsonBody := requestParam.JsonBody
	switch jsonBody.(type) {
	//body可以是string,[]byte类型,也可以在配置文件中配置key-value形式
	case string:
		request.Body(requestParam.JsonBody)
	case []byte:
		request.Body(requestParam.JsonBody)
	case map[interface{}]interface{}:
		body := common.TypeToMap(jsonBody)
		for k,v := range body{
			request.Param(common.TypeToString(k),common.TypeToString(v))
		}

	default:
		beego.Debug("Type error,please input string or map..")
	}

}

/**
 * function: 获取response
 * param : requestParam结构体
 * return: *http.Response,error
 */
func _GetResponse(request *httplib.BeegoHTTPRequest) (string,int){

	jsonResponse,_ := request.String()
	respond,err := request.Response()
	if err != nil{
		beego.Error(err.Error())
	}

	return jsonResponse,respond.StatusCode
}

