// scantaskschedule_test
package pipelines

import (
	"testing"
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/model"
)

func Test_startScanTaskSchedule(t *testing.T) {
	config.Init()
	go startTaskExecute()
	startScanTaskSchedule()
}

func Test_InsertTaskSchedule(t *testing.T) {
	config.Init()

	var taskSchedule model.TaskSchedule
	taskSchedule.ContractId = common.GenerateUUID()
	taskSchedule.StartTime = common.GenTimestamp()
	taskSchedule.EndTime = common.GenTimestamp()

	err := InsertTaskSchedule(taskSchedule)
	if err != nil {
		t.Error(err)
	}
}
