package chain

import (
	"fmt"
	"testing"
	"reflect"
)

/**
 * function : 1.合约查询 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestContractQuery(t *testing.T)  {
	result := ContractQuery("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 2.合约跟踪 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestContractTracking(t *testing.T) {
	result := ContractTracking("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 3.合约资产冻结 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestContractAssetFreeze(t *testing.T)  {
	result := ContractAssetFreeze("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 4.合约资产解冻 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestContractAssetThaw(t *testing.T)  {
	result := ContractAssetThaw("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 5.合约交易创建 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestCreateContractTran(t *testing.T)  {
	result := CreateContractTran("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 6.合约资产转移 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestTransferContractTran(t *testing.T)  {
	result := TransferContractTran("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}
