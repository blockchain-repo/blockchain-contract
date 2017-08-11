package chain

import (
	"errors"
	"unicontract/src/common"
	"unicontract/src/common/requestHandler"
	"unicontract/src/common/uniledgerlog"
)

/**
 * function : 1.单条payload交易创建接口
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func CreateByPayload(jsonBody interface{}) (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking CreateByPayload Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "CreateByPayload"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody, "")
		uniledgerlog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 2.根据交易ID获取交易
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryByID(jsonBody interface{}) (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking QueryByID Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryByID"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody, "")
		uniledgerlog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 3.获取区块链中的总交易条数
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryTxsTotal() (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking QueryTxsTotal Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryTxsTotal"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, "", "")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,"")
}

/**
 * function : 4.根据指定时间区间获取交易集
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryTxsByRange(jsonBody interface{}) (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking QueryTxsByRange Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryTxsByRange"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody, "")
		uniledgerlog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 5.获取每区块中包含的交易条数
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryGroupByBlock() (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking QueryGroupByBlock Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryGroupByBlock"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, "", "")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,"")
}

/**
 * function : 6.根据区块ID获取区块
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryBlockByID(jsonBody interface{}) (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking QueryBlockByID Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryBlockByID"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody, "")
		uniledgerlog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 7.根据区块ID获取区块中的交易
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryTxsByID(jsonBody interface{}) (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking QueryTxsByID Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryTxsByID"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody, "")
		uniledgerlog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 8.根据区块ID获取区块中的交易条数
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryTxsCountByID(jsonBody interface{}) (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking QueryTxsCountByID Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryTxsCountByID"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody, "")
		uniledgerlog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 9.获取区块链中的总区块数
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryBlockCount() (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking QueryBlockCount Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryBlockCount"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, "", "")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,"")
}

/**
 * function : 10.根据指定时间区间获取区块集
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryBlocksByRange(jsonBody interface{}) (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking QueryBlocksByRange Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryBlocksByRange"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody, "")
		uniledgerlog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 11.获取所有无效区块集
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryInvalidBlockTotal() (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking QueryInvalidBlockTotal Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryInvalidBlockTotal"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, "", "")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,"")
}

/**
 * function : 12.获取指定时间区间内的无效区块集
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func QueryInvalidBlockByRange(jsonBody interface{}) (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking QueryInvalidBlockByRange Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "QueryInvalidBlockByRange"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody, "")
		uniledgerlog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 13.获取参与投票的节点公钥集
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func PublickeySet() (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking PublickeySet Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "PublickeySet"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, "", "")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,"")
}

/**
 * function : 14.根据指定时间区间获取交易创建平均时间
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TxCreateAvgTimeByRange(jsonBody interface{}) (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking TxCreateAvgTimeByRange Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "TxCreateAvgTimeByRange"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody, "")
		uniledgerlog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 15.根据指定时间区间获取区块创建平均时间
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func BlockCreateAvgTimeByRange(jsonBody interface{}) (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking BlockCreateAvgTimeByRange Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "BlockCreateAvgTimeByRange"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody, "")
		uniledgerlog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 16.根据指定区块ID获取投票时间
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func VoteTimeByBlockID(jsonBody interface{}) (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking VoteTimeByBlockID Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "VoteTimeByBlockID"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody, "")
		uniledgerlog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 17.获取区块的平均投票时间
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func VoteAvgTimeByRange(jsonBody interface{}) (*requestHandler.ResponseResult, error) {

	uniledgerlog.Debug(" begin invoking VoteAvgTimeByRange Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "VoteAvgTimeByRange"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody, "")
		uniledgerlog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

func GetUnspentTxs(jsonBody interface{}) (*requestHandler.ResponseResult, error) {
	uniledgerlog.Debug(" begin invoking GetUnspentTxs Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "GetUnspentTxs"
	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody, "")
		uniledgerlog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
}

func GetFreezeUnspentTxs(jsonBody interface{}) (*requestHandler.ResponseResult, error) {
	uniledgerlog.Debug(" begin invoking GetFreezeUnspentTxs Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "GetFreezeUnspentTxs"
	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody, "")
		uniledgerlog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
}

func GetContractById(jsonBody interface{}) (*requestHandler.ResponseResult, error) {
	uniledgerlog.Debug(" begin invoking GetContractById Api")
	yamlName := "unichainApiConf.yaml"
	apiName := "GetContractById"
	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody, "")
		uniledgerlog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
	})

	return res, err
}
