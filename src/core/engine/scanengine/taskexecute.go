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
		jsonBody := fmt.Sprintf("{\"contract_id\":\"%s\"}", strContractTask.ContractId)
		responseResult, err := chain.GetContract(jsonBody)
		if err != nil || responseResult.Data == nil {
			beegoLog.Error(err)
			err := engineCommon.UpdateMonitorWait(strContractTask.NodePubkey,
				strContractTask.ContractId)
			if err != nil {
				beegoLog.Error(err)
			}
			continue
		}

		if responseResult.Code != _HTTP_OK {
			beegoLog.Error(responseResult.Message)
			continue
		}

		beegoLog.Debug("contract execute")
		contractData, ok := responseResult.Data.([]interface{})
		if !ok {
			beegoLog.Error("responseResult.Data is not ok for type string. type is %T", responseResult.Data)
			err := engineCommon.UpdateMonitorWait(strContractTask.NodePubkey,
				strContractTask.ContractId)
			if err != nil {
				beegoLog.Error(err)
			}
			continue
		}

		slContractData, _ := json.Marshal(contractData)
		beegoLog.Debug(string(slContractData))

		/*go*/
		func(data string) {
			_, err := execengine.Load(data)
			if err != nil {
				beegoLog.Error(err)
				return
			}

			execengine.Prepare()

			ret, err := execengine.Start()
			if err != nil {
				beegoLog.Error(err)
				return
			}
			if ret == 0 {
				beegoLog.Error("合约执行过程中，某任务没有达到执行条件，暂时退出，等待下轮扫描再次加载执行")
			} else if ret == -1 {
				beegoLog.Error("合约执行过程中，某任务执行失败，暂时退出，等待下轮扫描再次加载执行")
			} else if ret == 1 {
				beegoLog.Debug("合约执行完成")
			}
		}(string(slContractData))
	}

	gwgTaskExe.Done()
}

//---------------------------------------------------------------------------
