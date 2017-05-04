// scantaskschedule_test
package pipelines

import (
	"testing"
	"unicontract/src/config"
)

func Test_startScanTaskSchedule(t *testing.T) {
	config.Init()
	go startTaskExecute()
	startScanTaskSchedule()
}
