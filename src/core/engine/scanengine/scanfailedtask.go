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
	"unicontract/src/common/monitor"
	"unicontract/src/common/uniledgerlog"
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

	ticker := time.NewTicker(time.Second * time.Duration(scanEngineConf["sleep_time"].(int)))
	for _ = range ticker.C {
		uniledgerlog.Debug("query " + strLogFlag + " data")
		strNodePubkey := config.Config.Keypair.PublicKey
		retStr, err := engineCommon.GetMonitorNoSuccessData(strNodePubkey,
			scanEngineConf[strThresholdName].(int), flag)
		if err != nil {
			uniledgerlog.Error(err.Error())
			continue
		}

		if len(retStr) == 0 {
			uniledgerlog.Debug("no " + strLogFlag + " data")
			continue
		}

		uniledgerlog.Debug("get " + strLogFlag + " tasks")
		var slTasks []model.TaskSchedule
		json.Unmarshal([]byte(retStr), &slTasks)

		uniledgerlog.Debug("get task id slice")
		slID := _GetTaskID(slTasks)

		uniledgerlog.Debug("handle task")
		engineCommon.UpdateMonitorSendBatch(slID)

		uniledgerlog.Debug("record task")
		_Record(flag, slID)

		//task fail count send to monitor,modify value
		monitor.Monitor.Gauge(fmt.Sprintf("task_%s_count", strLogFlag), 1)
	}

	gwgTaskExe.Done()
}

//---------------------------------------------------------------------------
func _GetTaskID(slTasks []model.TaskSchedule) []interface{} {
	var slID []interface{}
	for index, _ := range slTasks {
		slID = append(slID, slTasks[index].Id)
	}
	return slID
}

//---------------------------------------------------------------------------
func _Record(flag int, slID []interface{}) {
	var strRecordFile string
	if flag == 0 {
		strRecordFile = scanEngineConf["record_f_file_path"].(string)
	} else if flag == 1 {
		strRecordFile = scanEngineConf["record_w_file_path"].(string)
	}

	var strID string
	for _, v := range slID {
		strID = fmt.Sprintf("%s\n%s", strID, v.(string))
	}

	writeCount, err := _WriteFile(strRecordFile, strID)
	if err != nil {
		uniledgerlog.Error(err)
	}
	if writeCount != len(strID) {
		uniledgerlog.Error("write count is error")
	}
}

//---------------------------------------------------------------------------
