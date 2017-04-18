package chain

import "unicontract/src/common/requestHandler"

/**
 * function : 1.单条payload交易创建接口
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func CreateByPayload(jsonBody interface{}) *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "CreateByPayload"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 2.根据交易ID获取交易
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryByID(jsonBody interface{}) *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryByID"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 3.获取区块链中的总交易条数
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryTxsTotal() *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryTxsTotal"
	return requestHandler.GetRequestResult(yamlName,apiName,"")
}

/**
 * function : 4.根据指定时间区间获取交易集
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryTxsByRange(jsonBody interface{}) *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryTxsByRange"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 5.获取每区块中包含的交易条数
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryGroupByBlock() *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryGroupByBlock"
	return requestHandler.GetRequestResult(yamlName,apiName,"")
}

/**
 * function : 6.根据区块ID获取区块
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryBlockByID(jsonBody interface{}) *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryBlockByID"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 7.根据区块ID获取区块中的交易
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryTxsByID(jsonBody interface{}) *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryTxsByID"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 8.根据区块ID获取区块中的交易条数
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryTxsCountByID(jsonBody interface{}) *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryTxsCountByID"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 9.获取区块链中的总区块数
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryBlockCount() *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryBlockCount"
	return requestHandler.GetRequestResult(yamlName,apiName,"")
}

/**
 * function : 10.根据指定时间区间获取区块集
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryBlocksByRange(jsonBody interface{}) *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryBlocksByRange"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 11.获取所有无效区块集
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryInvalidBlockTotal() *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryInvalidBlockTotal"
	return requestHandler.GetRequestResult(yamlName,apiName,"")
}

/**
 * function : 12.获取指定时间区间内的无效区块集
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryInvalidBlockByRange(jsonBody interface{}) *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryInvalidBlockByRange"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 13.获取参与投票的节点公钥集
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func PublickeySet() *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "PublickeySet"
	return requestHandler.GetRequestResult(yamlName,apiName,"")
}

/**
 * function : 14.根据指定时间区间获取交易创建平均时间
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TxCreateAvgTimeByRange(jsonBody interface{}) *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "TxCreateAvgTimeByRange"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 15.根据指定时间区间获取区块创建平均时间
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func BlockCreateAvgTimeByRange(jsonBody interface{}) *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "BlockCreateAvgTimeByRange"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 16.根据指定区块ID获取投票时间
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func VoteTimeByBlockID(jsonBody interface{}) *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "VoteTimeByBlockID"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 17.获取区块的平均投票时间
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func VoteAvgTimeByRange(jsonBody interface{}) *requestHandler.ResponseResult{
	yamlName := "unichainApiConf.yaml"
	apiName := "VoteAvgTimeByRange"
	return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}












































