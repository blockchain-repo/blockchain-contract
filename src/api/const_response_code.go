package api

//const HTTP_STATUS_CODE_OK = 200 // browser use!

//1位(类型,1~9)+2位(项目编号,01~99)+3位(具体错误码)
// response data code 业务级错误
const (
	RESPONSE_STATUS_OK = 0 //200 - 客户端请求已成功
	// 001~099 HTTP request relation
	RESPONSE_STATUS_CONTENT_TYPE_ERROR         = 201001 // Content-Type 错误
	RESPONSE_STATUS_REQUEST_METHOD_ERROR       = 201002 // request method 错误
	RESPONSE_STATUS_MISSING_REQUIRED_PARAMETER = 201003 // 必须参数丢失, token, sign, timestamp
	RESPONSE_STATUS_INVALID_PARAMETER          = 201004 // 请求参数名称非法
	RESPONSE_STATUS_INVALID_TIMESTAMP          = 201005 // 请求参数 timestamp 非法
	RESPONSE_STATUS_INVALID_TOKEN              = 201006 // 请求参数 token 非法
	RESPONSE_STATUS_INVALID_SIGN               = 201007 // 请求参数 sign 非法
	RESPONSE_STATUS_INVALID_APPID              = 201008 // 请求参数 appId 非法

	// controller PARAMETER_ERROR 101~199
	RESPONSE_STATUS_PARAMETER_ERROR_COUNT    = 201101 // 请求参数个数错误
	RESPONSE_STATUS_PARAMETER_ERROR_TYPE     = 201102 // 请求参数类型/格式 错误
	RESPONSE_STATUS_PARAMETER_ERROR_VALUE    = 201103 // 请求参数值错误, 值为空 或者范围等限制
	RESPONSE_STATUS_PARAMETER_ERROR_LENGTH   = 201104 // 请求参数长度错误
	RESPONSE_STATUS_CONTRACT_ERROR_MODEL     = 201105 // 合约 model 错误, invalid contractModel.Validate() false
	RESPONSE_STATUS_CONTRACT_ERROR_SIGNATURE = 201106 // 合约签名错误

	// DB ERROR 201~299
	RESPONSE_STATUS_DB_ERROR_CONN    = 201201 // 数据库连接错误
	RESPONSE_STATUS_DB_ERROR_TIMEOUT = 201202 // 数据库超时错误
	RESPONSE_STATUS_DB_ERROR_OP      = 201203 // 数据库操作错误
	// INTERNAL ERROR 301~399
	RESPONSE_STATUS_INTERNAL_ERROR = 201301 // 程序内部处理错误
	RESPONSE_STATUS_PROTO_ERROR    = 201302 // proto处理错误
	RESPONSE_STATUS_ERROR          = 201500
)

// 系统级错误代码 101~999
const (
	RESPONSE_STATUS_SYSTEM_ERROR                   = 101101 // 系统错误
	RESPONSE_STATUS_SERVICE_UNAVAILABLE            = 101102 // 服务暂停
	RESPONSE_STATUS_REMOTE_SERVICE_ERROR           = 101103 // 远程服务错误
	RESPONSE_STATUS_IP_LIMIT                       = 101104 // IP限制不能请求该资源
	RESPONSE_STATUS_PERMISSION_DENIED              = 101105 // 该资源需要appkey拥有授权
	RESPONSE_STATUS_MISSING_APPKEY                 = 101106 // 缺少source (appkey) 参数
	RESPONSE_STATUS_UNSUPPORT_MEDIATYPE            = 101107 // 不支持的MediaType
	RESPONSE_STATUS_JOB_EXPIRED                    = 101108 // 任务超时
	RESPONSE_STATUS_RPC_ERROR                      = 101109 // RPC错误
	RESPONSE_STATUS_ILLEGAL_REQUEST                = 101110 // 非法请求
	RESPONSE_STATUS_IP_REQUESTS_OUT_OF_RATE_LIMIT  = 101111 // IP请求频次超过上限
	RESPONSE_STATUS_APP_REQUESTS_OUT_OF_RATE_LIMIT = 101112 // 应用请求频次超过上限
)
