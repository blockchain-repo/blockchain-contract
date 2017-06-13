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
	beegoLog "github.com/astaxie/beego/logs"
)

import (
	"unicontract/src/chain"
	"unicontract/src/common/monitor"
	engineCommon "unicontract/src/core/engine/common"
	"unicontract/src/core/engine/execengine"
)

//---------------------------------------------------------------------------
func _TaskExecute() {
	for {
		beegoLog.Debug("wait for ContractTask ...")
		strContractTask, ok := <-gchTaskQueue
		if !ok {
			break
		}
		beegoLog.Debug("get ContractTask")
		//wsp@monitor
		monitor.Monitor.Count("task_running", -1)
		beegoLog.Debug("query contract base on contractId")
		jsonBody := fmt.Sprintf("{\"contract_hash_id\":\"%s\"}", strContractTask.ContractHashId)
		//responseResult:  requestHandler.ResponseResult, data中存的是完整的Output结构体
		responseResult, err := chain.GetTxByConHashId(jsonBody)
		if err != nil || responseResult.Data == nil {
			beegoLog.Error(err)
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}

		if responseResult.Code != _HTTP_OK {
			beegoLog.Error("responseResult.Code is [ %d ]", responseResult.Code)
			beegoLog.Error("responseResult.Message is [ %s ]", responseResult.Message)
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}

		beegoLog.Debug("get responseResult data")
		// 1
		contractData, ok := responseResult.Data.(interface{})
		if !ok {
			beegoLog.Error("responseResult.Data.(interface{}) is error")
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}

		// 2
		mapData, ok := contractData.([]interface{})
		if !ok {
			beegoLog.Error("contractData.([]map[string]interface{}) error")
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}

		if len(mapData) == 0 { // 没有查询到contract
			beegoLog.Error("get contract is null")
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}

		slContractData, err := json.Marshal(mapData[0])
		if err != nil {
			beegoLog.Error(err)
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}
		beegoLog.Debug(string(slContractData))

		beegoLog.Debug("contract execute")
		go _Execute(string(slContractData), strContractTask.ContractId, strContractTask.ContractHashId)
	}

	gwgTaskExe.Done()
}

//---------------------------------------------------------------------------
func _Execute(strData, strContractID, strContractHashID string) {
	task_execute_time := monitor.Monitor.NewTiming()
	contractExecuter := execengine.NewContractExecuter()
	//strData为完整的Output结构体
	err := contractExecuter.Load(strData)
	if err != nil {
		beegoLog.Error(err)
		_UpdateToFailed(strContractID, strContractHashID)
		return
	}
	//执行引擎初始化环境
	contractExecuter.Prepare()
	//执行机启动合约执行
	ret, err := contractExecuter.Start()
	if err != nil {
		beegoLog.Error(err)
		_UpdateToFailed(strContractID, strContractHashID)
		return
	}
	if ret == 0 {
		beegoLog.Error("合约执行过程中，某任务没有达到执行条件，暂时退出，等待下轮扫描再次加载执行")
		monitor.Monitor.Count("task_execute_fail", 1)
	} else if ret == -1 {
		beegoLog.Error("合约执行过程中，某任务执行失败，暂时退出，等待下轮扫描再次加载执行")
		monitor.Monitor.Count("task_execute_fail", 1)
	} else if ret == 1 {
		beegoLog.Debug("合约执行完成")
	}
	//执行机销毁合约
	contractExecuter.Destory()
	task_execute_time.Send("task_execute")
}

//---------------------------------------------------------------------------
func _UpdateToWait(strContractID, strContractHashID string) {
	err := engineCommon.UpdateMonitorWait(strContractID, strContractHashID, "0", "", 1)
	if err != nil {
		beegoLog.Error(err)
	}
}

//---------------------------------------------------------------------------
func _UpdateToFailed(strContractID, strContractHashID string) {
	err := engineCommon.UpdateMonitorFail(strContractID, strContractHashID, "0", "", 1)
	if err != nil {
		beegoLog.Error(err)
	}
}

//---------------------------------------------------------------------------
