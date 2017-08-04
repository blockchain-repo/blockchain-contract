package api

import (
	"testing"
	//"time"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/core/db/redis"
)

func Test_ExampleNewClient(t *testing.T) {

	conn, err := redis.GetConn()
	if err != nil {
		uniledgerlog.Error(err)
		return
	}
	defer conn.Close()
	val1, err := redis.GetString(conn, "key")
	if err != nil {
		uniledgerlog.Error(err)
		return
	}
	uniledgerlog.Info(val1)

	exist_token := redis.ExistKey(conn, "Token")

	uniledgerlog.Info("bb token is", exist_token)

	_, err = redis.SetVal(conn, "Token", "123")
	redis.SetExpire(conn, "Token", 10)
	if err != nil {
		uniledgerlog.Error(err)
		return
	}
	token, err := redis.GetString(conn, "Token")

	uniledgerlog.Info("token is", token)
	//val, err = redis.String(conn.Do("GET", "key"))
	//if err != nil {
	//	uniledgerlog.Error(err)
	//	return
	//}
	//uniledgerlog.Info(val)
	//val, err = redis.String(conn.Do("GET", "Token"))
	//uniledgerlog.Info(val)
	//
	//val, err = conn.Do("SET", "Token", time.Now().String())
	//conn.Do("EXPIRE", "TOKEN", 1)
	//uniledgerlog.Info("set", val)
	//val, err = redis.String(conn.Do("GET", "Token"))
	//uniledgerlog.Info(val)
}
