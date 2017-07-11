/*************************************************************************
  > File Name: taskexecute.go
  > Module:
  > Function: 从任务待执行队列（gchTaskQueue）中取任务，然后放入执行机执行
  > Author: wangyp
  > Company:
  > Department:
  > Mail: wangyepeng87@163.com
  > Created Time: 2017.05.08
 ************************************************************************/
package scanengine

import (
	"encoding/json"
	"fmt"
)

import (
	"unicontract/src/chain"
	"unicontract/src/common/monitor"
	"unicontract/src/common/uniledgerlog"
	engineCommon "unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine"
)

//---------------------------------------------------------------------------
func _TaskExecute() {
	for {
		uniledgerlog.Debug("wait for ContractTask ...")
		strContractTask, ok := <-gchTaskQueue
		if !ok {
			break
		}
		uniledgerlog.Debug("get ContractTask")
		//wsp@monitor
		monitor.Monitor.Count("task_running", -1)
		uniledgerlog.Debug("query contract base on contractId")
		jsonBody := fmt.Sprintf("{\"contract_hash_id\":\"%s\"}", strContractTask.ContractHashId)
		//responseResult:  requestHandler.ResponseResult, data中存的是完整的Output结构体
		responseResult, err := chain.GetTxByConHashId(jsonBody)
		if err != nil || responseResult.Data == nil {
			uniledgerlog.Error(err)
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}

		if responseResult.Code != _HTTP_OK {
			uniledgerlog.Error("responseResult.Code is [ %d ]", responseResult.Code)
			uniledgerlog.Error("responseResult.Message is [ %s ]", responseResult.Message)
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}

		uniledgerlog.Debug("get responseResult data")
		// 1
		contractData, ok := responseResult.Data.(interface{})
		if !ok {
			uniledgerlog.Error("responseResult.Data.(interface{}) is error")
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}

		// 2
		mapData, ok := contractData.([]interface{})
		if !ok {
			uniledgerlog.Error("contractData.([]map[string]interface{}) error")
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}

		if len(mapData) == 0 { // 没有查询到contract
			uniledgerlog.Error("get contract is null")
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}

		slContractData, err := json.Marshal(mapData[0])
		if err != nil {
			uniledgerlog.Error(err)
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}
		//uniledgerlog.Debug(string(slContractData))

		uniledgerlog.Debug("contract execute")
		go _Execute(string(slContractData), strContractTask.ContractId, strContractTask.ContractHashId)
	}

	gwgTaskExe.Done()
}

//---------------------------------------------------------------------------
func _Execute(strData, strContractID, strContractHashID string) {
	contractExecuter := execengine.NewContractExecuter()
	//strData为完整的Output结构体
	task_load_time := monitor.Monitor.NewTiming()
	err := contractExecuter.Load(strData)
	task_load_time.Send("task_load")
	if err != nil {
		uniledgerlog.Error(err)
		_UpdateToFailed(strContractID, strContractHashID)
		return
	}
	//执行引擎初始化环境
	contractExecuter.Prepare()
	//执行机启动合约执行
	task_execute_time := monitor.Monitor.NewTiming()
	ret, err := contractExecuter.Start()
	task_execute_time.Send("task_execute")
	if err != nil {
		uniledgerlog.Error(err)
		_UpdateToFailed(strContractID, strContractHashID)
		return
	}
	if ret == 0 {
		uniledgerlog.Warn("合约执行过程中，某任务没有达到执行条件，暂时退出，等待下轮扫描再次加载执行")
	} else if ret == -1 {
		uniledgerlog.Warn("合约执行过程中，某任务执行失败，暂时退出，等待下轮扫描再次加载执行")
	} else if ret == 1 {
		uniledgerlog.Info("合约任务执行完成")
	}
	//执行机销毁合约
	contractExecuter.Destory()

}

//---------------------------------------------------------------------------
func _UpdateToWait(strContractID, strContractHashID string) {
	err := engineCommon.UpdateMonitorWait(strContractID, strContractHashID, "0", "", 1)
	if err != nil {
		uniledgerlog.Error(err)
	}
}

//---------------------------------------------------------------------------
func _UpdateToFailed(strContractID, strContractHashID string) {
	err := engineCommon.UpdateMonitorFail(strContractID, strContractHashID, "0", "", 1)
	if err != nil {
		uniledgerlog.Error(err)
	}
}

//---------------------------------------------------------------------------
