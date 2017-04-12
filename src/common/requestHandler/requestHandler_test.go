package requestHandler

import (
	"testing"
	"fmt"
	"reflect"
)

/**
 * function:
 * param :
 * return nil:
 */
func TestNewRequestParam(t *testing.T) {
	head := make(map[interface{}]interface{})
	p :=NewRequestParam("sss","s",head,"ss")
	fmt.Println(reflect.TypeOf(p))
}

