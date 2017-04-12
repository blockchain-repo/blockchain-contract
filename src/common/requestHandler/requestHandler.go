package requestHandler

import (
	"strings"
	"github.com/astaxie/beego/httplib"
	"fmt"
	"log"
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
func RequestHandler(requestParam *RequestParam) (string,string){
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
		return "",fmt.Sprintf("%s is unsupported method", requestParam.Method)
	}
}

/**
 * function: get方法
 * param : requestParam结构体
 * return: *http.Response,error
 */
func _Get(requestParam *RequestParam) (string,string) {

	request := httplib.Get(requestParam.URL)
	jsonResponse,_ := request.String()

	return jsonResponse,_GetResponse(request)
}

/**
 * function: post方法
 * param : requestParam结构体
 * return: *http.Response,error
 */
func _Post(requestParam *RequestParam) (string,string) {

	request := httplib.Post(requestParam.URL)
	_SetParam(request,requestParam)
	jsonResponse,_ := request.String()

	return jsonResponse,_GetResponse(request)
}

/**
 * function: put方法
 * param : requestParam结构体
 * return: *http.Response,error
 */
func _Put(requestParam *RequestParam) (string,string){

	request := httplib.Put(requestParam.URL)
	_SetParam(request,requestParam)
	jsonResponse,_ := request.String()

	return jsonResponse,_GetResponse(request)

}

/**
 * function: delete方法
 * param : requestParam结构体
 * return: *http.Response,error
 */
func _Delete(requestParam *RequestParam) (string,string){

	request := httplib.Delete(requestParam.URL)
	_SetParam(request,requestParam)
	jsonResponse,_ := request.String()

	return jsonResponse,_GetResponse(request)
}

/**
 * function: 设置参数
 * param : request,requestParam结构体
 * return: *http.Response,error
 */
func _SetParam(request *httplib.BeegoHTTPRequest,requestParam *RequestParam){

	heads := requestParam.Heads
	if heads != nil{
		for k,v := range requestParam.Heads{

			request.Header(TypeToString(k),TypeToString(v))
		}
	}

	jsonBody := requestParam.JsonBody
	switch jsonBody.(type) {
	case string:
		request.Body(requestParam.JsonBody)

	case map[interface{}]interface{}:
		body,ok := jsonBody.(map[interface{}]interface{})
		if !ok{
			log.Fatal("Type conversion error")
		}
		for k,v := range body{
			key,_ := k.(string)
			value,_ := v.(string)
			request.Param(key,value)
		}

	default:
		log.Fatal("Type error,please input string or map..")
	}

}

/**
 * function: 获取response
 * param : requestParam结构体
 * return: *http.Response,error
 */
func _GetResponse(request *httplib.BeegoHTTPRequest) string{
	respond,err := request.Response()
	if err != nil{
		log.Fatal(err.Error())
	}
	return respond.Status
}

