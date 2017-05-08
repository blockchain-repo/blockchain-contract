// utils
package scanengine

import (
	"encoding/json"
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
// 根据公钥环为每个节点插入待执行任务
func InsertTaskSchedules(taskScheduleBase model.TaskSchedule) error {
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
func Start() {
	beegoLog.Info("CleanTaskSchedule start")
	gwgTaskExe.Add(1)
	go _CleanTaskSchedule()

	beegoLog.Info("TaskExecute start")
	gwgTaskExe.Add(1)
	go _TaskExecute()

	beegoLog.Info("ScanTaskSchedule start")
	gwgTaskExe.Add(1)
	go _ScanTaskSchedule()

	gwgTaskExe.Wait()
}
