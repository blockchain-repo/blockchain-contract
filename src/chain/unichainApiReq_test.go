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

	jsonBody := `{"beginTime":"1487066476", "endTime":"1487066776"}`

	result,err:= CreateByPayload(jsonBody)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(result.Data)
	fmt.Println(result.Message)
	fmt.Println(result.Code)
	fmt.Println(result.Status)
	fmt.Println(reflect.TypeOf(result.Data))
}

/**
 * function : 2.根据交易ID获取交易 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryByID(t *testing.T){
	jsonBody := `{"tx_id":"440b13ed5442cd1375c3bcd17b91fa9b17a8eb5efbd117cbbde3ee979ce50269"}`
	result,err := QueryByID(jsonBody)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 3.获取区块链中的总交易条数 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryTxsTotal(t *testing.T)  {
	result,err := QueryTxsTotal()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 4.根据指定时间区间获取交易集 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryTxsByRange(t *testing.T)  {
	jsonBody := `{"beginTime":"1493127874364", "endTime":"1493131075550"}`
	result,err := QueryTxsByRange(jsonBody)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 5.获取每区块中包含的交易条数 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryGroupByBlock(t *testing.T)  {
	result,err := QueryGroupByBlock()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 6.根据区块ID获取区块 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryBlockByID(t *testing.T)  {
	result,err := QueryBlockByID(`{"block_id":"e01e291dead26382274a7f3423fdf48ab7c9a930d5d6a1f7c59dee303987c346"}`)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 7.根据区块ID获取区块中的交易 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryTxsByID(t *testing.T)  {
	result,err := QueryTxsByID(`{"block_id":"e01e291dead26382274a7f3423fdf48ab7c9a930d5d6a1f7c59dee303987c346"}`)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 8.根据区块ID获取区块中的交易条数 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryTxsCountByID(t *testing.T)  {
	result,err := QueryTxsCountByID(`{"block_id":"e01e291dead26382274a7f3423fdf48ab7c9a930d5d6a1f7c59dee303987c346"}`)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 9.获取区块链中的总区块数 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryBlockCount(t *testing.T)  {
	result,err := QueryBlockCount()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 10.根据指定时间区间获取区块集 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryBlocksByRange(t *testing.T)  {
	jsonBody := `{"beginTime":"1493127874364", "endTime":"1493131075550"}`
	result,err := QueryBlocksByRange(jsonBody)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 11.获取所有无效区块集 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryInvalidBlockTotal(t *testing.T)  {
	result,err := QueryInvalidBlockTotal()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 12.获取指定时间区间内的无效区块集 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestQueryInvalidBlockByRange(t *testing.T)  {
	jsonBody := `{"beginTime":"1493127874364", "endTime":"1493131075550"}`
	result,err := QueryInvalidBlockByRange(jsonBody)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 13.获取参与投票的节点公钥集 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestPublickeySet(t *testing.T)  {
	result,err := PublickeySet()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 14.根据指定时间区间获取交易创建平均时间 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestTxCreateAvgTimeByRange(t *testing.T)  {
	jsonBody := `{"beginTime":"1493127874364", "endTime":"1493131075550"}`
	result,err := TxCreateAvgTimeByRange(jsonBody)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 15.根据指定时间区间获取区块创建平均时间 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestBlockCreateAvgTimeByRange(t *testing.T)  {
	jsonBody := `{"beginTime":"1493127874364", "endTime":"1493131075550"}`
	result,err := BlockCreateAvgTimeByRange(jsonBody)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 16.根据指定区块ID获取投票时间 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestVoteTimeByBlockID(t *testing.T)  {
	result,err := VoteTimeByBlockID(`{"block_id":"e01e291dead26382274a7f3423fdf48ab7c9a930d5d6a1f7c59dee303987c346"}`)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}

/**
 * function : 17.获取区块的平均投票时间 测试
 * param   : jsonBody interface{}
 * return : requestHandler.ResponseResult struct
 */
func TestVoteAvgTimeByRange(t *testing.T)  {
	result,err := VoteAvgTimeByRange("")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
}
