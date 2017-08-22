package api

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"golang.org/x/crypto/sha3"
	"hash"
	"sync"
	"time"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/db/redis"
)

var lock sync.RWMutex

// ratelimit token
var (
	// 45s temp store for apply the token
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
	auth_verify := beego.AppConfig.DefaultBool("auth_verify", true)
	auth_verify_token := beego.AppConfig.DefaultBool("auth_verify_token", true)
	auth_verify_rate_limit := beego.AppConfig.DefaultBool("auth_verify_rate_limit", true)
	if !auth_verify || !auth_verify_token && !auth_verify_rate_limit {
		uniledgerlog.Warn("no need redis install")
		return
	}
	uniledgerlog.Warn("need redis install, if not you can set \nthe api_auth = false in app.conf to skip the api verify!")
	uniledgerlog.Debug("init config %v,%v,%v", token_timeout, rate_limit_duration, rate_limit_count)
	token_timeout = beego.AppConfig.DefaultInt("api.token_timeout", token_timeout)
	rate_limit_key_suffix = beego.AppConfig.DefaultString("api.rate_limit_key_suffix", rate_limit_key_suffix)
	rate_limit_duration = beego.AppConfig.DefaultInt("api.rate_limit_duration", rate_limit_duration)
	rate_limit_count = beego.AppConfig.DefaultInt("api.rate_limit_count", rate_limit_count)
	uniledgerlog.Debug("load config %v,%v,%v", token_timeout, rate_limit_duration, rate_limit_count)

}

// TimeCost need cost_start := time.Now() and insert it before return
func TimeCost(start time.Time, ctx *context.Context, responseCode int32, msg string) func() {
	return func() {
		ctx.Request.ParseForm()
		parameters := ctx.Request.Form
		parameters.Del("token")
		requestParameters := parameters.Encode()
		result := fmt.Sprintf("API_INFO[method=%s code=%d from=%s to=%s cost=%.3fms parameters=(%s)] msg=%s", ctx.Request.Method, responseCode, ctx.Request.RemoteAddr, ctx.Request.Host,
			float32(time.Since(start).Nanoseconds())/1000000, requestParameters, msg)
		uniledgerlog.Info("%s", result)
	}
}

func ExistKey(token string) bool {
	return redis.ExistKey(token)
}

func UpdateToken(token string) bool {
	redis.SetValWithExpire(token, time.Now().String(), token_timeout)
	return true
}

func RateLimit(token string) (ok bool, msg string) {
	lock.Lock()
	defer lock.Unlock()

	token_rate_key := token + rate_limit_key_suffix
	times, _ := redis.GetInt(token_rate_key)
	ttl, err := redis.TTL(token_rate_key)
	if ttl <= -1 {
		redis.SetValWithExpire(token_rate_key, "1", rate_limit_duration)
		return true, "达到重置时间,重置 TTL"
	}
	if err != nil {
		uniledgerlog.Error(err)
	}

	if times > rate_limit_count {
		rate_limit_log := fmt.Sprintf("%s,[%ds,%d], 剩余重置时间 %ds", "用户访问频率超限", rate_limit_duration, rate_limit_count, ttl)
		uniledgerlog.Warn(rate_limit_log)
		return false, rate_limit_log
	} else {
		redis.Incr(token_rate_key)
	}

	rate_limit_log := fmt.Sprintf("%s,[%ds,%d], 剩余 %d 次, 剩余重置时间%ds", "用户访问频率", rate_limit_duration, rate_limit_count, rate_limit_count-times, ttl)
	return true, rate_limit_log
}

func hashData(val string) string {
	var hash hash.Hash
	var x string
	hash = sha3.New256()
	if hash != nil {
		hash.Write([]byte(val))
		x = hex.EncodeToString(hash.Sum(nil))
	}
	return x
}

func md5Encode(val string) string {
	var hash hash.Hash
	var x string
	hash = crypto.MD5.New()
	if hash != nil {
		hash.Write([]byte(val))
		x = hex.EncodeToString(hash.Sum(nil))
	}
	return x
}
