// utils
package scanengine

import (
	"os"
)

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
	go _ScanFailedTask(0)

	beegoLog.Info("ScanWaitTask start")
	gwgTaskExe.Add(1)
	go _ScanFailedTask(1)

	beegoLog.Info("TaskExecute start")
	gwgTaskExe.Add(1)
	go _TaskExecute()

	beegoLog.Info("ScanTaskSchedule start")
	gwgTaskExe.Add(1)
	go _ScanTaskSchedule()

	gwgTaskExe.Wait()
}

//---------------------------------------------------------------------------
func _WriteFile(fileName string, content string) (int, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.WriteString(content)
}

//---------------------------------------------------------------------------
