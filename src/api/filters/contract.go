package filters

import (
	//"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"hash"
	"strconv"
	"strings"
	"time"
	api "unicontract/src/api"
	"unicontract/src/common/uniledgerlog"
)

var (
	API_TIMEOUT       = int64(60)
	API_TOKEN_LEN     = 44
	API_SIGN_LEN      = 32
	API_TIMESTAMP_LEN = 13
)

// api auth verify
var (
	AUTH_VERIFY            = true
	AUTH_VERIFY_HTTP       = true
	AUTH_VERIFY_TIMESTAMP  = true
	AUTH_VERIFY_PARAMETERS = false
	AUTH_VERIFY_SIGN       = false
	AUTH_VERIFY_TOKEN      = false
	AUTH_VERIFY_RATE_LIMIT = false
)

type response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"data"`
}

func init() {
	API_TIMEOUT = beego.AppConfig.DefaultInt64("api_timeout", API_TIMEOUT)

	AUTH_VERIFY = beego.AppConfig.DefaultBool("auth_verify", AUTH_VERIFY)
	AUTH_VERIFY_HTTP = beego.AppConfig.DefaultBool("auth_verify_http", AUTH_VERIFY_HTTP)
	AUTH_VERIFY_TIMESTAMP = beego.AppConfig.DefaultBool("auth_verify_timestamp", AUTH_VERIFY_TIMESTAMP)
	AUTH_VERIFY_PARAMETERS = beego.AppConfig.DefaultBool("auth_verify_parameters", AUTH_VERIFY_PARAMETERS)
	AUTH_VERIFY_SIGN = beego.AppConfig.DefaultBool("auth_verify_sign", AUTH_VERIFY_SIGN)
	AUTH_VERIFY_TOKEN = beego.AppConfig.DefaultBool("auth_verify_token", AUTH_VERIFY_TOKEN)
	AUTH_VERIFY_RATE_LIMIT = beego.AppConfig.DefaultBool("auth_verify_rate_limit", AUTH_VERIFY_RATE_LIMIT)
}

func responseWithStatusCode(ctx *context.Context, status int, msg string) {
	result := response{Code: status, Msg: "", Result: msg}
	resultByte, err := json.Marshal(result)
	if err != nil {
		uniledgerlog.Error("responseWithStatusCode", err.Error())
	}
	ctx.ResponseWriter.Write(resultByte)
	return
}

//var ContentTypes = []string{"application/json", "application/x-protobuf"}
func APIHttpFilter(ctx *context.Context) {
	//cost_start := time.Now()

	contentType := ctx.Input.Header("Content-Type")
	if contentType == "" {
		resultMsg := fmt.Sprintf("%s 请求头 Content-Type 为空!", "Filter[APIContentTypeFilter]")
		uniledgerlog.Error(resultMsg)
		responseWithStatusCode(ctx, api.RESPONSE_STATUS_CONTENT_TYPE_ERROR, resultMsg)
		//defer api.TimeCost(cost_start, ctx, api.RESPONSE_STATUS_CONTENT_TYPE_ERROR, resultMsg)()
		return

	} else if contentType == "application/json" || contentType == "application/x-protobuf" {
		//uniledgerlog.Debug("RequestDataType is json!")
	} else {
		resultMsg := fmt.Sprintf("%s 请求头 Content-Type 值错误!", "Filter[APIContentTypeFilter]")
		uniledgerlog.Error(resultMsg)
		responseWithStatusCode(ctx, api.RESPONSE_STATUS_UNSUPPORT_MEDIATYPE, resultMsg)
		//defer api.TimeCost(cost_start, ctx, api.RESPONSE_STATUS_CONTENT_TYPE_ERROR, resultMsg)()
		return
	}
}

//todo  step 3
func APITimestampFilter(ctx *context.Context) {
	// 1. verify the timestamp format
	timestamp := ctx.Input.Query("timestamp")
	if len(timestamp) != API_TIMESTAMP_LEN {
		resultMsg := fmt.Sprintf("%s timestamp length error!", "Filter[APITimestampFilter]")
		uniledgerlog.Error(resultMsg)
		responseWithStatusCode(ctx, api.RESPONSE_STATUS_INVALID_TIMESTAMP, resultMsg)
		return
	}
	timestamp_int64, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		resultMsg := fmt.Sprintf("%s timestamp format error!", "Filter[APITimestampFilter]")
		uniledgerlog.Error(resultMsg)
		responseWithStatusCode(ctx, api.RESPONSE_STATUS_INVALID_TIMESTAMP, resultMsg)
		return
	}

	// 2. verify the timeout
	nanos := time.Now().UnixNano()
	//ms len=13
	current_unix_timestamp := nanos / 1000000
	cost := (current_unix_timestamp - timestamp_int64) / 1000
	uniledgerlog.Debug("time info", current_unix_timestamp, timestamp_int64, cost)
	if cost < 0 || cost > API_TIMEOUT {
		resultMsg := fmt.Sprintf("%s timestamp value error!", "Filter[APITimestampFilter]")
		uniledgerlog.Error(resultMsg, cost)
		responseWithStatusCode(ctx, api.RESPONSE_STATUS_INVALID_TIMESTAMP, resultMsg)
		return
	}
}

//todo  step 4
func APIParametersFilter(ctx *context.Context) {
	// params
	ctx.Request.ParseForm()
	parameters := ctx.Request.Form
	if len(parameters) > 0 {
		for k, _ := range parameters {
			if !api.ALLOW_REQUEST_PARAMETERS_ALL[k] {
				resultMsg := fmt.Sprintf("%s keyword %s is not allowed here!", "Filter[APIParametersFilter]", k)
				uniledgerlog.Error(resultMsg)
				responseWithStatusCode(ctx, api.RESPONSE_STATUS_INVALID_PARAMETER, resultMsg)
				return
			}
		}
	}
}

func generateSign(val string) string {
	var hash hash.Hash
	var x string
	hash = md5.New()
	if hash != nil {
		hash.Write([]byte(val))
		x = hex.EncodeToString(hash.Sum(nil))
	}
	return x
}

//todo  step 5
func APISignFilter(ctx *context.Context) {
	// params
	ctx.Request.ParseForm()
	parameters := ctx.Request.Form
	uniledgerlog.Warn(parameters.Encode())

	sign := ctx.Input.Query(api.REQUEST_FIELD_AUTH_SIGN)
	if len(sign) != API_SIGN_LEN {
		resultMsg := fmt.Sprintf("%s sign length error!", "Filter[APISignFilter]")
		uniledgerlog.Error(resultMsg)
		responseWithStatusCode(ctx, api.RESPONSE_STATUS_INVALID_SIGN, resultMsg)
		return
	}

	// get real sign
	parameters.Del(api.REQUEST_FIELD_AUTH_SIGN)
	realSign := parameters.Encode()
	realSign = strings.ToUpper(generateSign(realSign))
	if sign != realSign {
		resultMsg := fmt.Sprintf("%s sign(%s)错误!", "Filter[APISignFilter]", sign)
		uniledgerlog.Error(resultMsg, "input is "+sign+", real is "+realSign)
		responseWithStatusCode(ctx, api.RESPONSE_STATUS_INVALID_SIGN, resultMsg)
		return
	}

}

// todo get the token with app_id and app_key; if not exist, response 111000403.
// APIAuthorizationFilter step 2 授权,请求token
// accesskey need store in redis for next step verify accessKey
//func APIAuthorizationFilter(ctx *context.Context) {
//	cost_start := time.Now()
//	app_id := ctx.Input.Query("app_id")
//	app_key := ctx.Input.Query("app_key")
//
//	/******************* test data, each has the static  *******************/
//	if app_id == "" {
//		app_id = "0123456789"
//	}
//	if app_key == "" {
//		app_key = "uni-ledger.com"
//	}
//	/******************* test data *******************/
//
//	exist := api.CheckExistAppUser(api.GenerateAccessKey(app_id, app_key))
//
//	if exist {
//		access_key := api.GenerateAccessKey(app_id, app_key)
//		uniledgerlog.Info("APIAuthorizationFilter exist user!", access_key)
//		//todo
//		_ = api.StoreAccessKey(app_id, access_key)
//		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_OK, access_key)
//	} else {
//		uniledgerlog.Error("APIAuthorizationFilter not exist user!")
//		// not exist, generateToken and put redis
//		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_OK, "not exist user!")
//		defer api.TimeCost(cost_start, ctx, api.RESPONSE_STATUS_Forbidden)()
//		return
//	}
//}
//
//func APIGetTokenFilter(ctx *context.Context) {
//	cost_start := time.Now()
//	access_key := ctx.Input.Query("accessKey")
//	app_id := ctx.Input.Query("app_id")
//	if app_id == "" {
//		app_id = "0123456789"
//	}
//	if access_key != "" {
//		token, ok := api.GetToken(app_id, access_key)
//		if ok {
//			uniledgerlog.Info("APIGetTokenFilter exist accessKey!", token)
//			responseWithStatusCode(ctx, api.RESPONSE_STATUS_OK, token)
//
//		} else {
//			uniledgerlog.Info("APIGetTokenFilter not exist accessKey!", token)
//			responseWithStatusCode(ctx, api.RESPONSE_STATUS_BadRequest, token)
//		}
//	} else {
//		uniledgerlog.Error("APIGetTokenFilter accessKey error!")
//		// not exist, generateToken and put redis
//		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_OK, "accessKey error!")
//		defer api.TimeCost(cost_start, ctx, api.RESPONSE_STATUS_Forbidden)()
//		return
//	}
//
//}
//
//// APIAuthorizationFilter step 3 API 参数及权限认证
//func APIAuthFilter(ctx *context.Context) {
//	// 1. verify the parameters is miss, timestamp, token,sign
//	token := ctx.Input.Query("token")
//	timestamp := ctx.Input.Query("timestamp")
//	sign := ctx.Input.Query("sign")
//
//	if token == "" || timestamp == "" || sign == "" {
//		uniledgerlog.Error("APIAuthFilter parameters miss!")
//		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter parameters miss!")
//		return
//	}
//	if len(token) != API_TOKEN_LEN {
//		uniledgerlog.Error("APIAuthFilter parameters token error!")
//		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter parameters token error!")
//		return
//	}
//
//	//if len(sign) != API_SIGN_LEN {
//	//	uniledgerlog.Error("APIAuthFilter parameters sign error!")
//	//	responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter parameters sign error!")
//	//	return
//	//}
//
//	//if len(token) != API_TOKEN_LEN || len(timestamp) != API_TIMESTAMP_LEN || len(sign) != API_SIGN_LEN {
//	//	uniledgerlog.Error("APIAuthFilter parameters error!")
//	//	responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter parameters error!")
//	//	return
//	//}
//
//	// 3. verify the token
//	if !api.ExistKey(token) {
//		uniledgerlog.Error("APIAuthFilter token not exist!")
//		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter token not exist! "+strconv.FormatInt(cost, 10)+" s")
//		return
//	}
//	// 4. verify the sign
//	if !api.VerifySign(token, timestamp, sign) {
//		uniledgerlog.Error("APIAuthFilter sign error!", cost)
//		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter sign error!"+strconv.FormatInt(cost, 10)+"s")
//		return
//	}
//}
//
//// APIRateLimitFilter step 4 API Ratelimit 验证
//func APIRateLimitFilter(ctx *context.Context) {
//	//cost_start := time.Now()
//	token := ctx.Input.Query("token")
//	if !api.RateLimit(token) {
//		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_TOO_MANY_REQUESTS, "APIRateLimitFilter error!"+string(token))
//		//defer api.TimeCost(cost_start, ctx, api.HTTP_STATUS_CODE_TOO_MANY_REQUESTS)()
//		return
//	}
//	// last filter, will reset the token valid time
//	api.UpdateToken(token)
//}
