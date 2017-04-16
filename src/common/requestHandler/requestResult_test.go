package requestHandler

import (
	"testing"
	"fmt"
	"reflect"
)
/**
 * function : 
 * param   :
 * return : 
 */
func TestGetRequestResult(t *testing.T) {

	yamlName := "unicontractApiConf.yaml"
	apiName := "ContractTracking"
	//data := GetData2()
	responseResult := GetRequestResult(yamlName,apiName,"")
	fmt.Println(responseResult)
	fmt.Println(reflect.TypeOf(responseResult))
}

func TestGetRequestResult1(t *testing.T) {
	yamlName := "unicontractApiConf.yaml"
	apiName := "ContractTracking"
	//data := GetData2()
	result,status:= GetRequestResult1(yamlName,apiName,"")
	fmt.Println(status)
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}