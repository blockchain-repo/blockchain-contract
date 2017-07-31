package acrosschain

import (
	"github.com/astaxie/beego/logs"
)

/**
both select and insert
*/
func operateMultiChainWithNoTrans(chainList []map[string]interface{}) {
	var dataChannel = make(chan map[string]interface{})
	var chainLen = len(chainList)
	logs.Info("chainListLen:", chainLen)
	var resList = make([]map[string]interface{}, 0)
	logs.Info("resListLen:", len(resList))

	for index, chainMap := range chainList {
		logs.Info("index:", index)
		logs.Info("chainMap:", chainMap)
		go func() {
			var resMap = requestToUnichain(index, chainMap)
			dataChannel <- resMap
		}()
	}
	for i := 0; i < chainLen; i++ {
		//TODO timeout
		res := <-dataChannel
		logs.Info("chan--", res)
		resList = append(resList, res)
	}

	logs.Info("list--", resList)
}

func insertMultiChainWithTrans(chainList []map[string]interface{}) {
	var dataChannel = make(chan map[string]interface{})
	var chainLen = len(chainList)
	logs.Info("chainListLen:", chainLen)
	var resList = make([]map[string]interface{}, 0)
	logs.Info("resListLen:", len(resList))

	for index, chainMap := range chainList {
		logs.Info("index:", index)
		logs.Info("chainMap:", chainMap)
		go func() {
			var resMap = insertOutputAndFreeze(index, chainMap)
			dataChannel <- resMap
		}()
	}
	for i := 0; i < chainLen; i++ {
		//TODO timeout
		res := <-dataChannel
		logs.Info("chan--", res)
		resList = append(resList, res)
	}

	logs.Info("list--", resList)
}

func unFreezeMultiChainAssert(chainList []map[string]interface{}) {
	//todo unfreeze
}

func resetMultiChainAssert() {
	//todo reserve owner and transfer

}

func insertOutputAndFreeze(index int, chainMap map[string]interface{}) map[string]interface{} {
	var resMap map[string]interface{} = make(map[string]interface{})
	resMap["index"] = index
	logs.Info("saveData:", resMap)
	//TODO generate transaction to freeze

	//TODO return txId and output cid
	return resMap
}

func requestToUnichain(index int, chainMap map[string]interface{}) map[string]interface{} {
	// TODO accoding to chainMap  request unichain,, add response to resMap

	var resMap map[string]interface{} = make(map[string]interface{})
	resMap["index"] = index
	logs.Info("saveData:", resMap)
	return resMap
}
