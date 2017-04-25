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
	jsonBody := `{"beginTime":"1487066476", "endTime":"1487066776"}`
	result,err:= CreateContract(jsonBody)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 2.创建合约交易 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestCreateContractTx(t *testing.T) {
	jsonBody := `{"beginTime":"1487066476", "endTime":"1487066776"}`
	result,err := CreateContractTx(jsonBody)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 3.获取合约 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestGetContract(t *testing.T)  {
	jsonBody := `{"contract_id":"1487066476"}`
	result,err := GetContract(jsonBody)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 4.获取合约交易 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestGetContractTx(t *testing.T)  {
	jsonBody := `{"contract_id":"1487066476"}`
	result,err := GetContractTx(jsonBody)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 5.获取合约记录 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestGetContractRecord(t *testing.T)  {
	jsonBody := `{"contract_id":"1487066476"}`
	result,err := GetContractRecord(jsonBody)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 6.冻结资产 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestFreezeAsset(t *testing.T)  {
	jsonBody := `{"contract_id":"1487066476"}`
	result,err := FreezeAsset(jsonBody)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 7.解冻资产 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestUnfreezeAsset(t *testing.T)  {
	jsonBody := `{"contract_id":"1487066476"}`
	result,err := UnfreezeAsset(jsonBody)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 8.查询冻结资产 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestFrozenAsset(t *testing.T)  {
	jsonBody := `{"public_key":"1487066476"}`
	result,err := FrozenAsset(jsonBody)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}
