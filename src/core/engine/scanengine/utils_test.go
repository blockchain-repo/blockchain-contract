// scantaskschedule_test
package scanengine

import (
	"strconv"
	"testing"
	"time"
	"unicontract/src/common"
	"unicontract/src/config"
	"unicontract/src/core/model"
)

func Test_Start(t *testing.T) {
	config.Init()

	Start()
}

func Test_InsertTaskSchedule(t *testing.T) {
	config.Init()

	var taskSchedule model.TaskSchedule
	taskSchedule.ContractId = common.GenerateUUID()
	taskSchedule.StartTime = common.GenTimestamp()
	taskSchedule.EndTime = strconv.FormatInt(time.Now().Add(time.Hour*24*5).UnixNano()/1000000, 10)

	err := InsertTaskSchedules(taskSchedule)
	if err != nil {
		t.Error(err)
	}
}
