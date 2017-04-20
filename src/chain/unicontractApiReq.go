package chain

import "unicontract/src/common/requestHandler"

/**
 * function : 1.创建合约
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func CreateContract(jsonBody interface{})  *requestHandler.ResponseResult{
	yamlName := "unicontractApiConf.yaml"
	apiName := "CreateContract"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 2.创建合约交易
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func CreateContractTx(jsonBody interface{}) *requestHandler.ResponseResult{
	yamlName := "unicontractApiConf.yaml"
	apiName := "CreateContractTx"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 3.获取合约
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func GetContract(jsonBody interface{})  *requestHandler.ResponseResult{
	yamlName := "unicontractApiConf.yaml"
	apiName := "GetContract"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 4.获取合约交易
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func GetContractTx(jsonBody interface{})  *requestHandler.ResponseResult{
	yamlName := "unicontractApiConf.yaml"
	apiName := "GetContractTx"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 5.获取合约记录
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func GetContractRecord(jsonBody interface{})  *requestHandler.ResponseResult{
	yamlName := "unicontractApiConf.yaml"
	apiName := "GetContractRecord"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 6.冻结资产
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func FreezeAsset(jsonBody interface{})  *requestHandler.ResponseResult{
	yamlName := "unicontractApiConf.yaml"
	apiName := "FreezeAsset"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 7.解冻资产
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func UnfreezeAsset(jsonBody interface{})  *requestHandler.ResponseResult{
	yamlName := "unicontractApiConf.yaml"
	apiName := "UnfreezeAsset"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 8.查询冻结资产
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func FrozenAsset(jsonBody interface{})  *requestHandler.ResponseResult{
	yamlName := "unicontractApiConf.yaml"
	apiName := "FrozenAsset"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}
