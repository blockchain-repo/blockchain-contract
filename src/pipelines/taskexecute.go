// taskexecute
package pipelines

import (
	"fmt"
)

import (
	beegoLog "github.com/astaxie/beego/logs"
)

import (
	"unicontract/src/chain"
	"unicontract/src/core/db/rethinkdb"
)

const (
	_TASKLISTLEN = 20
	_THRESHOLD   = 50
)

var (
	gchTaskListID chan string
)

func init() {
	gchTaskListID = make(chan string, _TASKLISTLEN)
}

func _TaskExecute() {
	for {
		strContractID, ok := <-gchTaskListID
		if !ok {
			break
		}

		jsonBody := fmt.Sprintf("{\"contract_id\":\"%s\"}", strContractID)
		responseResult, err := chain.GetContract(jsonBody)
		if err != nil {
			beegoLog.Error(err)
			continue
		}

		contractData := responseResult.Data.(string)
		go func(data string) {
			// TODO 调用执行机接口进行load和start，根据返回结果决定后续操作
			if true { // TODO 执行成功

			} else { // TODO 执行失败
				err = rethinkdb.SetTaskScheduleSend(strContractID)
				if err != nil {
					beegoLog.Error(err)
					return
				}

				failedCount, err := rethinkdb.SetTaskScheduleFailedCount(strContractID)
				if err != nil {
					beegoLog.Error(err)
					return
				}

				if failedCount >= _THRESHOLD {
					// TODO 执行失败次数超过阈值，要进行处理
				}
			}
		}(contractData)
	}
}

func startTaskExecute() {
	go _TaskExecute()
}
