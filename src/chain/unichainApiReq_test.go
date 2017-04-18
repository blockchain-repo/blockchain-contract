package chain

import (
	"testing"
	"fmt"
	"reflect"
)

/**
 * function : 1.单条payload交易创建接口 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestCreateByPayload(t *testing.T) {
	result := CreateByPayload("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 2.根据交易ID获取交易 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryByID(t *testing.T){
	result := QueryByID("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 3.获取区块链中的总交易条数 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryTxsTotal(t *testing.T)  {
	result := QueryTxsTotal()
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 4.根据指定时间区间获取交易集 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryTxsByRange(t *testing.T)  {
	result := QueryTxsByRange("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 5.获取每区块中包含的交易条数 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryGroupByBlock(t *testing.T)  {
	result := QueryGroupByBlock()
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 6.根据区块ID获取区块 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryBlockByID(t *testing.T)  {
	result := QueryBlockByID("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 7.根据区块ID获取区块中的交易 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryTxsByID(t *testing.T)  {
	result := QueryTxsByID("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 8.根据区块ID获取区块中的交易条数 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryTxsCountByID(t *testing.T)  {
	result := QueryTxsCountByID("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 9.获取区块链中的总区块数 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryBlockCount(t *testing.T)  {
	result := QueryBlockCount()
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 10.根据指定时间区间获取区块集 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryBlocksByRange(t *testing.T)  {
	result := QueryBlocksByRange("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 11.获取所有无效区块集 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryInvalidBlockTotal(t *testing.T)  {
	result := QueryInvalidBlockTotal()
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 12.获取指定时间区间内的无效区块集 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryInvalidBlockByRange(t *testing.T)  {
	result := QueryInvalidBlockByRange("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 13.获取参与投票的节点公钥集 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestPublickeySet(t *testing.T)  {
	result := PublickeySet()
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 14.根据指定时间区间获取交易创建平均时间 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestTxCreateAvgTimeByRange(t *testing.T)  {
	result := TxCreateAvgTimeByRange("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 15.根据指定时间区间获取区块创建平均时间 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestBlockCreateAvgTimeByRange(t *testing.T)  {
	result := BlockCreateAvgTimeByRange("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 16.根据指定区块ID获取投票时间 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestVoteTimeByBlockID(t *testing.T)  {
	result := VoteTimeByBlockID("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 17.获取区块的平均投票时间 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestVoteAvgTimeByRange(t *testing.T)  {
	result := VoteAvgTimeByRange("")
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}
