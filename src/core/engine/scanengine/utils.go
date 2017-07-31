// utils
package scanengine

import (
	"os"
)

import (
	"unicontract/src/common/uniledgerlog"
)

//---------------------------------------------------------------------------
func Start() {
	if scanEngineConf["clean_data_on"].(int) == 1 {
		uniledgerlog.Info("CleanTaskSchedule start")
		gwgTaskExe.Add(1)
		go _CleanTaskSchedule()
	}

	uniledgerlog.Info("ScanFailedTask start")
	gwgTaskExe.Add(1)
	go _ScanFailedTask(0)

	uniledgerlog.Info("ScanWaitTask start")
	gwgTaskExe.Add(1)
	go _ScanFailedTask(1)

	uniledgerlog.Info("execute multi-thread start")
	threadNum, _ := scanEngineConf["execute_thread_num"].(int)
	gPool := new(ThreadPool)
	defer gPool.Stop()
	gPool.Init(threadNum)
	for i := 0; i < threadNum; i++ {
		gPool.AddTask(func() error {
			return _Execute()
		})
	}
	go gPool.Start()

	uniledgerlog.Info("TaskExecute start")
	gwgTaskExe.Add(1)
	go _TaskExecute()

	uniledgerlog.Info("ScanTaskSchedule start")
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
