package chain

import (
	"fmt"
	"testing"
)

var jsonBody string= `{
						"status": "success",
						"code": "200",
						"message": "wsp",
						"data": "aa:bb"

						}`

/**
 * function: 获取response struct
 * param : jsonBody
 * return: 返回ResponseResult struct
 */
func TestContractQuery2(t *testing.T) {
	jsonBody := `{
    "status": "success",
    "code": "200",
    "message": "wsp",
    "data": "aa:bb"
	}`
	responseResult := _GetResponseData(jsonBody)

	fmt.Println(responseResult)
	fmt.Println(responseResult.status)
	fmt.Println(responseResult.code)
	fmt.Println(responseResult.message)
	fmt.Println(responseResult.data)
}

/**
 * function : 合约查询测试
 * param   :
 * return : data,response status
 */
func TestContractQuery(t *testing.T) {
	responseResult,status := ContractQuery(jsonBody)
	fmt.Println(responseResult)
	fmt.Println(status)

}

/**
 * function : 合约跟踪
 * param   :
 * return : data,response status
 */
func TestContractTracking(t *testing.T) {
	responseResult,status := ContractTracking(jsonBody)
	fmt.Println(responseResult)
	fmt.Println(status)
}

/**
 * function : 合约资产冻结测试
 * param   :
 * return : data,response status
 */
func TestContractAssetFreeze(t *testing.T) {
	responseResult,status := ContractAssetFreeze(jsonBody)
	fmt.Println(responseResult)
	fmt.Println(status)
}

/**
 * function : 合约资产解冻测试
 * param   :
 * return : data,response status
 */
func TestContractAssetThaw(t *testing.T) {
	responseResult,status := ContractAssetThaw(jsonBody)
	fmt.Println(responseResult)
	fmt.Println(status)
}

/**
 * function : 合约交易创建测试
 * param   :
 * return : data,response status
 */
func TestCreateContractTran(t *testing.T) {
	responseResult,status := CreateContractTran(jsonBody)
	fmt.Println(responseResult)
	fmt.Println(status)
}

/**
 * function : 合约资产转移测试
 * param   :
 * return : data,response status
 */
func TestTransferContractTran(t *testing.T) {
	responseResult,status := TransferContractTran(jsonBody)
	fmt.Println(responseResult)
	fmt.Println(status)
}
