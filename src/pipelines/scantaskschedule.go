// scantaskschedule
package pipelines

import (
	"encoding/json"
	"sync"
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

var (
	gwgScan sync.WaitGroup
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
			goto CONSUME
		}

		if len(retStr) == 0 {
			beegoLog.Debug("query result is null")
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
			if _SendToList(value) {
				err = rethinkdb.SetTaskScheduleSend(value.Id)
				if err != nil {
					beegoLog.Error(err.Error())
					goto CONSUME
				}
			}
		}

	CONSUME:
		consume := time.Since(start)
		if consume.Seconds() < float64(_SLEEPTIME) {
			time.Sleep((time.Duration(float64(_SLEEPTIME) - consume.Seconds())) * time.Second)
		}
	}

	gwgScan.Done()
}

func _SendToList(task model.TaskSchedule) bool {
	beegoLog.Debug("contract [%s] enter queue", task.ContractId)
	gchTaskQueue <- task
	return true
}

func startScanTaskSchedule() {
	beegoLog.Debug("ScanTaskSchedule start")
	gwgScan.Add(1)
	go _ScanTaskSchedule()
	gwgScan.Wait()
}
