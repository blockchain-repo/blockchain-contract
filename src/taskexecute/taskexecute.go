// taskexecute
package taskexecute

// 从任务待执行队列（gchTaskQueue）中取任务，然后放入执行机执行

import (
	"fmt"
)

import (
	beegoLog "github.com/astaxie/beego/logs"
)

import (
	"unicontract/src/chain"
	//"unicontract/src/common/requestHandler"
	"unicontract/src/core/db/rethinkdb"
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
		if err != nil {
			beegoLog.Error(err)
			err := rethinkdb.SetTaskScheduleNoSend(strContractTask.Id)
			if err != nil {
				beegoLog.Error(err)
			}
			continue
		}

		beegoLog.Debug("contract execute")
		contractData := responseResult.Data.(string)
		go func(data string) {
			// TODO 调用执行机接口进行load和start，返回结果分为三种：
			//  1：整个合约执行成功
			//  0：合约中某个步骤没有达到执行条件
			// -1：合约中某个步骤执行失败

			/* 以下修改数据库的部分挪到执行机内部操作
			ret := false
			if ret { // TODO 执行成功
				beegoLog.Debug("execute success")
			} else { // TODO 执行失败
				beegoLog.Debug("execute failed")
				err := rethinkdb.SetTaskScheduleNoSend(strContractTask.Id)
				if err != nil {
					beegoLog.Error(err)
					return
				}

				failedCount, err := rethinkdb.SetTaskScheduleFailedCount(strContractTask.Id)
				if err != nil {
					beegoLog.Error(err)
					return
				}

				if failedCount >= _THRESHOLD {
					// TODO 执行失败次数超过阈值，要进行处理
				}
			}*/
		}(contractData)
	}

	gwgTaskExe.Done()
}

//---------------------------------------------------------------------------
