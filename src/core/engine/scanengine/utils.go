// utils
package scanengine

import (
	beegoLog "github.com/astaxie/beego/logs"
)

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
