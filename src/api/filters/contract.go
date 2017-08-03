package filters

import (
	"crypto"
	"encoding/hex"
	"github.com/astaxie/beego/context"
	"go/token"
	"strconv"
	"time"
	http_status "unicontract/src/api"
	"unicontract/src/common"
	"unicontract/src/common/uniledgerlog"
)

func responseWithStatusCode(ctx *context.Context, status int, output string) {
	//ctx.Output.SetStatus(status) // invalid setStatus
	ctx.ResponseWriter.WriteHeader(status)
	ctx.ResponseWriter.Write([]byte(output))
}

func checkExistAppUser(store_secure_id string) (exist bool) {
	// get from db , exist records return token ,generate
	// now temp return generate token!
	return true
}

func generateToken() string {
	token := "ffffffffffffffffffffffffffffffffffffffffffffffff"
	return token
}

//var ContentTypes = []string{"application/json", "application/x-protobuf"}

// get token use app_id and app_key from redis
func APIAuthorizationFilter(ctx *context.Context) {
	app_id := ctx.Input.Query("app_id")
	app_key := ctx.Input.Query("app_key")
	// app_id app_key encode -> store_secure_id
	//
	h := crypto.SHA3_512.New()
	h.Write([]byte(app_id + "_" + app_key))
	x := h.Sum(nil)
	y := make([]byte, 32)
	hex.Encode(y, x)

	store_secure_id := string(y)
	exist := checkExistAppUser(store_secure_id)
	if exist {
		token := generateToken()
		uniledgerlog.Info("APIAuthorizationFilter exist user!", token)
		responseWithStatusCode(ctx, http_status.HTTP_STATUS_CODE_OK, token)
	} else {
		uniledgerlog.Error("APIAuthorizationFilter not exist user!")
		responseWithStatusCode(ctx, http_status.HTTP_STATUS_CODE_Forbidden, "not exist user!")
		return
	}

}

// 请求参数验证
func APIContentTypeFilter(ctx *context.Context) {
	contentType := ctx.Input.Header("Content-Type")

	if contentType == "" {
		result := make(map[string]interface{})
		result["msg"] = "error Headers"
		result["status"] = http_status.HTTP_STATUS_CODE_BadRequest
		uniledgerlog.Error("APIContentTypeFilter contentType or requestDataType is empty!")
		responseWithStatusCode(ctx, http_status.HTTP_STATUS_CODE_BadRequest, common.StructSerialize(result))
		return

	} else if contentType == "application/json" || contentType == "application/x-protobuf" {
		//uniledgerlog.Debug("RequestDataType is json!")
	} else {
		uniledgerlog.Error("APIContentTypeFilter contentType error!")
		responseWithStatusCode(ctx, http_status.HTTP_STATUS_CODE_BadRequest, "APIContentTypeFilter contentType error!")
		return
	}
}

//签名身份验证
func APIAuthFilter(ctx *context.Context) {
	// 5s
	// test 60*60*10 10hours
	timeout := int64(60 * 60 * 10)
	t := time.Now()
	nanos := t.UnixNano()
	//ms len=13
	current_unix_timestamp := nanos / 1000000
	timestamp := ctx.Input.Query("timestamp")
	if len(timestamp) != 13 {
		uniledgerlog.Error("APIAuthFilter timestamp error!")
		responseWithStatusCode(ctx, http_status.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter timestamp error!")
		return
	}

	token := ctx.Input.Query("token")
	if len(token) != 48 {
		uniledgerlog.Error("APIAuthFilter token error!", token)
		responseWithStatusCode(ctx, http_status.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter token error!"+string(token))
		return
	}

	timestamp_int64, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		uniledgerlog.Error("APIAuthFilter error!", err)
		return
	}
	timecost := (current_unix_timestamp - timestamp_int64) / 1000
	uniledgerlog.Info("time info", current_unix_timestamp, timestamp_int64, timecost)
	if timecost < 0 || timecost > timeout {
		uniledgerlog.Error("APIAuthFilter timestamp invalid!", timecost)
		responseWithStatusCode(ctx, http_status.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter timestamp invalid!"+string(timecost)+"s")
		return
	}

	// sign verify
	sign := ctx.Input.Query("sign")
	if len(sign) == 0 {
		uniledgerlog.Error("APIAuthFilter token error!", token)
		responseWithStatusCode(ctx, http_status.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter sign error!"+string(token))
		return
	}

}
