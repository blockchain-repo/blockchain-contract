/*************************************************************************
  > File Name: scanfailedtask.go
  > Module:
  > Function: 扫描任务待执行表（TaskSchedule），过滤出表内执行次数已经超过阈值的任务
  > Author: wangyp
  > Company:
  > Department:
  > Mail: wangyepeng87@163.com
  > Created Time: 2017.05.16
 ************************************************************************/
package scanengine

import (
	"encoding/json"
	"time"
)

import (
	beegoLog "github.com/astaxie/beego/logs"
)

import (
	"unicontract/src/common/monitor"
	"unicontract/src/config"
	engineCommon "unicontract/src/core/engine/common"
	"unicontract/src/core/model"
)

//---------------------------------------------------------------------------
func _ScanFailedTask() {
	for {
		start := time.Now()
		var slTasks []model.TaskSchedule
		var slID []string

		beegoLog.Debug("query failed data")
		strNodePubkey := config.Config.Keypair.PublicKey
		retStr, err := engineCommon.GetMonitorFailedData(strNodePubkey,
			gParam.FailedCountThreshold)
		if err != nil {
			beegoLog.Error(err.Error())
			goto CONSUME
		}

		if len(retStr) == 0 {
			beegoLog.Debug("no failed data")
			goto CONSUME
		}

		beegoLog.Debug("get failed tasks")
		json.Unmarshal([]byte(retStr), &slTasks)

		beegoLog.Debug("handle task")
		slID = getFailedTaskID(slTasks)

		beegoLog.Debug("handle task")
		engineCommon.UpdateMonitorSendBatch(slID)

		//task fail count send to monitor,modify value
		monitor.Monitor.Gauge("task_fail_count", 1)

	CONSUME:
		consume := time.Since(start)
		if consume.Seconds() < float64(gParam.SleepTime) {
			time.Sleep((time.Duration(float64(gParam.SleepTime) -
				consume.Seconds())) * time.Second)
		}
	}

	gwgTaskExe.Done()
}

//---------------------------------------------------------------------------
func getFailedTaskID(slTasks []model.TaskSchedule) []string {
	var slID []string
	for index, _ := range slTasks {
		slID = append(slID, slTasks[index].Id)
	}
	return slID
}

//---------------------------------------------------------------------------
