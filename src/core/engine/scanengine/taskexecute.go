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
type executeParam struct {
	strData           string
	strContractID     string
	strContractHashID string
}

//---------------------------------------------------------------------------
func _TaskExecute() {
	for {
		uniledgerlog.Info(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "wait for ContractTask ..."))
		strContractTask, ok := <-gchTaskQueue
		if !ok {
			break
		}
		uniledgerlog.Info(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "get ContractTask"))
		//wsp@monitor
		monitor.Monitor.Count("task_running", -1)
		uniledgerlog.Info(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "query contract base on contractId"))
		jsonBody := fmt.Sprintf("{\"contract_hash_id\":\"%s\"}", strContractTask.ContractHashId)
		//responseResult:  requestHandler.ResponseResult, data中存的是完整的Output结构体
		responseResult, err := chain.GetTxByConHashId(jsonBody)
		if err != nil || responseResult.Data == nil {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, err.Error()))
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}

		if responseResult.Code != _HTTP_OK {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, fmt.Sprintf("responseResult.Code is [ %d ]", responseResult.Code)))
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, fmt.Sprintf("responseResult.Message is [ %s ]", responseResult.Message)))
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}

		uniledgerlog.Info(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "get responseResult data"))
		// 1
		contractData, ok := responseResult.Data.(interface{})
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, "responseResult.Data.(interface{}) is error"))
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}

		// 2
		mapData, ok := contractData.([]interface{})
		if !ok {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, "contractData.([]map[string]interface{}) error"))
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}

		if len(mapData) == 0 { // 没有查询到contract
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, "get contract is null"))
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}

		slContractData, err := json.Marshal(mapData[0])
		if err != nil {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.SERIALIZE_ERROR, err.Error()))
			_UpdateToWait(strContractTask.ContractId, strContractTask.ContractHashId)
			continue
		}

		uniledgerlog.Info(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "contract execute"))
		param := executeParam{
			strData:           string(slContractData),
			strContractID:     strContractTask.ContractId,
			strContractHashID: strContractTask.ContractHashId,
		}
		gchExecParamQueue <- param
	}

	gwgTaskExe.Done()
}

//---------------------------------------------------------------------------
func _Execute() error {
	uniledgerlog.Info(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "Execute into ......"))
	for {
		param, ok := <-gchExecParamQueue
		if !ok {
			break
		}

		contractExecuter := execengine.NewContractExecuter()
		//strData为完整的Output结构体
		task_load_time := monitor.Monitor.NewTiming()
		err := contractExecuter.Load(param.strData)
		task_load_time.Send("task_load")
		if err != nil {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, err.Error()))
			_UpdateToFailed(param.strContractID, param.strContractHashID)
			continue
		}
		//执行引擎初始化环境
		contractExecuter.Prepare()
		//执行机启动合约执行
		task_execute_time := monitor.Monitor.NewTiming()
		ret, err := contractExecuter.Start()
		task_execute_time.Send("task_execute")
		if err != nil {
			uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, err.Error()))
			_UpdateToFailed(param.strContractID, param.strContractHashID)
			continue
		}
		if ret == 0 {
			uniledgerlog.Warn(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "合约执行过程中，某任务没有达到执行条件，暂时退出，等待下轮扫描再次加载执行"))
		} else if ret == -1 {
			uniledgerlog.Warn(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "合约执行过程中，某任务执行失败，暂时退出，等待下轮扫描再次加载执行"))
		} else if ret == 1 {
			uniledgerlog.Info(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "合约任务执行完成"))
		}
		//执行机销毁合约
		contractExecuter.Destory()
	}
	return nil
}

//---------------------------------------------------------------------------
func _UpdateToWait(strContractID, strContractHashID string) {
	var waitStruct engineCommon.UpdateMonitorWaitStruct
	waitStruct.WstrContractID = strContractID
	waitStruct.WstrContractHashID = strContractHashID
	waitStruct.WstrTaskId = "0"
	waitStruct.WstrTaskState = ""
	waitStruct.WnTaskExecuteIndex = 1
	slWaitData, _ := json.Marshal(waitStruct)
	err := engineCommon.UpdateMonitorWait(string(slWaitData))
	if err != nil {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, err.Error()))
	}
}

//---------------------------------------------------------------------------
func _UpdateToFailed(strContractID, strContractHashID string) {
	var failStruct engineCommon.UpdateMonitorFailStruct
	failStruct.FstrContractID = strContractID
	failStruct.FstrContractHashID = strContractHashID
	failStruct.FstrTaskId = "0"
	failStruct.FstrTaskState = ""
	failStruct.FnTaskExecuteIndex = 1
	slFailData, _ := json.Marshal(failStruct)
	err := engineCommon.UpdateMonitorFail(string(slFailData))
	if err != nil {
		uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, err.Error()))
	}
}

//---------------------------------------------------------------------------
