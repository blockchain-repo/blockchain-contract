// common_test
package headNodeMonitor

import (
	"testing"
	"unicontract/src/core/control"
)

func Test_Init(t *testing.T) {
	t.Logf("%+v\n", gslPublicKeys)
	t.Logf("%+v\n", headNodeMonitorConf)
	control.Init()
	Init()
	t.Logf("%+v\n", gslPublicKeys)
	t.Logf("%+v\n", headNodeMonitorConf)
}
