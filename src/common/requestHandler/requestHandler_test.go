package requestHandler

import (
	"testing"
	"fmt"
	"reflect"
	"unicontract/src/core/protos"
	"github.com/golang/protobuf/proto"
)

/**
 * function: 测试获取请求参数结构体
 * param :
 * return:
 */
func TestNewRequestParam(t *testing.T) {
	head := make(map[interface{}]interface{})
	p :=NewRequestParam("sss","s",head,"ss")
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(p)
}
func GetData() []byte{
	contract := protos.ContractProto{ // proto-buf
		Id:         "2",
		NodePubkey: "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		MainPubkey: "93TEovPuYo6BQFm4ia9ta4qtL1TbAmnk9fV5kxmesAG5",
		Signature:  "2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
		Voters: []string{
			"2kdD14DHpccekjRgK55bgzEuAF5JLubhq3tBRm1sXqDc",
			"JBMja2vDAJxkj9bxxjGzxQpTtavLxajxij41geufRXzs",
			"EtQVTBXJ8onJmXLnkzGBhbxhE3bSPgqvCkeaKtT22Cet",
		},
		Timestamp: "12321",
		Version:   "v1.0",
	}
	data := protos.ContractData{
		Data: &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}
	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}
	return requestBody
}
/**
 * function: 测试请求单个api
 * param :
 * return:
 */
func TestRequestHandler1(t *testing.T) {

	//requestBody := GetData()

	yamlConfig := GetYamlConfig("unicontractApiConf.yaml")
	method,url,head,_ := GetParam(yamlConfig,"ContractTracking")
	param := NewRequestParam(method,url,head,"")
	result,status := RequestHandler(param)
	fmt.Println("==================================")
	fmt.Println(result)
	fmt.Println(status)
	fmt.Println(reflect.TypeOf(status))
}

/**
 * function: 测试多次请求多个api
 * param :
 * return:
 */
func TestRequestHandler(t *testing.T) {

	requestBody := GetData()

	for i := 1;i<=5;i++{
		yamlConfig := GetYamlConfig("unicontractApiConf.yaml")
		method,url,head,_ := GetParam(yamlConfig,"ContractTracking")
		param := NewRequestParam(method,url,head,"")
		result,status := RequestHandler(param)
		fmt.Println(result)
		fmt.Println(status)
		fmt.Println("========================")
		method1,url1,head1,_ := GetParam(yamlConfig,"ContractAssetFreeze")
		param1 := NewRequestParam(method1,url1,head1,requestBody)
		result1,status1 := RequestHandler(param1)
		fmt.Println(result1)
		fmt.Println(status1)
		fmt.Println("========================")
	}

}

