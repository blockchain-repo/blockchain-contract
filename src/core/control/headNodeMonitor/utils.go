// utils
package headNodeMonitor

import (
	"unicontract/src/common/uniledgerlog"
)

//---------------------------------------------------------------------------
func Start() {
	uniledgerlog.Info("HeadNodeMonitor start")
	gwgHeadNodeMonitor.Add(1)
	go _HeadNodeMonitor()

	gwgHeadNodeMonitor.Wait()
}

//---------------------------------------------------------------------------
