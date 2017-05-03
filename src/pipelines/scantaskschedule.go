// scantaskschedule
package pipelines

import (
	"encoding/json"
	"time"
)

import (
	beegoLog "github.com/astaxie/beego/logs"
)

import (
	"unicontract/src/config"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
)

const (
	_SLEEPTIME = 30 // 单位是秒
)

func _ScanTaskSchedule() {
	for {
		start := time.Now()
		var slTasks []model.TaskSchedule

		beegoLog.Debug("query database")
		strNodePubkey := config.Config.Keypair.PublicKey
		retStr, err := rethinkdb.GetTaskSchedules(strNodePubkey)
		if err != nil {
			beegoLog.Error(err.Error())
			// TODO error handle
			goto CONSUME
		}

		beegoLog.Debug("get tasks")
		err = json.Unmarshal([]byte(retStr), &slTasks)
		if err != nil {
			beegoLog.Error(err.Error())
			// TODO error handle
			goto CONSUME
		}

		beegoLog.Debug("handle task")
		for _, value := range slTasks {
			beegoLog.Debug("send task")
			if _SendToList(value.ContractId) {
				err = rethinkdb.SetTaskScheduleSend(value.Id)
				if err != nil {
					beegoLog.Error(err.Error())
					// TODO error handle
					goto CONSUME
				}
			}
		}

	CONSUME:
		consume := time.Since(start)
		if consume < _SLEEPTIME {
			time.Sleep((_SLEEPTIME - consume) * time.Second)
		}
	}
}

func _SendToList(strContractID string) bool {
	gchTaskListID <- strContractID
	return true
}

func startScanTaskSchedule() {
	go _ScanTaskSchedule()
}
