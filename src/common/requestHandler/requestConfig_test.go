package requestHandler

import (
	"testing"
	"fmt"
)

/**
 * function : 获取config对象
 * param   :
 * return : 
 */
func TestGetConfig(t *testing.T) {
	confi := _GetConfig()
	fmt.Println(confi)
}

/**
 * function : 获取对应api的参数
 * param   :
 * return :
 */
func TestGetParam(t *testing.T) {
	fmt.Println(GetParam("creatTrac"))
}

