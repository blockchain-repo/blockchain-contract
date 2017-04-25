package chain

import (
	"unicontract/src/common/requestHandler"
	"unicontract/src/common"
	"errors"
)

import (
	beegoLog "github.com/astaxie/beego/logs"
)

/**
 * function : 1.创建合约
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func CreateContract(jsonBody interface{})  (*requestHandler.ResponseResult,error){

	beegoLog.Debug(" begin invoking CreateContract Api")
	yamlName := "unicontractApiConf.yaml"
	apiName := "CreateContract"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody)
		beegoLog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
		beegoLog.Error(err)
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 2.创建合约交易
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func CreateContractTx(jsonBody interface{}) (*requestHandler.ResponseResult,error){

	beegoLog.Debug(" begin invoking CreateContractTx Api")
	yamlName := "unicontractApiConf.yaml"
	apiName := "CreateContractTx"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody)
		beegoLog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
		beegoLog.Error(err)
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 3.获取合约
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func GetContract(jsonBody interface{})  (*requestHandler.ResponseResult,error){

	beegoLog.Debug(" begin invoking GetContract Api")
	yamlName := "unicontractApiConf.yaml"
	apiName := "GetContract"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody)
		beegoLog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
		beegoLog.Error(err)
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 4.获取合约交易
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func GetContractTx(jsonBody interface{})  (*requestHandler.ResponseResult,error){

	beegoLog.Debug(" begin invoking GetContractTx Api")
	yamlName := "unicontractApiConf.yaml"
	apiName := "GetContractTx"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody)
		beegoLog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
		beegoLog.Error(err)
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 5.获取合约记录
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func GetContractRecord(jsonBody interface{})  (*requestHandler.ResponseResult,error){

	beegoLog.Debug(" begin invoking GetContractRecord Api")
	yamlName := "unicontractApiConf.yaml"
	apiName := "GetContractRecord"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody)
		beegoLog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
		beegoLog.Error(err)
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 6.冻结资产
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func FreezeAsset(jsonBody interface{})  (*requestHandler.ResponseResult,error){

	beegoLog.Debug(" begin invoking FreezeAsset Api")
	yamlName := "unicontractApiConf.yaml"
	apiName := "FreezeAsset"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody)
		beegoLog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
		beegoLog.Error(err)
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 7.解冻资产
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func UnfreezeAsset(jsonBody interface{})  (*requestHandler.ResponseResult,error){

	beegoLog.Debug(" begin invoking UnfreezeAsset Api")
	yamlName := "unicontractApiConf.yaml"
	apiName := "UnfreezeAsset"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody)
		beegoLog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
		beegoLog.Error(err)
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}

/**
 * function : 8.查询冻结资产
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func FrozenAsset(jsonBody interface{})  (*requestHandler.ResponseResult,error){

	beegoLog.Debug(" begin invoking FrozenAsset Api")
	yamlName := "unicontractApiConf.yaml"
	apiName := "FrozenAsset"

	var res *requestHandler.ResponseResult
	var err error
	common.Try(func() {
		res = requestHandler.GetRequestResult(yamlName, apiName, jsonBody)
		beegoLog.Debug("request finish....")
	}, func(e interface{}) {
		err = errors.New("connect reflused")
		beegoLog.Error(err)
	})

	return res, err
	//return requestHandler.GetRequestResult(yamlName,apiName,jsonBody)
}
