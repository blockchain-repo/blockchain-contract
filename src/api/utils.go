package api

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/btcsuite/btcutil/base58"
	"hash"
	"time"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/db/redis"
)

// ratelimit token
var (
	// 45s temp store for apply the token
	access_key_timeout = 45
	access_key_blur    = "uniledger"
	// 1 hour
	token_timeout = 60 * 60
	// {token}_rate
	rate_limit_key_suffix = "_rate"
	// 100 tps
	rate_limit_duration = 1
	rate_limit_count    = 100
)

func init() {
	//limit the request timeout s
	//API_TIMEOUT = beego.AppConfig.DefaultInt64("api_timeout", API_TIMEOUT)
	// if true, add the api filter
	api_auth := beego.AppConfig.DefaultBool("api_auth", true)
	if !api_auth {
		uniledgerlog.Warn("no need redis install")
		return
	}
	uniledgerlog.Warn("need redis install, if not you can set \nthe api_auth = false in app.conf to skip the api verify!")
	uniledgerlog.Debug("init config %v,%v,%v", access_key_timeout, access_key_blur, token_timeout)
	access_key_timeout = beego.AppConfig.DefaultInt("api.access_key_timeout", access_key_timeout)
	access_key_blur = beego.AppConfig.DefaultString("api.access_key_blur", access_key_blur)
	token_timeout = beego.AppConfig.DefaultInt("api.token_timeout", token_timeout)
	rate_limit_key_suffix = beego.AppConfig.DefaultString("api.rate_limit_key_suffix", rate_limit_key_suffix)
	rate_limit_duration = beego.AppConfig.DefaultInt("api.rate_limit_duration", rate_limit_duration)
	rate_limit_count = beego.AppConfig.DefaultInt("api.rate_limit_count", rate_limit_count)
	uniledgerlog.Debug("load config %v,%v,%v", access_key_timeout, access_key_blur, token_timeout)

}

// TimeCost need cost_start := time.Now() and insert it before return
func TimeCost(start time.Time, ctx *context.Context, responseCode int32, msg string) func() {
	return func() {
		result := fmt.Sprintf("API_INFO[method=%s code=%d from=%s to=%s cost=%.3fms msg=%s]", ctx.Request.Method, responseCode, ctx.Request.RemoteAddr, ctx.Request.Host,
			float32(time.Since(start).Nanoseconds())/1000000, msg)
		uniledgerlog.Info("%s", result)
	}
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
	accessKeyTemp := app_id + access_key_blur + access_key
	_, err := redis.SetVal(app_id, accessKeyTemp)
	if err != nil {
		uniledgerlog.Error(err.Error())
		return false
	}
	_, err = redis.SetExpire(app_id, access_key_timeout)
	if err != nil {
		uniledgerlog.Error(err.Error())
		return false
	}
	return true
}

func ExistKey(token string) bool {
	return redis.ExistKey(token)
}

func VerifySign(token string, timestamp string, sign string) bool {
	//return sign == md5Encode(token+"_"+timestamp)
	return true
}

// 1. 根据用户的app_id 和 app_key，生成access_key; 2. 通过access_key获取token
func GetToken(app_id string, access_key string) (string, bool) {
	// access_key=token, judge the expire
	// token=counter judge the rate limit, expired will reproduce

	existAccessKey := ExistKey(app_id)
	if !existAccessKey {
		// maybe invalid access_key
		return "", false
	}
	//md5Encode
	// default_app_id-unix_timestamp
	token := generateToken(access_key)
	redis.SetVal(token, time.Nanosecond.String())
	redis.SetExpire(token, token_timeout)

	// set the rate limit token
	token_rate_key := token + rate_limit_key_suffix
	redis.SetVal(token_rate_key, "0")
	redis.SetExpire(token_rate_key, rate_limit_duration)
	uniledgerlog.Debug("not exist token and generate the token!", token)
	// when generate the token success, must expire the access_key
	redis.SetExpire(app_id, 0)
	return token, true
}

func UpdateToken(token string) bool {
	redis.SetVal(token, time.Nanosecond.String())
	redis.SetExpire(token, token_timeout)
	return true
}

func RateLimit(token string) (ok bool) {
	token_rate_key := token + rate_limit_key_suffix
	if !redis.ExistKey(token_rate_key) {
		uniledgerlog.Warn("token_rate is expired, will reset!")
	}

	times, _ := redis.GetInt(token_rate_key)
	if times == 0 {
		redis.Incr(token_rate_key)
		redis.SetExpire(token_rate_key, rate_limit_duration)
	} else {
		if times < rate_limit_count {
			redis.Incr(token_rate_key)
		} else {
			ttl, err := redis.TTL(token_rate_key)
			if err != nil {
				uniledgerlog.Error(err)
			}
			rate_limit_log := fmt.Sprintf("%s,[%ds,%d], 剩余重置时间 %ds", "用户访问频率超限", rate_limit_duration, rate_limit_count, ttl)
			uniledgerlog.Warn(rate_limit_log)
			return false
		}
	}
	ttl, _ := redis.TTL(token_rate_key)
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
