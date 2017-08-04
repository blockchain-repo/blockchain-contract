package redis

import (
	"github.com/garyburd/redigo/redis"
	"os"
	"unicontract/src/common/uniledgerlog"
)

func GetConn() (conn redis.Conn, err error) {
	conn, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		uniledgerlog.Error(err)
		os.Exit(2)
	}
	return
}

// CloseConn close the redis Conn
func CloseConn(conn redis.Conn) bool {
	return conn.Close() != nil
}

// SetVal set the key use val
func SetVal(conn redis.Conn, key, val string) (interface{}, error) {
	return conn.Do("SET", key, val)
}

// SetExpire
func SetExpire(conn redis.Conn, key string, expire int64) (interface{}, error) {
	return conn.Do("EXPIRE", key, expire)
}

// GetVal
func GetVal(conn redis.Conn, key string) (interface{}, error) {
	return conn.Do("GET", key)
}

func GetString(conn redis.Conn, key string) (string, error) {
	return redis.String(conn.Do("GET", key))

}

func GetBool(conn redis.Conn, key string) (bool, error) {
	return redis.Bool(conn.Do("GET", key))

}

func GetInt64(conn redis.Conn, key string) (int64, error) {
	return redis.Int64(conn.Do("GET", key))

}

func GetInt(conn redis.Conn, key string) (int, error) {
	return redis.Int(conn.Do("GET", key))

}

func Incr(conn redis.Conn, key string) (interface{}, error) {
	return conn.Do("INCR", key)
}

func TTL(conn redis.Conn, key string) (int, error) {
	return redis.Int(conn.Do("TTL", key))
}

// ExistKey
func ExistKey(conn redis.Conn, key string) bool {
	exist_key, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		uniledgerlog.Error(err)
	}
	return exist_key
}
