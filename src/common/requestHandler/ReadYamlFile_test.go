package requestHandler

import (
	"testing"
	"fmt"
	"reflect"
)

/**
 * function : 获取config参数
 * param   :
 * return : 
 */
func TestGetConfig2(t *testing.T) {
	config := GetYamlConfig("unichainApiConf.yaml")
	fmt.Println(config)
	fmt.Println(reflect.TypeOf(config))
}
