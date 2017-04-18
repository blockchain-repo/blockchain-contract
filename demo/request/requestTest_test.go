package request

import (
	"testing"
)

/**
 * function:
 * param :
 * return nil:
 */
func TestGetValue(t *testing.T) {
	CreatTransaction()
}

func TestWeather(t *testing.T) {
	Weather()
}
func TestToday(t *testing.T) {
	Today()
	//switch statusCode {
	//
	//case statusCode >=100 && statusCode < 200:
	//	//信息，服务器收到请求，需要请求者继续执行操作
	//	message := "Extraordinary Response"
	//	responseData = _ResponseFailDataHandler(statusCode,message)
	//
	//case statusCode >=200 && statusCode < 300:
	//	//成功，操作被成功接收并处理
	//	responseData = _ResponseDataHandler(responseBody)
	//
	//case statusCode >=300 && statusCode < 400:
	//	//重定向，需要进一步的操作以完成请求
	//	message := "Re-order New Destination Site"
	//	responseData = _ResponseFailDataHandler(statusCode,message)
	//
	//case statusCode >=400 && statusCode < 500:
	//	//客户端错误，请求包含语法错误或无法完成请求
	//	message := "Customer Service request Error"
	//	responseData = _ResponseFailDataHandler(statusCode,message)
	//
	//case statusCode >=500 && statusCode < 600:
	//	//服务器错误，服务器在处理请求的过程中发生了错误
	//	message := "Server Service Internal Error"
	//	responseData = _ResponseFailDataHandler(statusCode,message)
	//
	//default:
	//	//其它错误
	//	message := "Unexecption Error"
	//	responseData = _ResponseFailDataHandler(statusCode,message)
	//}
}

func TestTest(t *testing.T) {
	Test()
}