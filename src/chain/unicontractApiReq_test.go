package chain

import (
	"fmt"
	"testing"
	"reflect"
)

/**
 * function : 1.创建合约 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestCreateContract(t *testing.T)  {
	result := CreateContract("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 2.创建合约交易 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestCreateContractTx(t *testing.T) {
	result := CreateContractTx("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 3.获取合约 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestGetContract(t *testing.T)  {
	result := GetContract("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 4.获取合约交易 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestGetContractTx(t *testing.T)  {
	result := GetContractTx("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 5.获取合约记录 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestGetContractRecord(t *testing.T)  {
	result := GetContractRecord("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 6.冻结资产 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestFreezeAsset(t *testing.T)  {
	result := FreezeAsset("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 7.解冻资产 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestUnfreezeAsset(t *testing.T)  {
	result := UnfreezeAsset("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 8.查询冻结资产 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestFrozenAsset(t *testing.T)  {
	result := FrozenAsset("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}
