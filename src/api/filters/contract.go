package filters

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strconv"
	"time"
	api "unicontract/src/api"
	"unicontract/src/common"
	"unicontract/src/common/uniledgerlog"
)

var (
	API_TIMEOUT       = int64(60)
	API_TOKEN_LEN     = 44
	API_TIMESTAMP_LEN = 13
)

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func init() {
	API_TIMEOUT = beego.AppConfig.DefaultInt64("api_timeout", API_TIMEOUT)
}

func responseWithStatusCode(ctx *context.Context, status int, output interface{}) {
	result := response{Code: status, Msg: "", Data: output}
	resultByte, err := json.Marshal(result)
	if err != nil {
		uniledgerlog.Error("responseWithStatusCode", err.Error())
	}
	ctx.ResponseWriter.Write(resultByte)
	return
}

//var ContentTypes = []string{"application/json", "application/x-protobuf"}

// todo filter the error Content-Type
// APIContentTypeFilter step 1
func APIContentTypeFilter(ctx *context.Context) {
	cost_start := time.Now()

	contentType := ctx.Input.Header("Content-Type")

	if contentType == "" {
		result := make(map[string]interface{})
		result["msg"] = "error Headers"
		result["status"] = api.RESPONSE_STATUS_BadRequest
		uniledgerlog.Error("APIContentTypeFilter contentType is empty!")
		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_OK, common.StructSerialize(result))
		defer api.TimeCost(cost_start, ctx, api.RESPONSE_STATUS_BadRequest)()
		return

	} else if contentType == "application/json" || contentType == "application/x-protobuf" {
		//uniledgerlog.Debug("RequestDataType is json!")
	} else {
		uniledgerlog.Error("APIContentTypeFilter contentType error!")
		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_OK, "APIContentTypeFilter contentType error!")
		defer api.TimeCost(cost_start, ctx, api.RESPONSE_STATUS_BadRequest)()
		return
	}

}

// todo get the token with app_id and app_key; if not exist, response 111000403.
// APIAuthorizationFilter step 2 授权,请求token
// accesskey need store in redis for next step verify accessKey
func APIAuthorizationFilter(ctx *context.Context) {
	cost_start := time.Now()
	app_id := ctx.Input.Query("app_id")
	app_key := ctx.Input.Query("app_key")

	/******************* test data, each has the static  *******************/
	if app_id == "" {
		app_id = "0123456789"
	}
	if app_key == "" {
		app_key = "uni-ledger.com"
	}
	/******************* test data *******************/

	exist := api.CheckExistAppUser(api.GenerateAccessKey(app_id, app_key))

	if exist {
		access_key := api.GenerateAccessKey(app_id, app_key)
		uniledgerlog.Info("APIAuthorizationFilter exist user!", access_key)
		//todo
		_ = api.StoreAccessKey(app_id, access_key)
		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_OK, access_key)
	} else {
		uniledgerlog.Error("APIAuthorizationFilter not exist user!")
		// not exist, generateToken and put redis
		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_OK, "not exist user!")
		defer api.TimeCost(cost_start, ctx, api.RESPONSE_STATUS_Forbidden)()
		return
	}
}

func APIGetTokenFilter(ctx *context.Context) {
	cost_start := time.Now()
	access_key := ctx.Input.Query("accessKey")
	app_id := ctx.Input.Query("app_id")
	if app_id == "" {
		app_id = "0123456789"
	}
	if access_key != "" {
		token, ok := api.GetToken(app_id, access_key)
		if ok {
			uniledgerlog.Info("APIGetTokenFilter exist accessKey!", token)
			responseWithStatusCode(ctx, api.RESPONSE_STATUS_OK, token)

		} else {
			uniledgerlog.Info("APIGetTokenFilter not exist accessKey!", token)
			responseWithStatusCode(ctx, api.RESPONSE_STATUS_BadRequest, token)
		}
	} else {
		uniledgerlog.Error("APIGetTokenFilter accessKey error!")
		// not exist, generateToken and put redis
		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_OK, "accessKey error!")
		defer api.TimeCost(cost_start, ctx, api.RESPONSE_STATUS_Forbidden)()
		return
	}

}

// APIAuthorizationFilter step 3 API 参数及权限认证
func APIAuthFilter(ctx *context.Context) {
	// 1. verify the parameters is miss, timestamp, token,sign
	token := ctx.Input.Query("token")
	timestamp := ctx.Input.Query("timestamp")
	sign := ctx.Input.Query("sign")

	if token == "" || timestamp == "" || sign == "" {
		uniledgerlog.Error("APIAuthFilter parameters miss!")
		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter parameters miss!")
		return
	}
	if len(token) != API_TOKEN_LEN {
		uniledgerlog.Error("APIAuthFilter parameters token error!")
		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter parameters token error!")
		return
	}
	if len(timestamp) != API_TIMESTAMP_LEN {
		uniledgerlog.Error("APIAuthFilter parameters timestamp error!")
		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter parameters timestamp error!")
		return
	}

	//if len(sign) != API_SIGN_LEN {
	//	uniledgerlog.Error("APIAuthFilter parameters sign error!")
	//	responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter parameters sign error!")
	//	return
	//}

	//if len(token) != API_TOKEN_LEN || len(timestamp) != API_TIMESTAMP_LEN || len(sign) != API_SIGN_LEN {
	//	uniledgerlog.Error("APIAuthFilter parameters error!")
	//	responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter parameters error!")
	//	return
	//}

	// 2. verify the timestamp
	nanos := time.Now().UnixNano()
	//ms len=13
	current_unix_timestamp := nanos / 1000000

	timestamp_int64, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		uniledgerlog.Error("APIAuthFilter error!", err)
		return
	}
	cost := (current_unix_timestamp - timestamp_int64) / 1000
	uniledgerlog.Debug("time info", current_unix_timestamp, timestamp_int64, cost)
	if cost < 0 || cost > API_TIMEOUT {
		uniledgerlog.Error("APIAuthFilter timestamp invalid!", cost)
		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter timestamp invalid!"+strconv.FormatInt(cost, 10)+"s")
		return
	}

	// 3. verify the token
	if !api.ExistKey(token) {
		uniledgerlog.Error("APIAuthFilter token not exist!")
		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter token not exist! "+strconv.FormatInt(cost, 10)+" s")
		return
	}
	// 4. verify the sign
	if !api.VerifySign(token, timestamp, sign) {
		uniledgerlog.Error("APIAuthFilter sign error!", cost)
		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter sign error!"+strconv.FormatInt(cost, 10)+"s")
		return
	}
}

// APIRateLimitFilter step 4 API Ratelimit 验证
func APIRateLimitFilter(ctx *context.Context) {
	//cost_start := time.Now()
	token := ctx.Input.Query("token")
	if !api.RateLimit(token) {
		responseWithStatusCode(ctx, api.HTTP_STATUS_CODE_TOO_MANY_REQUESTS, "APIRateLimitFilter error!"+string(token))
		//defer api.TimeCost(cost_start, ctx, api.HTTP_STATUS_CODE_TOO_MANY_REQUESTS)()
		return
	}
	// last filter, will reset the token valid time
	api.UpdateToken(token)
}
