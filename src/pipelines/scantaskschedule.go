// scantaskschedule
package pipelines

import (
	"fmt"
	"time"
)

import (
	beegoLog "github.com/astaxie/beego/logs"
)

var (
	_ = 123
)

const (
	_TableNameTaskSchedule = "Votes"
	_SLEEPTIME             = 10
)

func _ScanTaskSchedule() {
	beegoLog.Debug("_ScanTaskSchedule")
	for {
		start := time.Now()

		//TODO real handle

		consume := time.Since(start)
		if consume < _SLEEPTIME {
			time.Sleep((_SLEEPTIME - consume) * time.Second)
		}
	}
}

func _SendToList(strTaskSchedule string) bool {
	return true
}

func startScanTaskSchedule() {
	fmt.Println("startScanTaskSchedule")
	go _ScanTaskSchedule()
}
