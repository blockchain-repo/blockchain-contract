package request

import (
	"fmt"
	"unicontract/src/common/requestHandler"
	"unicontract/src/core/protos"
	"github.com/golang/protobuf/proto"
)

/**
 * function:
 * param :
 * return nil:
 */
func CreatTransaction(){

	yamlConfig := requestHandler.GetYamlConfig("unichainApiConf.yaml")
	_,_,head,_ := requestHandler.GetParam(yamlConfig,"ContractTracking")
	param := requestHandler.NewRequestParam("get","http://www.weather.com.cn/data/sk/101010100.html",head,make(map[interface{}]interface{}))
	requestHandler.RequestHandler(param)
	fmt.Println(param)
}

func Weather(){
	yamlConfig := requestHandler.GetYamlConfig("unichainApiConf.yaml")
	_,_,head,_ := requestHandler.GetParam(yamlConfig,"ContractTracking")
	param := requestHandler.NewRequestParam("get","http://www.weather.com.cn/data/sk/101010100.html",head,"ss:aa")
	jsonResponse,status := requestHandler.RequestHandler(param)
	fmt.Println(jsonResponse)
	fmt.Println(status)
}

func Today(){
	//url := "http://api.kdniao.cc/Ebusiness/EbusinessOrderHandle.aspx"
	//method := "post"
	//headKey := "Content-Type"
	//headValue := "application/x-www-form-urlencoded;charset=utf-8"
	//jsonBody := `{
	//"RequestType":"1002"
	//"OrderCode": "",
     //   "ShipperCode": "SF",
     //   "LogisticCode": "118650888018"
	//}`
	////_,_,_,_,a := requestHandler.GetParam("creatTrac")
	//param := requestHandler.NewRequestParam(method,url,headKey,headValue,jsonBody)
	//jsonResponse,status:= requestHandler.RequestHandler(param)
	//fmt.Println(jsonResponse)
	//fmt.Println(status)
}

func Test(){
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
		Version:   1,
	}
	data := protos.ContractData{
		Data: &contract,
		Token: "ZDNkM0xtWjFkSFZ5WlhabGNpNWpiMjA9",
	}
	requestBody, err := proto.Marshal(&data)
	if err != nil {
		fmt.Println("error ", err.Error())
	}

	for i := 1;i<=5;i++{
		yamlConfig := requestHandler.GetYamlConfig("unichainApiConf.yaml")
		method,url,head,_ := requestHandler.GetParam(yamlConfig,"ContractTracking")
		param := requestHandler.NewRequestParam(method,url,head,requestBody)
		result,status := requestHandler.RequestHandler(param)
		fmt.Println(result)
		fmt.Println(status)
		fmt.Println("========================")
		method1,url1,head1,_ := requestHandler.GetParam(yamlConfig,"ContractAssetFreeze")
		param1 := requestHandler.NewRequestParam(method1,url1,head1,requestBody)
		result1,status1 := requestHandler.RequestHandler(param1)
		fmt.Println(result1)
		fmt.Println(status1)
		fmt.Println("========================")
	}

}
