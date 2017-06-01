// utils
package headNodeMonitor

import (
	beegoLog "github.com/astaxie/beego/logs"
)

//---------------------------------------------------------------------------
func Start() {
	beegoLog.Info("HeadNodeMonitor start")
	gwgHeadNodeMonitor.Add(1)
	go _HeadNodeMonitor()

	gwgHeadNodeMonitor.Wait()
}

//---------------------------------------------------------------------------
