package api

import (
	"github.com/garyburd/redigo/redis"
	"testing"
	"unicontract/src/common/uniledgerlog"
)

func ExampleNewClient(t *testing.T) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		uniledgerlog.Error(err)
		return
	}
	defer conn.Close()
}
