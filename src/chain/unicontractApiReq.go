package chain

import "unicontract/src/common/requestHandler"

/**
 * function : 1.合约查询
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func ContractQuery(jsonBody interface{})  *requestHandler.ResponseResult{
	yamlName := "unicontractApiConf.yaml"
	apiName := "ContractQuery"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 2.合约跟踪
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func ContractTracking(jsonBody interface{}) *requestHandler.ResponseResult{
	yamlName := "unicontractApiConf.yaml"
	apiName := "ContractTracking"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 3.合约资产冻结
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func ContractAssetFreeze(jsonBody interface{})  *requestHandler.ResponseResult{
	yamlName := "unicontractApiConf.yaml"
	apiName := "ContractAssetFreeze"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 4.合约资产解冻
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func ContractAssetThaw(jsonBody interface{})  *requestHandler.ResponseResult{
	yamlName := "unicontractApiConf.yaml"
	apiName := "ContractAssetThaw"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 5.合约交易创建
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func CreateContractTran(jsonBody interface{})  *requestHandler.ResponseResult{
	yamlName := "unicontractApiConf.yaml"
	apiName := "CreateContractTran"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 6.合约资产转移
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TransferContractTran(jsonBody interface{})  *requestHandler.ResponseResult{
	yamlName := "unicontractApiConf.yaml"
	apiName := "TransferContractTran"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}
