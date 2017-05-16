package filters

import (
	"github.com/astaxie/beego/context"
	"unicontract/src/common/monitor"
)

func MonitorFilter(ctx *context.Context) {
	defer monitor.Monitor.NewTiming().Send("all_request")
}
