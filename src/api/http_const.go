package api

const (
	HTTP_STATUS_CODE_OK             = 111000200 //200 - 客户端请求已成功
	HTTP_STATUS_CODE_BadRequest     = 111000400 //400 - 请求出现语法错误
	HTTP_STATUS_CODE_Unauthorized   = 111000401 //401 - 访问被拒绝
	HTTP_STATUS_CODE_Forbidden      = 111000403 //403 - 禁止访问 资源不可用
	HTTP_STATUS_CODE_NotFound       = 111000404 //404 - 无法找到指定位置的资源
	HTTP_STATUS_CODE_NotAcceptable  = 111000406 //406 - 指定的资源已经找到，但它的MIME类型和客户在Accpet头中所指定的不兼容
	HTTP_STATUS_CODE_RequestTimeout = 111000408 //408 - 在服务器许可的等待时间内，客户一直没有发出任何请求。客户可以在以后重复同一请求。
)
