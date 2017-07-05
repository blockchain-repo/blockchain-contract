/*************************************************************************
  > File Name: scantaskschedule.go
  > Module:
  > Function: 扫描任务待执行表（TaskSchedule），过滤出表内属于本节点的任务，
              放入任务待执行队列（gchTaskQueue）
  > Author: wangyp
  > Company:
  > Department:
  > Mail: wangyepeng87@163.com
  > Created Time: 2017.05.08
 ************************************************************************/
package scanengine

import (
	"encoding/json"
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
func _ScanTaskSchedule() {
	ticker := time.NewTicker(time.Second * time.Duration(scanEngineConf["sleep_time"].(int)))
	for _ = range ticker.C {
		uniledgerlog.Debug("query no send data")
		strNodePubkey := config.Config.Keypair.PublicKey
		retStr, err := engineCommon.GetMonitorNoSendData(strNodePubkey,
			scanEngineConf["failed_count_threshold"].(int))
		if err != nil {
			uniledgerlog.Error(err.Error())
			continue
		}

		if len(retStr) == 0 {
			uniledgerlog.Debug("all send")
			continue
		}

		uniledgerlog.Debug("get no send tasks")
		var slTasks []model.TaskSchedule
		json.Unmarshal([]byte(retStr), &slTasks)

		uniledgerlog.Debug("handle task")
		for _, value := range slTasks {
			uniledgerlog.Debug("contract [%v] enter queue", value)
			gchTaskQueue <- value
			//wsp@monitor
			monitor.Monitor.Count("task_running", 1)
			err = engineCommon.UpdateMonitorSend(value.Id)
			if err != nil {
				uniledgerlog.Error(err.Error())
			}
		}
	}

	gwgTaskExe.Done()
}

//---------------------------------------------------------------------------
