// scantaskschedule
package taskexecute

// 扫描任务待执行表（TaskSchedule），过滤出表内属于本节点的任务，放入任务待执行队列（gchTaskQueue）

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

//---------------------------------------------------------------------------
func _ScanTaskSchedule() {
	for {
		start := time.Now()
		var slTasks []model.TaskSchedule

		beegoLog.Debug("query no send data")
		strNodePubkey := config.Config.Keypair.PublicKey
		retStr, err := rethinkdb.GetTaskSchedulesNoSend(strNodePubkey)
		if err != nil {
			beegoLog.Error(err.Error())
			goto CONSUME
		}

		if len(retStr) == 0 {
			beegoLog.Debug("all send")
			goto CONSUME
		}

		beegoLog.Debug("get tasks")
		json.Unmarshal([]byte(retStr), &slTasks)

		beegoLog.Debug("handle task")
		for _, value := range slTasks {
			beegoLog.Debug("contract [%s] enter queue", value)
			gchTaskQueue <- value
			err = rethinkdb.SetTaskScheduleSend(value.Id)
			if err != nil {
				beegoLog.Error(err.Error())
				goto CONSUME
			}
		}

	CONSUME:
		consume := time.Since(start)
		if consume.Seconds() < float64(_SLEEPTIME) {
			time.Sleep((time.Duration(float64(_SLEEPTIME) - consume.Seconds())) * time.Second)
		}
	}

	gwgTaskExe.Done()
}

//---------------------------------------------------------------------------
