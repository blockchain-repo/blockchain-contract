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
	"net/url"
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
	AUTH_VERIFY                    = true
	AUTH_VERIFY_HEADER             = true
	AUTH_VERIFY_TIMESTAMP          = true
	AUTH_VERIFY_BASIC_PARAMETERS   = false
	AUTH_VERIFY_ALLOWED_PARAMETERS = false
	AUTH_VERIFY_SIGN               = false
	AUTH_VERIFY_TOKEN              = false
	AUTH_VERIFY_RATE_LIMIT         = false
)

type response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"data"`
}

func init() {
	API_TIMEOUT = beego.AppConfig.DefaultInt64("api_timeout", API_TIMEOUT)

	AUTH_VERIFY = beego.AppConfig.DefaultBool("auth_verify", AUTH_VERIFY)
	AUTH_VERIFY_HEADER = beego.AppConfig.DefaultBool("auth_verify_header", AUTH_VERIFY_HEADER)
	AUTH_VERIFY_TIMESTAMP = beego.AppConfig.DefaultBool("auth_verify_timestamp", AUTH_VERIFY_TIMESTAMP)
	AUTH_VERIFY_BASIC_PARAMETERS = beego.AppConfig.DefaultBool("auth_verify_basic_parameters", AUTH_VERIFY_BASIC_PARAMETERS)
	AUTH_VERIFY_ALLOWED_PARAMETERS = beego.AppConfig.DefaultBool("auth_verify_all_parameters", AUTH_VERIFY_ALLOWED_PARAMETERS)
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

func verifyContentType(contentType string) (int, string) {
	if contentType == "" {
		resultMsg := fmt.Sprintf("%s 请求头 Content-Type 为空!", "Filter[APIBasicFilter]")
		return api.RESPONSE_STATUS_CONTENT_TYPE_ERROR, resultMsg

	} else if contentType == "application/json" || contentType == "application/x-protobuf" {
		return api.RESPONSE_STATUS_OK, ""
	} else {
		resultMsg := fmt.Sprintf("%s 请求头 Content-Type 值错误!", "Filter[APIBasicFilter]")
		return api.RESPONSE_STATUS_UNSUPPORT_MEDIATYPE, resultMsg
	}
}

func verifyTimestamp(timestamp string, api_timeout int64, api_timestamp_len int) (int, string) {
	// 1. verify the timestamp format
	if len(timestamp) != api_timestamp_len {
		resultMsg := fmt.Sprintf("%s timestamp length error!", "Filter[APIBasicFilter]")
		return api.RESPONSE_STATUS_INVALID_TIMESTAMP, resultMsg
	}
	timestamp_int64, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		resultMsg := fmt.Sprintf("%s timestamp format error!", "Filter[APIBasicFilter]")
		return api.RESPONSE_STATUS_INVALID_TIMESTAMP, resultMsg
	}
	// 2. verify the timeout
	nanos := time.Now().UnixNano()
	//ms len=13
	current_unix_timestamp := nanos / 1000000
	cost := (current_unix_timestamp - timestamp_int64) / 1000
	uniledgerlog.Debug("time info", current_unix_timestamp, timestamp_int64, cost)
	if cost < 0 || cost > api_timeout {
		resultMsg := fmt.Sprintf("%s timestamp value error!", "Filter[APIBasicFilter]")
		uniledgerlog.Debug(resultMsg, cost)
		return api.RESPONSE_STATUS_INVALID_TIMESTAMP, resultMsg
	}
	return 0, ""
}

func verifyBasicParameters(appId string, timestamp string, token string, sign string) (int, string) {
	if len(appId) != 0 {
		resultMsg := fmt.Sprintf("%s appId length error!", "Filter[APIBasicFilter]")
		return api.RESPONSE_STATUS_INVALID_TIMESTAMP, resultMsg
	}

	if len(timestamp) != API_TIMESTAMP_LEN {
		resultMsg := fmt.Sprintf("%s timestamp length error!", "Filter[APIBasicFilter]")
		return api.RESPONSE_STATUS_INVALID_TIMESTAMP, resultMsg
	}

	if len(sign) != API_SIGN_LEN {
		resultMsg := fmt.Sprintf("%s sign length error!", "Filter[APIBasicFilter]")
		return api.RESPONSE_STATUS_INVALID_SIGN, resultMsg
	}
	if len(token) != API_TOKEN_LEN {
		resultMsg := fmt.Sprintf("%s token length error!", "Filter[APIBasicFilter]")
		return api.RESPONSE_STATUS_INVALID_TOKEN, resultMsg
	}
	return 0, ""
}

func verifyAllowRequestParameters(parameters url.Values) (int, string) {
	// params
	if len(parameters) > 0 {
		for k, _ := range parameters {
			if !api.ALLOW_REQUEST_PARAMETERS_ALL[k] {
				resultMsg := fmt.Sprintf("%s keyword %s is not allowed here!", "Filter[APIBasicFilter]", k)
				return api.RESPONSE_STATUS_INVALID_PARAMETER, resultMsg
			}
		}
	}
	return 0, ""
}

func generateSign(val string) string {
	var hashObj hash.Hash
	var x string
	hashObj = md5.New()
	if hashObj != nil {
		hashObj.Write([]byte(val))
		x = hex.EncodeToString(hashObj.Sum(nil))
	}
	return x
}

func verifySign(parameters url.Values, sign string) (int, string) {
	// get real sign
	parameters.Del(api.REQUEST_FIELD_AUTH_SIGN)
	realSign := parameters.Encode()
	realSign = strings.ToUpper(generateSign(realSign))
	if sign != realSign {
		resultMsg := fmt.Sprintf("%s sign(%s)错误!", "Filter[APISignFilter]", sign)
		uniledgerlog.Debug(resultMsg, "input is "+sign+", real is "+realSign)
		return api.RESPONSE_STATUS_INVALID_SIGN, resultMsg
	}
	return 0, ""
}

// verify the content-type, sign, timestamp and token
func APIBasicFilter(ctx *context.Context) {
	statusCode, resultMsg := 0, ""

	// 1. verify the content-type
	if AUTH_VERIFY_HEADER {
		contentType := ctx.Input.Header("Content-Type")
		//todo
		statusCode, resultMsg = verifyContentType(contentType)
		if statusCode != api.RESPONSE_STATUS_OK {
			uniledgerlog.Error(resultMsg)
			responseWithStatusCode(ctx, statusCode, resultMsg)
			return
		}

	}
	// 2. verify the timestamp
	timestamp := ctx.Input.Query(api.REQUEST_FIELD_AUTH_TIMESTAMP)

	if AUTH_VERIFY_TIMESTAMP {
		statusCode, resultMsg = verifyTimestamp(timestamp, API_TIMEOUT, API_TIMESTAMP_LEN)
		if statusCode != api.RESPONSE_STATUS_OK {
			uniledgerlog.Error(resultMsg)
			responseWithStatusCode(ctx, statusCode, resultMsg)
			return
		}

	}
	// 3. verify the basic required parameters: timestamp, token, sign
	token := ctx.Input.Query(api.REQUEST_FIELD_AUTH_TOKEN)
	sign := ctx.Input.Query(api.REQUEST_FIELD_AUTH_SIGN)
	appId := ctx.Input.Query(api.REQUEST_FIELD_AUTH_APPID)
	if AUTH_VERIFY_BASIC_PARAMETERS {
		statusCode, resultMsg = verifyBasicParameters(appId, timestamp, token, sign)
		if statusCode != api.RESPONSE_STATUS_OK {
			uniledgerlog.Error(resultMsg)
			responseWithStatusCode(ctx, statusCode, resultMsg)
			return
		}
	}

	// 4. verify all the request parameters are allowed
	ctx.Request.ParseForm()
	parameters := ctx.Request.Form
	if AUTH_VERIFY_ALLOWED_PARAMETERS {
		statusCode, resultMsg = verifyAllowRequestParameters(parameters)
		if statusCode != api.RESPONSE_STATUS_OK {
			uniledgerlog.Error(resultMsg)
			responseWithStatusCode(ctx, statusCode, resultMsg)
			return
		}
	}

	// 5. verify the sign
	if AUTH_VERIFY_SIGN {
		statusCode, resultMsg = verifySign(parameters, sign)
		if statusCode != api.RESPONSE_STATUS_OK {
			uniledgerlog.Error(resultMsg)
			responseWithStatusCode(ctx, statusCode, resultMsg)
			return
		}
	}
	// 6. verify the token exist
	if AUTH_VERIFY_TOKEN {
		tokenKey := appId + "_" + token
		if verifyTheToken(token, appId, api.TOKEN_MAP[tokenKey]) {
			_ = api.UpdateToken(tokenKey)
		} else {
			resultMsg := fmt.Sprintf("%s token(%s)错误!", "Filter[APIBasicFilter]", sign)
			uniledgerlog.Error(resultMsg)
			responseWithStatusCode(ctx, api.RESPONSE_STATUS_INVALID_TOKEN, resultMsg)
			return
		}
	}
}

func verifyTheToken(token string, appId string, accessKey string) bool {
	//tokens[token_appId] = accessKey(server store)
	// token = appId + generateTimestamp + randomString+ endDate
	return true
}

func APIRateLimitFilter(ctx *context.Context) {
	appId := ctx.Input.Query(api.REQUEST_FIELD_AUTH_APPID)
	token := ctx.Input.Query(api.REQUEST_FIELD_AUTH_TOKEN)
	tokenKey := appId + "_" + token
	exist := api.ExistKey(tokenKey)
	if exist {
		_ = api.UpdateToken(tokenKey)
		ok, resultMsg := api.RateLimit(tokenKey)
		if ok {
			uniledgerlog.Debug(resultMsg)
		} else {
			uniledgerlog.Error(resultMsg)
			responseWithStatusCode(ctx, api.RESPONSE_STATUS_APP_REQUESTS_OUT_OF_RATE_LIMIT, resultMsg)
			return
		}
	} else {
		isValid := verifyTheToken(token, appId, api.TOKEN_MAP[tokenKey])
		if isValid {
			api.UpdateToken(tokenKey)
		}
	}
}
