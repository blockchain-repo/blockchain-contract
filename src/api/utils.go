package api

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/btcsuite/btcutil/base58"
	"hash"
	"time"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/db/redis"
)

// TimeCose need cost_start := time.Now() and insert it before return
func TimeCost(start time.Time, ctx *context.Context, responseCode int32) func() {
	return func() {
		result := fmt.Sprintf("API_INFO[code=%s method=%d from=%s to=%s cost=%.3fms]", ctx.Request.Method, responseCode, ctx.Request.RemoteAddr, ctx.Request.Host,
			float32(time.Since(start).Nanoseconds())/1000000)
		uniledgerlog.Debug("%s", result)
	}
}

// todo test
func getAppConfig() (app_id string, app_key string) {
	app_id = "0123456789"
	app_key = "uni-ledger.com"
	return
}

func GetAccessKey(app_id string, app_key string) string {
	return GenerateAccessKey(app_id, app_key)
}

//todo temp deal for accessKey
// GenerateAccessKey 用于获取token
func GenerateAccessKey(app_id string, app_key string) string {
	return hashData(fmt.Sprintf("%s-%d-%s", app_id, time.Now().UnixNano(), app_key))
}

func CheckExistAppUser(store_secure_id string) (exist bool) {
	// get from db , exist records return token ,generate
	// now temp return generate token!
	return true
}

// todo need add salt and other deals
func generateToken(str string) string {
	md5Str := md5Encode(str)
	return base58.Encode([]byte(md5Str))
}

// accesskey temp store app_id=accessKeyTemp, app_id maybe replace with user_id
func StoreAccessKey(app_id string, access_key string) bool {
	conn, _ := redis.GetConn()
	accessKeyTemp := app_id + access_key_blur + access_key
	_, err := redis.SetVal(conn, app_id, accessKeyTemp)
	if err != nil {
		uniledgerlog.Error(err.Error())
		return false
	}
	_, err = redis.SetExpire(conn, app_id, access_key_timeout)
	if err != nil {
		uniledgerlog.Error(err.Error())
		return false
	}
	return true
}

func ExistKey(token string) bool {
	conn, _ := redis.GetConn()
	return redis.ExistKey(conn, token)

}

func VerifySign(token string, timestamp string, sign string) bool {
	//return sign == md5Encode(token+"_"+timestamp)
	return true
}

// 1. 根据用户的app_id 和 app_key，生成access_key; 2. 通过access_key获取token
func GetToken(app_id string, access_key string) (string, bool) {
	// access_key=token, judge the expire
	// token=counter judge the rate limit, expired will reproduce

	conn, _ := redis.GetConn()
	existAccessKey := ExistKey(app_id)
	if !existAccessKey {
		// maybe invalid access_key
		return "", false
	}
	//md5Encode
	// default_app_id-unix_timestamp
	token := generateToken(access_key)
	redis.SetVal(conn, token, time.Nanosecond.String())
	redis.SetExpire(conn, token, token_timeout)

	// set the rate limit token
	token_rate_key := token + rate_token_key
	redis.SetVal(conn, token_rate_key, "0")
	redis.SetExpire(conn, token_rate_key, rate_limit_duration)
	uniledgerlog.Debug("not exist token and generate the token!", token)
	// when generate the token success, must expire the access_key
	redis.SetExpire(conn, app_id, 0)
	return token, true

	//exist := redis.ExistKey(conn, access_key)
	//if exist {
	//	token, _ := redis.GetString(conn, access_key)
	//	// set the rate limit token
	//	redis.SetExpire(conn, access_key, token_timeout)
	//	uniledgerlog.Debug("exist token and return!", token)
	//	return token
	//} else {
	//	// default_app_id-unix_timestamp
	//	token := generateToken(access_key)
	//	redis.SetVal(conn, access_key, token)
	//	redis.SetExpire(conn, access_key, token_timeout)
	//
	//	// set the rate limit token
	//	redis.SetVal(conn, token, "0")
	//	redis.SetExpire(conn, token, rate_limit_duration)
	//	uniledgerlog.Debug("not exist token and generate the token!", token)
	//	return token
	//}
}

func UpdateToken(token string) bool {
	conn, _ := redis.GetConn()
	redis.SetVal(conn, token, time.Nanosecond.String())
	redis.SetExpire(conn, token, token_timeout)
	return true
}

func RateLimit(token string) (ok bool) {
	conn, err := redis.GetConn()
	if err != nil {
		uniledgerlog.Error(err)
		return false
	}

	token_rate_key := token + rate_token_key
	if !redis.ExistKey(conn, token_rate_key) {
		uniledgerlog.Error("token is expired, will reset!")
	}

	times, err := redis.GetInt(conn, token_rate_key)
	if times == 0 {
		redis.Incr(conn, token_rate_key)
		redis.SetExpire(conn, token_rate_key, rate_limit_duration)
	} else {
		if times < rate_limit_count {
			redis.Incr(conn, token_rate_key)
		} else {
			ttl, err := redis.TTL(conn, token_rate_key)
			if err != nil {
				uniledgerlog.Error(err)
			}
			rate_limit_log := fmt.Sprintf("%s,[%ds,%d], 剩余重置时间 %ds", "用户访问频率超限", rate_limit_duration, rate_limit_count, ttl)
			uniledgerlog.Debug(rate_limit_log)
			return false
		}
	}
	ttl, err := redis.TTL(conn, token_rate_key)
	result := fmt.Sprintf("%s,[%ds,%d], 剩余 %d 次, 剩余重置时间%ds", "用户访问频率", rate_limit_duration, rate_limit_count, rate_limit_count-times-1, ttl)
	uniledgerlog.Debug(result)
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
