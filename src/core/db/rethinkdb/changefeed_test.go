package rethinkdb

import (
	"fmt"
	"testing"
)

func Test_Changefeed(t *testing.T) {
	var value interface{}
	res := Changefeed("bigchain", "backlog")
	for res.Next(&value) {
		fmt.Println(value)
	}
}
