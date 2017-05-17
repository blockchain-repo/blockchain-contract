// utils
package scanengine

import (
	beegoLog "github.com/astaxie/beego/logs"
)

//---------------------------------------------------------------------------
func Start() {
	if scanEngineConf["clean_data_on"].(int) == 1 {
		beegoLog.Info("CleanTaskSchedule start")
		gwgTaskExe.Add(1)
		go _CleanTaskSchedule()
	}

	beegoLog.Info("ScanFailedTask start")
	gwgTaskExe.Add(1)
	go _ScanFailedTask()

	beegoLog.Info("TaskExecute start")
	gwgTaskExe.Add(1)
	go _TaskExecute()

	beegoLog.Info("ScanTaskSchedule start")
	gwgTaskExe.Add(1)
	go _ScanTaskSchedule()

	gwgTaskExe.Wait()
}

//---------------------------------------------------------------------------
