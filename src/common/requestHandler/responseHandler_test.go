package requestHandler

import (
	"testing"
	"fmt"
	"reflect"
)


var responseBody string = `{"status":"success", "code":200, "message":"xxxxxx", "data":{"id":123}}
`


/**
 * function : 
 * param   :
 * return : 
 */
func TestGetResponseString(t *testing.T) {
	result,status:= GetResponseString(responseBody,200)
	fmt.Println(status)
	fmt.Println(result)
}

func TestGetResponseData(t *testing.T) {
	result := GetResponseData(responseBody,200)
	fmt.Println(result)
	fmt.Println(result.Status)
	fmt.Println(result.Code)
	fmt.Println(result.Message)
	fmt.Println(result.Data)
	fmt.Println(reflect.TypeOf(result.Data))
}