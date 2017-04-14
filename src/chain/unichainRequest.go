package chain

import (
	"encoding/json"
	"log"
	"unicontract/src/common/requestHandler"
)

/**
 * 响应结果struct
 */
type ResponseResult struct {
	Status       string
	Code         string
	Message      string
	Data         string
}

/**
 * function: 初始化响应结果
 * param : status,code,message,data
 * return: 返回response result
 */
func _NewResponseResult(status string,code string,message string,data string) *ResponseResult{

	result := &ResponseResult{
		status,
		code,
		message,
		data,
	}
	return result
}

/**
 * function: 获取response struct
 * param : jsonBody
 * return: 返回ResponseResult struct
 */
func _GetResponseData(jsonBody string) *ResponseResult{

	responseMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonBody),&responseMap)
	if err != nil{
		log.Fatal(err.Error())
	}

	status := requestHandler.TypeToString(responseMap["status"])
	code := requestHandler.TypeToString(responseMap["code"])
	message := requestHandler.TypeToString(responseMap["message"])
	data := requestHandler.TypeToString(responseMap["data"])

	return _NewResponseResult(status,code,message,data)
}

/**
 * function : 合约查询
 * param   :
 * return : data,response status
 */
func ContractQuery(jsonBody string) (*ResponseResult,string){

	method,url,head,_ := requestHandler.GetParam("ContractQuery")
	param := requestHandler.NewRequestParam(method,url,head,jsonBody)

	responseStr,status := requestHandler.RequestHandler(param)
	responseResult := _GetResponseData(responseStr)

	return responseResult,status
}

/**
 * function : 合约跟踪
 * param   :
 * return : data,response status
 */
func ContractTracking(jsonBody string) (*ResponseResult,string){

	method,url,head,_ := requestHandler.GetParam("ContractTracking")
	param := requestHandler.NewRequestParam(method,url,head,jsonBody)

	responseStr,status := requestHandler.RequestHandler(param)
	responseResult := _GetResponseData(responseStr)

	return responseResult,status
}

/**
 * function : 合约资产冻结
 * param   :
 * return : data,response status
 */
func ContractAssetFreeze(jsonBody string) (*ResponseResult,string){

	method,url,head,_ := requestHandler.GetParam("ContractAssetFreeze")
	param := requestHandler.NewRequestParam(method,url,head,jsonBody)

	responseStr,status := requestHandler.RequestHandler(param)
	responseResult := _GetResponseData(responseStr)

	return responseResult,status
}

/**
 * function : 合约资产解冻
 * param   :
 * return : data,response status
 */
func ContractAssetThaw(jsonBody string) (*ResponseResult,string){

	method,url,head,_ := requestHandler.GetParam("ContractAssetThaw")
	param := requestHandler.NewRequestParam(method,url,head,jsonBody)

	responseStr,status := requestHandler.RequestHandler(param)
	responseResult := _GetResponseData(responseStr)

	return responseResult,status
}

/**
 * function : 合约交易创建
 * param   :
 * return : data,response status
 */
func CreateContractTran(jsonBody string) (*ResponseResult,string){

	method,url,head,_ := requestHandler.GetParam("CreateContractTran")
	param := requestHandler.NewRequestParam(method,url,head,jsonBody)

	responseStr,status := requestHandler.RequestHandler(param)
	responseResult := _GetResponseData(responseStr)

	return responseResult,status
}

/**
 * function : 合约资产转移
 * param   :
 * return : data,response status
 */
func TransferContractTran(jsonBody string) (*ResponseResult,string){

	method,url,head,_ := requestHandler.GetParam("TransferContractTran")
	param := requestHandler.NewRequestParam(method,url,head,jsonBody)

	responseStr,status := requestHandler.RequestHandler(param)
	responseResult := _GetResponseData(responseStr)

	return responseResult,status
}
