package acrosschain

import (
	"unicontract/src/common/uniledgerlog"
)

/**
select or insert(operator unichain not contract)
*/
func requestMultiChain(chainList []map[string]interface{}) {
	var dataChannel = make(chan map[string]interface{})
	var chainLen = len(chainList)
	uniledgerlog.Info("chainListLen:", chainLen)
	var resList = make([]map[string]interface{}, 0)
	uniledgerlog.Info("resListLen:", len(resList))

	for index, chainMap := range chainList {
		uniledgerlog.Info("index:", index)
		uniledgerlog.Info("chainMap:", chainMap)
		go func() {
			var resMap = requestToUnichain(index, chainMap)
			dataChannel <- resMap
		}()
	}
	for i := 0; i < chainLen; i++ {
		//TODO timeout
		res := <-dataChannel
		uniledgerlog.Info("chan--", res)
		resList = append(resList, res)
	}

	uniledgerlog.Info("list--", resList)
}

func insertMultiChainWithoutTrans(chainList []map[string]interface{}) {
	var dataChannel = make(chan map[string]interface{})
	var chainLen = len(chainList)
	uniledgerlog.Info("chainListLen:", chainLen)
	var resList = make([]map[string]interface{}, 0)
	uniledgerlog.Info("resListLen:", len(resList))

	for index, chainMap := range chainList {
		uniledgerlog.Info("index:", index)
		uniledgerlog.Info("chainMap:", chainMap)
		go func() {
			var resMap = requestToUnichain(index, chainMap)
			dataChannel <- resMap
		}()
	}
	for i := 0; i < chainLen; i++ {
		//TODO timeout
		res := <-dataChannel
		uniledgerlog.Info("chan--", res)
		resList = append(resList, res)
	}

	uniledgerlog.Info("list--", resList)
}

func insertMultiChainWithTrans(chainList []map[string]interface{}) {
	var dataChannel = make(chan map[string]interface{})
	var chainLen = len(chainList)
	uniledgerlog.Info("chainListLen:", chainLen)
	var resList = make([]map[string]interface{}, 0)
	uniledgerlog.Info("resListLen:", len(resList))

	for index, chainMap := range chainList {
		uniledgerlog.Info("index:", index)
		uniledgerlog.Info("chainMap:", chainMap)
		go func() {
			var resMap = insertOutputAndFreeze(index, chainMap)
			dataChannel <- resMap
		}()
	}
	for i := 0; i < chainLen; i++ {
		//TODO timeout
		res := <-dataChannel
		uniledgerlog.Info("chan--", res)
		resList = append(resList, res)
	}

	uniledgerlog.Info("list--", resList)
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
	uniledgerlog.Info("saveData:", resMap)
	//TODO generate transaction to freeze

	//TODO return txId and output cid
	return resMap
}

func requestToUnichain(index int, chainMap map[string]interface{}) map[string]interface{} {
	// TODO accoding to chainMap  request unichain,, add response to resMap

	var resMap map[string]interface{} = make(map[string]interface{})
	resMap["index"] = index
	uniledgerlog.Info("saveData:", resMap)
	return resMap
}
