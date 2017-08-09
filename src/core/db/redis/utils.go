package redis

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"os"
	"time"
	"unicontract/src/common/uniledgerlog"
)

var (
	redis_client  *redis.Pool
	redis_address string
	redis_db      int
)

func init() {
	uniledgerlog.Info("redis pool init start")
	redis_host := beego.AppConfig.String("redis.host")
	redis_port := beego.AppConfig.String("redis.port")
	redis_address = redis_host + ":" + redis_port
	redis_db, _ = beego.AppConfig.Int("redis.db")
	redis_password := beego.AppConfig.String("redis.password")
	// create redis pool
	redis_client = &redis.Pool{
		MaxIdle:     beego.AppConfig.DefaultInt("redis.maxidle", 2),
		MaxActive:   beego.AppConfig.DefaultInt("redis.maxactive", 4),
		IdleTimeout: 30 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redis_address, redis.DialPassword(redis_password))
			if err != nil {
				uniledgerlog.Error("redis 初始化错误，请检测服务器是否安装了redis及其配置参数\n"+
					"api的token及ratelimit验证使用了redis，可以通过设置app.conf中的api_auth = false来禁止这些功能以跳过此错误", err)
				os.Exit(0)
				return nil, err
			}
			// choose db
			_, err = c.Do("SELECT", redis_db)
			if err != nil {
				uniledgerlog.Error(err)
				return nil, err
			}
			return c, nil
		},
	}
	uniledgerlog.Info("redis pool init end")
	config_info := fmt.Sprintf("host=%s, port=%s, address=%s,db=%d,password=%s,maxidle=%d,maxactive=%d", redis_host, redis_port, redis_address, redis_db, redis_password,
		beego.AppConfig.DefaultInt("redis.maxidle", 1), beego.AppConfig.DefaultInt("redis.maxactive", 10))
	uniledgerlog.Debug("redis pool init end, config info: \n%s", config_info)
}

func closePool() bool {
	return redis_client.Close() != nil
}

func GetConn() (conn redis.Conn) {
	return redis_client.Get()
}

// SetVal set the key use val
func SetVal(key, val string) (interface{}, error) {
	conn := GetConn()
	defer conn.Close()
	return conn.Do("SET", key, val)
}

// SetExpire
func SetExpire(key string, expire int) (interface{}, error) {
	conn := GetConn()
	defer conn.Close()
	return conn.Do("EXPIRE", key, expire)
}

// GetVal
func GetVal(key string) (interface{}, error) {
	conn := GetConn()
	defer conn.Close()
	return conn.Do("GET", key)
}

func GetString(key string) (string, error) {
	conn := GetConn()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))

}

func GetBool(key string) (bool, error) {
	conn := GetConn()
	defer conn.Close()
	return redis.Bool(conn.Do("GET", key))

}

func GetInt64(key string) (int, error) {
	conn := GetConn()
	defer conn.Close()
	return redis.Int(conn.Do("GET", key))

}

func GetInt(key string) (int, error) {
	conn := GetConn()
	defer conn.Close()
	return redis.Int(conn.Do("GET", key))

}

func Incr(key string) (interface{}, error) {
	conn := GetConn()
	defer conn.Close()
	return conn.Do("INCR", key)
}

func TTL(key string) (int, error) {
	conn := GetConn()
	defer conn.Close()
	return redis.Int(conn.Do("TTL", key))
}

// ExistKey
func ExistKey(key string) bool {
	conn := GetConn()
	defer conn.Close()
	exist_key, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		uniledgerlog.Error(err)
	}
	return exist_key
}
