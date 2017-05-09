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
	beegoLog "github.com/astaxie/beego/logs"
)

import (
	"unicontract/src/config"
	"unicontract/src/core/db/rethinkdb"
	"unicontract/src/core/model"
)

//---------------------------------------------------------------------------
func _CleanTaskSchedule() {
	for {
		ticker := time.NewTicker(time.Minute * (time.Duration)(gParam.CleanTime))
		beegoLog.Debug("wait for clean data...")
		select {
		case <-ticker.C:
			beegoLog.Debug("query all success task")
			strSuccessTask, err :=
				rethinkdb.GetTaskSchedulesSuccess(config.Config.Keypair.PublicKey)
			if err != nil {
				beegoLog.Error(err)
				continue
			}

			if len(strSuccessTask) == 0 {
				beegoLog.Debug("success task is null")
				continue
			}

			var slTasks []model.TaskSchedule
			json.Unmarshal([]byte(strSuccessTask), &slTasks)

			beegoLog.Debug("success task filter")
			slID := _TaskFilter(slTasks)

			if len(slID) == 0 {
				beegoLog.Debug("_TaskFilter return is null")
				continue
			}

			beegoLog.Debug("success task delete")
			deleteNum, slerr := rethinkdb.DeleteTaskSchedules(slID)
			if deleteNum != len(slID) {
				for _, value := range slerr {
					beegoLog.Error("id is [%s] delete failed", value.Error())
				}
			}
		}
	}
	gwgTaskExe.Done()
}

//---------------------------------------------------------------------------
func _TaskFilter(slTasks []model.TaskSchedule) []string {
	var slID []string

	// 过滤时间点
	cleanTimePoint := time.Now().
		Add(-time.Hour*24*(time.Duration)(gParam.CleanDataTime)).
		UnixNano() / 1000000
	for index, value := range slTasks {
		nTimePoint, err := strconv.Atoi(value.LastExecuteTime)
		if err != nil {
			beegoLog.Error(err)
			continue
		}
		if int64(nTimePoint) < cleanTimePoint {
			slID = append(slID, slTasks[index].Id)
		}
	}

	return slID
}

//---------------------------------------------------------------------------
