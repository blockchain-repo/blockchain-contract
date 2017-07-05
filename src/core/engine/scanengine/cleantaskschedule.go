/*************************************************************************
  > File Name: cleantaskschedule.go
  > Module:
  > Function: 清理数据表（TaskSchedule）中的过期或者已经执行成功的任务
  > Author: wangyp
  > Company:
  > Department:
  > Mail: wangyepeng87@163.com
  > Created Time: 2017.05.08
 ************************************************************************/
package scanengine

import (
	"encoding/json"
	"strconv"
	"time"
)

import (
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/config"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
)

//---------------------------------------------------------------------------
func _CleanTaskSchedule() {
	ticker := time.NewTicker(time.Minute * (time.Duration)(scanEngineConf["clean_time"].(int)))
	for _ = range ticker.C {
		uniledgerlog.Debug("query all success task")
		strSuccessTask, err :=
			rethinkdb.GetTaskSchedulesSuccess(config.Config.Keypair.PublicKey)
		if err != nil {
			uniledgerlog.Error(err)
			continue
		}

		if len(strSuccessTask) == 0 {
			uniledgerlog.Debug("success task is null")
			continue
		}

		var slTasks []model.TaskSchedule
		json.Unmarshal([]byte(strSuccessTask), &slTasks)

		uniledgerlog.Debug("success task filter")
		slID := _TaskFilter(slTasks)

		if len(slID) == 0 {
			uniledgerlog.Debug("_TaskFilter return is null")
			continue
		}

		uniledgerlog.Debug("success task delete")
		deleteNum, err := rethinkdb.DeleteTaskSchedules(slID)
		if deleteNum != len(slID) {
			uniledgerlog.Error(err)
		}
	}
	gwgTaskExe.Done()
}

//---------------------------------------------------------------------------
func _TaskFilter(slTasks []model.TaskSchedule) []interface{} {
	var slID []interface{}

	// 过滤时间点
	cleanTimePoint := time.Now().
		Add(-time.Hour*24*(time.Duration)(scanEngineConf["clean_data_time"].(int))).
		UnixNano() / 1000000
	for index, value := range slTasks {
		nTimePoint, err := strconv.Atoi(value.LastExecuteTime)
		if err != nil {
			uniledgerlog.Error(err)
			continue
		}
		if int64(nTimePoint) < cleanTimePoint {
			slID = append(slID, slTasks[index].Id)
		}
	}

	return slID
}

//---------------------------------------------------------------------------
