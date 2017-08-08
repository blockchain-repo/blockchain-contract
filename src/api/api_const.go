package api

// origin http status
const (
	HTTP_STATUS_CODE_OK         = 200 //200 - 客户端请求已成功
	HTTP_STATUS_CODE_BadRequest = 400 //400 - 请求出现语法错误
	//HTTP_STATUS_CODE_Unauthorized      = 111000401 //401 - 访问被拒绝
	HTTP_STATUS_CODE_Forbidden = 403 //403 - 禁止访问 资源不可用
	//HTTP_STATUS_CODE_NotFound          = 404 //404 - 无法找到指定位置的资源
	//HTTP_STATUS_CODE_NotAcceptable     = 406 //406 - 指定的资源已经找到，但它的MIME类型和客户在Accpet头中所指定的不兼容
	//HTTP_STATUS_CODE_RequestTimeout    = 408 //408 - 在服务器许可的等待时间内，客户一直没有发出任何请求。客户可以在以后重复同一请求。
	HTTP_STATUS_CODE_TOO_MANY_REQUESTS = 429 // 429 - 你需要限制客户端请求某个服务数量时，该状态码就很有用，也就是请求速度限制。
)

// response data code
const (
	RESPONSE_STATUS_QUERY_ERROR   = 211000001 // query error
	RESPONSE_STATUS_INSERT_ERROR  = 211000002 // insert error
	RESPONSE_STATUS_CONVERT_ERROR = 211000003 // proto convert error

)

// response data code
const (
	RESPONSE_STATUS_OK                = 0         //200 - 客户端请求已成功
	RESPONSE_STATUS_BadRequest        = 111000400 //400 - 请求出现语法错误
	RESPONSE_STATUS_Unauthorized      = 111000401 //401 - 访问被拒绝
	RESPONSE_STATUS_Forbidden         = 111000403 //403 - 禁止访问 资源不可用
	RESPONSE_STATUS_NotFound          = 111000404 //404 - 无法找到指定位置的资源
	RESPONSE_STATUS_NotAcceptable     = 111000406 //406 - 指定的资源已经找到，但它的MIME类型和客户在Accpet头中所指定的不兼容
	RESPONSE_STATUS_RequestTimeout    = 111000408 //408 - 在服务器许可的等待时间内，客户一直没有发出任何请求。客户可以在以后重复同一请求。
	RESPONSE_STATUS_TOO_MANY_REQUESTS = 111000429 // 429 - 你需要限制客户端请求某个服务数量时，该状态码就很有用，也就是请求速度限制。
)

const (
	// 10 hours
	API_TIMEOUT       = int64(60 * 60 * 10)
	API_SIGN_LEN      = int(32)
	API_TOKEN_LEN     = int(44)
	API_TIMESTAMP_LEN = int(13)
)

// ratelimit token
const (
	// 60s temp store for apply the token
	access_key_timeout = 60
	// confuse the access_key, as  app_id{access_key_blur}access_key
	access_key_blur = "uniledger"
	// 30 minutes
	token_timeout = 60 * 30
	// {token}_rate
	rate_token_key = "_rate"
	// 2s
	rate_limit_duration = 10
	rate_limit_count    = 100
)
