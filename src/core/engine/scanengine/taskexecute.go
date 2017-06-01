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

		beegoLog.Debug("query contract base on contractId")
		jsonBody := fmt.Sprintf("{\"contract_hash_id\":\"%s\"}", strContractTask.ContractHashId)
		//responseResult:  requestHandler.ResponseResult, data中存的是完整的Output结构体
		responseResult, err := chain.GetTxByConHashId(jsonBody)
		if err != nil || responseResult.Data == nil {
			beegoLog.Error(err)
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId, "")
			continue
		}

		if responseResult.Code != _HTTP_OK {
			beegoLog.Error("responseResult.Code is [ %d ]", responseResult.Code)
			beegoLog.Error("responseResult.Message is [ %s ]", responseResult.Message)
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId, "")
			continue
		}

		beegoLog.Debug("contract execute")
		contractData, ok := responseResult.Data.([]interface{})
		if !ok || len(contractData) == 0 {
			beegoLog.Error("responseResult.Data is not ok for type []interface {}. type is %T, or value is [], value is %+v",
				responseResult.Data, responseResult.Data)
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId, "")
			continue
		}

		slContractData, _ := json.Marshal(contractData[0])
		beegoLog.Debug(string(slContractData))

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
		_UpdateToFailed(strContractID, strContractHashID, "")
		return
	}
	//执行引擎初始化环境
	contractExecuter.Prepare()
	//执行机启动合约执行
	ret, err := contractExecuter.Start()
	if err != nil {
		beegoLog.Error(err)
		_UpdateToFailed(strContractID, strContractHashID, "")
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
func _UpdateToWait(strContractID, strContractHashID, strTaskState string) {
	err := engineCommon.UpdateMonitorWait(strContractID, strContractHashID, strTaskState)
	if err != nil {
		beegoLog.Error(err)
	}
}

//---------------------------------------------------------------------------
func _UpdateToFailed(strContractID, strContractHashID, strTaskState string) {
	err := engineCommon.UpdateMonitorFail(strContractID, strContractHashID, strTaskState)
	if err != nil {
		beegoLog.Error(err)
	}
}

//---------------------------------------------------------------------------
