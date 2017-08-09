package api

import (
	"testing"
	//"time"
	"fmt"
	"time"
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

func sum(a, b int) {
	//defer func() { result += a }() //被延时执行的匿名函数甚至可以修改函数返回给调用者的返回值
	//defer timecost()("sum is " + string(result)) // note:不要忘记defer语句后的圆括号，否则本该在进入时执行的操作会在退出时执行，而本该在退出时执行的，永远不会执行
	start := time.Now()
	result := 0
	fmt.Printf("enter... %v\n", result)
	result += a
	result += b
	fmt.Printf("enter... %v\n", result)
	if a <= 4 {
		defer timecost(start, " GET|200 https://www.uni-ledger.com/v1/contract-app/contracts?token=ADFADAEGLEQODGSL"+
			"&timestamp=1501225200000 192.168.0.1 33.33.12.12")() // note:不要忘
		return
	}
	defer timecost(start, " GET|200 https://www.uni-ledger.com/v1/contract-app/contracts?token=ADFADAEGLEQODGSL"+
		"&timestamp=1501225200000 192.168.0.1 33.33.12.12")() // note:不要忘记defer语句后的圆括号，否则本该在进入时执行的操作会在退出时执行，而本该在退出时执行的，永远不会执行
	//return result
}

//func timecost() func(resultLog string) {
//	fmt.Printf("enter... %v\n", 11)
//	start := time.Now()
//	return func(resultLog string) {
//		uniledgerlog.Info("%s %s", resultLog, time.Since(start))
//	}
//}

func timecost(start time.Time, resultLog string) func() {
	fmt.Printf("enter... %v\n", 11)
	return func() {
		uniledgerlog.Info("%s %d", resultLog, time.Since(start))
		//uniledgerlog.Info("%s %d", resultLog, time.Since(start).Nanoseconds()/1000000)
	}
}

func TestSDFF(t *testing.T) {
	sum(3, 6)
	//count := sum(3, 6)
	//_ = count
	//fmt.Printf("%v\n", count)
}
