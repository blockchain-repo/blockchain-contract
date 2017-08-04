package filters

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/btcsuite/btcutil/base58"
	"hash"
	//"strconv"
	//"time"
	http_status "unicontract/src/api"
	"unicontract/src/common"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/db/redis"
)

func responseWithStatusCode(ctx *context.Context, status int, output string) {
	//ctx.Output.SetStatus(status) // invalid setStatus
	ctx.ResponseWriter.WriteHeader(status)
	ctx.ResponseWriter.Write([]byte(output))
}

const (
	// 30 minutes
	token_timeout    = 60 * 30
	rate_limit_time  = 2
	rate_limit_count = 100
)

func checkExistAppUser(store_secure_id string) (exist bool) {
	// get from db , exist records return token ,generate
	// now temp return generate token!
	return true
}

func hashData(val string) string {
	var hash hash.Hash
	var x string = ""
	hash = crypto.SHA3_256.New()
	if hash != nil {
		hash.Write([]byte(val))
		x = hex.EncodeToString(hash.Sum(nil))
	}
	return x
}

func md5Encode(val string) string {
	var hash hash.Hash
	var x string = ""
	hash = crypto.MD5.New()
	if hash != nil {
		hash.Write([]byte(val))
		x = hex.EncodeToString(hash.Sum(nil))
	}
	return x
}

func generateToken(str string) string {
	md5Str := md5Encode(str)
	return base58.Encode([]byte(md5Str))
}

func generateSecureStoreKey(app_id, app_key string) string {
	return hashData(app_id + "-" + app_key)
}

func getToken() string {
	//token := "ffffffffffffffffffffffffffffffffffffffffffffffff"
	default_app_id := "0123456789"
	default_app_key := "uni-ledger.com"
	secure_token_key := generateSecureStoreKey(default_app_id, default_app_key)
	conn, _ := redis.GetConn()
	exist := redis.ExistKey(conn, secure_token_key)
	if exist {
		token, _ := redis.GetString(conn, secure_token_key)
		redis.SetExpire(conn, secure_token_key, token_timeout)
		uniledgerlog.Info("APIAuthorizationFilter get token!", token)
		return token
	} else {
		// default_app_id-unix_timestamp
		token := generateToken(generateSecureStoreKey(default_app_id, common.GenTimestamp()))
		redis.SetVal(conn, secure_token_key, token)
		redis.SetExpire(conn, secure_token_key, token_timeout)

		// token aid/userid
		redis.SetVal(conn, token, default_app_id)
		redis.SetExpire(conn, token, rate_limit_time)

		uniledgerlog.Info("APIAuthorizationFilter generateToken token!", token)
		return token
	}
}

//var ContentTypes = []string{"application/json", "application/x-protobuf"}

// get token use app_id and app_key from redis
func APIAuthorizationFilter(ctx *context.Context) {
	app_id := ctx.Input.Query("app_id")
	app_key := ctx.Input.Query("app_key")
	//
	exist := checkExistAppUser(generateSecureStoreKey(app_id, app_key))
	if exist {
		token := getToken()
		uniledgerlog.Info("APIAuthorizationFilter exist user!", token)
		responseWithStatusCode(ctx, http_status.HTTP_STATUS_CODE_OK, token)
	} else {
		uniledgerlog.Error("APIAuthorizationFilter not exist user!")
		// not exist, generateToken and put redis
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

func rateLimit(token string) (rateOK bool) {
	//ctx.Request.Host
	conn, err := redis.GetConn()
	if err != nil {
		uniledgerlog.Error(err)
	}
	if !redis.ExistKey(conn, token) {
		uniledgerlog.Error("token is not exist, will generate!")
		//return false
	}
	times, err := redis.GetInt(conn, token)
	if times == 0 {
		redis.Incr(conn, token)
		redis.SetExpire(conn, token, rate_limit_time)
	} else {
		if times < rate_limit_count {
			redis.Incr(conn, token)
		} else {
			ttl, err := redis.TTL(conn, token)
			if err != nil {
				uniledgerlog.Error(err)
			}
			fmt.Sprintf("%s,[%ds,%d],left %d", "用户访问频率超限", rate_limit_time, rate_limit_count, ttl)
			uniledgerlog.Error(fmt.Sprintf("%s,[%ds,%d],left %d", "用户访问频率超限", rate_limit_time, rate_limit_count, ttl))
			return false
		}
	}
	ttl, err := redis.TTL(conn, token)
	result := fmt.Sprintf("%s,[%ds,%d],剩余 %d, 重置时间%d", "用户访问频率", rate_limit_time, rate_limit_count, rate_limit_count-times, ttl)
	uniledgerlog.Info(result)
	return true
}

//签名身份验证
func APIAuthFilter(ctx *context.Context) {
	// 5s
	// test 60*60*10 10hours
	//timeout := int64(60 * 60 * 10)
	//t := time.Now()
	//nanos := t.UnixNano()
	////ms len=13
	//current_unix_timestamp := nanos / 1000000
	//timestamp := ctx.Input.Query("timestamp")
	//if len(timestamp) != 13 {
	//	uniledgerlog.Error("APIAuthFilter timestamp error!")
	//	responseWithStatusCode(ctx, http_status.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter timestamp error!")
	//	return
	//}

	token := ctx.Input.Query("token")
	if len(token) != 44 {
		uniledgerlog.Error("APIAuthFilter token error!", token)
		responseWithStatusCode(ctx, http_status.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter token error!"+string(token))
		return
	}
	if !rateLimit(token) {
		responseWithStatusCode(ctx, http_status.HTTP_STATUS_CODE_Forbidden, "APIAuthFilter token error!"+string(token))
		return
	}

	//timestamp_int64, err := strconv.ParseInt(timestamp, 10, 64)
	//if err != nil {
	//	uniledgerlog.Error("APIAuthFilter error!", err)
	//	return
	//}
	//timecost := (current_unix_timestamp - timestamp_int64) / 1000
	//uniledgerlog.Info("time info", current_unix_timestamp, timestamp_int64, timecost)
	//if timecost < 0 || timecost > timeout {
	//	uniledgerlog.Error("APIAuthFilter timestamp invalid!", timecost)
	//	responseWithStatusCode(ctx, http_status.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter timestamp invalid!"+string(timecost)+"s")
	//	return
	//}

	// sign verify
	//sign := ctx.Input.Query("sign")
	//if len(sign) == 0 {
	//	uniledgerlog.Error("APIAuthFilter token error!", token)
	//	responseWithStatusCode(ctx, http_status.HTTP_STATUS_CODE_BadRequest, "APIAuthFilter sign error!"+string(token))
	//	return
	//}

}
