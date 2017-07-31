package acrosschain

import "testing"

func Test_operateMultiChainWithNoTrans(t *testing.T) {
	var chainMap1 map[string]interface{} = make(map[string]interface{})
	chainMap1["name"] = "map1"
	var chainMap2 map[string]interface{} = make(map[string]interface{})
	chainMap2["name"] = "map2"
	var chainMap3 map[string]interface{} = make(map[string]interface{})
	chainMap3["name"] = "map3"
	var chainList []map[string]interface{}
	chainList = append(chainList, chainMap1)
	chainList = append(chainList, chainMap2)
	chainList = append(chainList, chainMap3)

	operateMultiChainWithNoTrans(chainList)
}
