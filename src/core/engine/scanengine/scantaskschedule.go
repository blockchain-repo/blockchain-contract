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
	"fmt"
	"time"
)

import (
	"unicontract/src/common/monitor"
	"unicontract/src/common/uniledgerlog"
	"unicontract/src/config"
	engineCommon "unicontract/src/core/engine/common"
	"unicontract/src/core/engine/common/db"
)

//---------------------------------------------------------------------------
func _ScanTaskSchedule() {
	sleep_time, ok := scanEngineConf["sleep_time"].(int)
	if !ok {
		panic("scanEngineConf[\"sleep_time\"].(int) assert error")
	}
	ticker := time.NewTicker(time.Second * time.Duration(sleep_time))

	for _ = range ticker.C {
		open, ok := scanEngineConf["open"].(bool)
		if !ok {
			panic("scanEngineConf[\"open\"].(bool) assert error")
		}
		if open {
			uniledgerlog.Info(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "query no send data"))
			strNodePubkey := config.Config.Keypair.PublicKey
			failed_count_threshold, ok := scanEngineConf["failed_count_threshold"].(int)
			if !ok {
				panic("scanEngineConf[\"failed_count_threshold\"].(int) assert error")
			}
			retStr, err := engineCommon.GetMonitorNoSendData(strNodePubkey, failed_count_threshold)
			if err != nil {
				uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, err.Error()))
				continue
			}

			if len(retStr) == 0 {
				uniledgerlog.Info(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "all send"))
				continue
			}

			uniledgerlog.Info(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "get no send tasks"))
			var slTasks []db.TaskSchedule
			json.Unmarshal([]byte(retStr), &slTasks)

			uniledgerlog.Info(fmt.Sprintf("[%s][%s]", uniledgerlog.NO_ERROR, "handle task"))
			for _, value := range slTasks {
				gchTaskQueue <- value
				//wsp@monitor
				monitor.Monitor.Count("task_running", 1)
				err = engineCommon.UpdateMonitorSend(value.Id)
				if err != nil {
					uniledgerlog.Error(fmt.Sprintf("[%s][%s]", uniledgerlog.OTHER_ERROR, err.Error()))
				}
			}
		}
	}

	gwgTaskExe.Done()
}

//---------------------------------------------------------------------------
