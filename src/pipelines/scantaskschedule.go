// scantaskschedule
package pipelines

// 扫描任务待执行表（TaskSchedule），过滤出表内属于本节点的任务，放入任务待执行队列（gchTaskQueue）

import (
	"encoding/json"
	"sync"
	"time"
)

import (
	beegoLog "github.com/astaxie/beego/logs"
)

import (
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
)

//---------------------------------------------------------------------------
const (
	_SLEEPTIME = 30 // 单位是秒
)

var (
	gwgScan sync.WaitGroup
)

//---------------------------------------------------------------------------
func startScanTaskSchedule() {
	beegoLog.Debug("ScanTaskSchedule start")
	gwgScan.Add(1)
	go _ScanTaskSchedule()
	gwgScan.Wait()
}

//---------------------------------------------------------------------------
func _ScanTaskSchedule() {
	for {
		start := time.Now()
		var slTasks []model.TaskSchedule

		beegoLog.Debug("query database")
		strNodePubkey := config.Config.Keypair.PublicKey
		retStr, err := rethinkdb.GetTaskSchedulesNoSend(strNodePubkey)
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

//---------------------------------------------------------------------------
func _SendToList(task model.TaskSchedule) bool {
	beegoLog.Debug("contract [%s] enter queue", task.ContractId)
	gchTaskQueue <- task
	return true
}

//---------------------------------------------------------------------------
// 根据公钥环为每个节点插入带执行任务
func InsertTaskSchedule(taskScheduleBase model.TaskSchedule) error {
	var err error
	allPublicKeys := config.GetAllPublicKey()
	for index, _ := range allPublicKeys {
		var taskSchedule model.TaskSchedule
		taskSchedule.Id = common.GenerateUUID()
		taskSchedule.ContractId = taskScheduleBase.ContractId
		taskSchedule.NodePubkey = allPublicKeys[index]
		taskSchedule.StartTime = taskScheduleBase.StartTime
		taskSchedule.EndTime = taskScheduleBase.EndTime
		taskSchedule.FailedCount = 0
		taskSchedule.SendFlag = 0

		slJson, _ := json.Marshal(taskSchedule)
		err = rethinkdb.InsertTaskSchedule(string(slJson))
		if err != nil {
			beegoLog.Error("insert [%s] TaskSchedule is error, error is %s",
				taskScheduleBase.ContractId, err.Error())
			break
		}
	}
	return err
}

//---------------------------------------------------------------------------
