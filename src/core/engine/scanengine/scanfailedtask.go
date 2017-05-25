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
	"fmt"
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
func _ScanFailedTask(flag int) {
	var strThresholdName, strLogFlag string
	if flag == 0 { // fail
		strThresholdName = "failed_count_threshold"
		strLogFlag = "failed"
	} else if flag == 1 { // wait
		strThresholdName = "wait_count_threshold"
		strLogFlag = "wait"
	}

	for {
		start := time.Now()
		var slTasks []model.TaskSchedule
		var slID []interface{}

		beegoLog.Debug("query " + strLogFlag + " data")
		strNodePubkey := config.Config.Keypair.PublicKey
		retStr, err := engineCommon.GetMonitorNoSuccessData(strNodePubkey,
			scanEngineConf[strThresholdName].(int), flag)
		if err != nil {
			beegoLog.Error(err.Error())
			goto CONSUME
		}

		if len(retStr) == 0 {
			beegoLog.Debug("no " + strLogFlag + " data")
			goto CONSUME
		}

		beegoLog.Debug("get " + strLogFlag + " tasks")
		json.Unmarshal([]byte(retStr), &slTasks)

		beegoLog.Debug("get task id slice")
		slID = getTaskID(slTasks)

		beegoLog.Debug("handle task")
		engineCommon.UpdateMonitorSendBatch(slID)

		//task fail count send to monitor,modify value
		monitor.Monitor.Gauge(fmt.Sprintf("task_%s_count", strLogFlag), 1)

	CONSUME:
		consume := time.Since(start)
		if consume.Seconds() < float64(scanEngineConf["sleep_time"].(int)) {
			time.Sleep((time.Duration(float64(scanEngineConf["sleep_time"].(int)) -
				consume.Seconds())) * time.Second)
		}
	}

	gwgTaskExe.Done()
}

//---------------------------------------------------------------------------
func getTaskID(slTasks []model.TaskSchedule) []interface{} {
	var slID []interface{}
	for index, _ := range slTasks {
		slID = append(slID, slTasks[index].Id)
	}
	return slID
}

//---------------------------------------------------------------------------
